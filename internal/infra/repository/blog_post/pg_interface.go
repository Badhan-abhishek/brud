package blogpostrepo

import (
	blogpost "be.blog/internal/domain/blog_post"
	"be.blog/internal/domain/pagination"
)

type Repository interface {
	CreatePost(post *blogpost.BlogPost) (*blogpost.BlogPost, error)
	GetAll(pn *pagination.PaginationInput) (*[]blogpost.BlogPost, int, error)
	Get(id string) (*blogpost.BlogPost, error)
	Delete(id string) error
	Update(post *blogpost.BlogPostUpdate) error
}
