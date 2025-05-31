package blogpostrepo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	entity "be.blog/internal/domain/blog_post"
	"be.blog/internal/domain/pagination"
	customerrors "be.blog/pkg/custom_errors"
	errorutils "be.blog/pkg/utils/error_utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type blogpostRepo struct {
	db *pgxpool.Pool
}

func NewBlogPostRepo(db *pgxpool.Pool) *blogpostRepo {
	return &blogpostRepo{db}
}

func (r *blogpostRepo) CreatePost(post *entity.BlogPost) (*entity.BlogPost, error) {
	row, err := r.db.Query(
		context.Background(),
		`INSERT INTO blog_post (title, body, author, descriptions, slug, published, meta_description) 
	 VALUES ($1, $2, $3, $4, $5, $6, $7) returning id, created_at, updated_at, title, body, author, descriptions, slug, published, meta_description`,
		post.Title, post.Body, post.Author, post.Descriptions, post.Slug, post.Published, post.MetaDescription,
	)
	if err != nil {
		return nil, errorutils.MapRepoError(err)
	}
	var blogPost entity.BlogPost
	if row.Next() {
		err = row.Scan(&blogPost.ID, &blogPost.CreatedAt, &blogPost.UpdatedAt, &blogPost.Title, &blogPost.Body, &blogPost.Author, &blogPost.Descriptions, &blogPost.Slug, &blogPost.Published, &blogPost.MetaDescription)
		if err != nil {
			return nil, errorutils.MapRepoError(err)
		}
	} else {
		return nil, errorutils.MapRepoError(err)
	}
	return &blogPost, nil
}

func (r *blogpostRepo) GetAll(pn *pagination.PaginationInput) (*[]entity.BlogPost, int, error) {
	offset := pagination.CalculateOffset(*pn.Skip, *pn.Limit)
	var posts []entity.BlogPost

	var count int = 0

	err := r.db.QueryRow(
		context.Background(),
		`SELECT COUNT(*) FROM blog_post`,
	).Scan(&count)
	if err != nil {
		return nil, count, errorutils.MapRepoError(err)
	}

	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, title, body, author, descriptions, slug, published, meta_description, created_at, updated_at FROM blog_post limit $1 offset $2`,
		pn.Limit,
		offset,
	)
	if err != nil {
		return nil, count, errorutils.MapRepoError(err)
	}

	for rows.Next() {
		var post entity.BlogPost
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Author, &post.Descriptions, &post.Slug, &post.Published, &post.MetaDescription, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, count, errorutils.MapRepoError(err)
		}
		posts = append(posts, post)
	}
	return &posts, count, nil
}

func (r *blogpostRepo) Get(id string) (*entity.BlogPost, error) {
	var post entity.BlogPost
	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, title, body, author, descriptions, slug, published, meta_description, created_at, updated_at FROM blog_post WHERE id = $1`,
		id,
	).Scan(&post.ID, &post.Title, &post.Body, &post.Author, &post.Descriptions, &post.Slug, &post.Published, &post.MetaDescription, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, errorutils.MapRepoError(err)
	}
	return &post, nil
}

func (r *blogpostRepo) Delete(id string) error {
	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM blog_post WHERE id = $1`,
		id,
	)
	if err != nil {
		return errorutils.MapRepoError(err)
	}

	return nil
}

func (r *blogpostRepo) Update(post *entity.BlogPostUpdate) error {

	if post.ID == "" {
		return errors.New(customerrors.RepoIDNotFound)
	}

	_, err := r.Get(post.ID)

	if err != nil {
		return errorutils.MapRepoError(err)
	}

	columns := []string{}
	args := []any{}
	i := 1

	if post.Title != nil {
		columns = append(columns, "title = $"+fmt.Sprint(i))
		args = append(args, post.Title)
		i++
	}

	if post.Body != nil {
		columns = append(columns, "body = $"+fmt.Sprint(i))
		args = append(args, post.Body)
		i++
	}

	if post.Author != nil {
		columns = append(columns, "author = $"+fmt.Sprint(i))
		args = append(args, post.Author)
		i++
	}

	if post.Descriptions != nil && len(*post.Descriptions) > 0 {
		columns = append(columns, "descriptions = $"+fmt.Sprint(i))
		args = append(args, post.Descriptions)
		i++
	}

	if post.Published != nil {
		columns = append(columns, "published = $"+fmt.Sprint(i))
		args = append(args, post.Published)
		i++
	}

	if post.Slug != nil {
		columns = append(columns, "slug = $"+fmt.Sprint(i))
		args = append(args, post.Slug)
		i++
	}

	args = append(args, post.ID)

	query := fmt.Sprintf("UPDATE blog_post SET %s WHERE id = $%d", strings.Join(columns, ", "), i)

	_, err = r.db.Exec(context.Background(), query, args...)

	if err != nil {
		return errorutils.MapRepoError(err)
	}
	return nil
}
