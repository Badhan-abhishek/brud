package blogpost

type CreateBlogPost struct {
	Title        string   `json:"title" validate:"required"`
	Body         string   `json:"body" validate:"required"`
	Author       string   `json:"author" validate:"required"`
	Descriptions []string `json:"descriptions" validate:"dive,required"`
}

type UpdateBlogPost struct {
	Title        *string   `json:"title"`
	Body         *string   `json:"body"`
	Author       *string   `json:"author"`
	Descriptions *[]string `json:"descriptions"`
}
