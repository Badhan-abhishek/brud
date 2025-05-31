package blogpost

import (
	blogpost "be.blog/internal/domain/blog_post"
	"be.blog/internal/domain/pagination"
)

type Usecase interface {
	CreatePost(title, content, author string, descriptions []string, slug *string) (*blogpost.BlogPost, error)
	GetAll(limit *int, skip *int) (*pagination.PaginationOutput, error)
	Get(id string) (*blogpost.BlogPost, error)
	Delete(id string) error
	Update(id string, title, author, body *string, descriptions *[]string, published *bool) error
}
