package blogpost

import (
	entity "be.blog/internal/domain/blog_post"
	"be.blog/internal/domain/pagination"
	blogpostrepo "be.blog/internal/infra/repository/blog_post"
)

type usecase struct {
	repo blogpostrepo.Repository
}

func NewBlogPostUsecase(repo blogpostrepo.Repository) Usecase {
	return &usecase{repo: repo}
}

func (s *usecase) CreatePost(title, content, author string, tags []string, slug *string) (*entity.BlogPost, error) {
	post, err := entity.NewBlogPost(title, content, author, tags, slug)
	if err != nil {
		return nil, err
	}
	newBlogPost, err := s.repo.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return newBlogPost, nil
}

func (s *usecase) GetAll(limit *int, skip *int) (*pagination.PaginationOutput, error) {
	pn := pagination.CreatePagination(*skip, *limit)
	posts, count, err := s.repo.GetAll(pn)
	if err != nil {
		return nil, err
	}
	return pagination.CreatePaginationOutput(count, *pn.Skip, *pn.Limit, posts), nil
}

func (s *usecase) Get(id string) (*entity.BlogPost, error) {
	post, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *usecase) Delete(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *usecase) Update(id string, title, author, body *string, descriptions *[]string, published *bool) error {
	post, err := entity.UpdateBlogPost(id, title, author, body, descriptions, published)
	if err != nil {
		return err
	}
	if err := s.repo.Update(post); err != nil {
		return err
	}
	return nil
}
