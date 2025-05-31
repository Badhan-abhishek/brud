package blogpost

import (
	"errors"
	"time"

	"be.blog/pkg/custom_errors"
	stringutils "be.blog/pkg/utils/string_utils"
)

type BlogPost struct {
	ID              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Title           string
	Body            string
	Author          string
	Descriptions    []string
	Published       bool
	Slug            string
	MetaDescription string
}

type BlogPostUpdate struct {
	ID              string
	Title           *string
	Body            *string
	Author          *string
	Descriptions    *[]string
	Published       *bool
	MetaDescription *string
	Slug            *string
}

func NewBlogPost(title string, content string, author string, tags []string, slug *string) (*BlogPost, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if content == "" {
		return nil, errors.New("content cannot be empty")
	}

	if author == "" {
		author = "Anonymous"
	}

	if len(tags) == 0 {
		tags = []string{"general"}
	}

	if slug == nil {
		s := stringutils.Slugify(title)
		slug = &s
	}
	return &BlogPost{
		Title:           title,
		Body:            content,
		Slug:            *slug,
		Descriptions:    tags,
		Author:          author,
		Published:       true,
		MetaDescription: stringutils.Truncate(content, 160),
	}, nil
}

func UpdateBlogPost(id string, title, author, body *string, descriptions *[]string, published *bool) (*BlogPostUpdate, error) {
	if id == "" {
		return nil, errors.New(customerrors.DomainMissingID)
	}

	var metaDescription *string
	if body != nil {
		desc := stringutils.Truncate(*body, 160)
		metaDescription = &desc
	}

	var slug *string
	if title != nil {
		s := stringutils.Slugify(*title)
		slug = &s
	}

	return &BlogPostUpdate{
		ID:              id,
		Title:           title,
		Author:          author,
		Body:            body,
		Descriptions:    descriptions,
		Published:       published,
		MetaDescription: metaDescription,
		Slug:            slug,
	}, nil
}
