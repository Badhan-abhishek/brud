package blogpost

import (
	"errors"
	"log"

	_ "be.blog/docs"
	usecase "be.blog/internal/usecase/blog_post"
	customerrors "be.blog/pkg/custom_errors"
	"be.blog/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase usecase.Usecase
}

func NewBlogPostHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {

	r.Post("/", h.CreateBlogPost)

	r.Get("/", h.GetAll)

	r.Get("/:id", h.Get)

	r.Delete("/:id", h.Delete)

	r.Patch("/:id", h.Update)
}

// @Tags			BlogPost
// @Summary		Create a new blog post
// @Description	Create a new blog post with title, body, author, and descriptions
// @Accept			json
// @Produce		json
// @Param			blogPost	body		CreateBlogPost	true	"Blog Post"
// @Success		201			{object}	blogpost.BlogPost
// @Failure		400			{object}	response.ErrorResponse
// @Failure		409			{object}	response.ErrorResponse
// @Failure		500			{object}	response.ErrorResponse
// @Router			/api/blog-post [post]
func (handler *Handler) CreateBlogPost(c *fiber.Ctx) error {
	errorHandler := response.NewErrorHandler(c, "Blog Post", "Blog Posts")

	var input CreateBlogPost
	if err := c.BodyParser(&input); err != nil {
		return errorHandler.SendError(fiber.StatusBadRequest, "Invalid input data")
	}
	post, err := handler.usecase.CreatePost(input.Title, input.Body, input.Author, input.Descriptions, nil)
	if err != nil {
		log.Println("Error creating blog post:", err)
		return errorHandler.SendCustomError(err)
	}

	return errorHandler.SendData(fiber.StatusCreated, post)
}

// @Tags			BlogPost
// @Summary		Get all blog posts
// @Description	Retrieve all blog posts with optional limit and skip parameters
// @Accept			json
// @Produce		json
// @Param			limit	query		int	false	"Limit the number of blog posts returned"
// @Param			skip	query		int	false	"Skip the first N blog posts"
// @Success		200		{array}		blogpost.BlogPost
// @Failure		500		{object}	response.ErrorResponse
// @Router			/api/blog-post [get]
func (handler *Handler) GetAll(c *fiber.Ctx) error {
	errorHandler := response.NewErrorHandler(c, "Blog Post", "Blog Posts")
	limit := c.QueryInt("limit", 10)
	skip := c.QueryInt("skip", 0)
	posts, err := handler.usecase.GetAll(&limit, &skip)
	if err != nil {
		log.Println("Error fetching blog posts:", err)
		return errorHandler.SendCustomError(err)
	}

	return errorHandler.SendData(fiber.StatusOK, posts)
}

// @Tags			BlogPost
// @Summary		Get a blog post by ID
// @Description	Retrieve a blog post by its ID
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Blog Post ID"
// @Success		200	{object}	blogpost.BlogPost
// @Failure		400	{object}	response.ErrorResponse
// @Failure		404	{object}	response.ErrorResponse
// @Failure		500	{object}	response.ErrorResponse
// @Router			/api/blog-post/{id} [get]
func (handler *Handler) Get(c *fiber.Ctx) error {
	errorHandler := response.NewErrorHandler(c, "Blog Post", "Blog Posts")
	id := c.Params("id")
	if id == "" {
		return errorHandler.SendError(fiber.StatusBadRequest, "Blog post ID is required")
	}

	post, err := handler.usecase.Get(id)
	if err != nil {
		log.Println("Error fetching blog post:", err)
		return errorHandler.SendCustomError(err)
	}

	if post == nil {
		return errorHandler.SendError(fiber.StatusNotFound, "Blog post not found")
	}

	return errorHandler.SendData(fiber.StatusOK, post)
}

// @Tags			BlogPost
// @Summary		Delete a blog post by ID
// @Description	Delete a blog post by its ID
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Blog Post ID"
// @Success		200	{object}	response.SuccessResponse
// @Failure		400	{object}	response.ErrorResponse
// @Failure		404	{object}	response.ErrorResponse
// @Failure		500	{object}	response.ErrorResponse
// @Router			/api/blog-post/{id} [delete]
func (handler *Handler) Delete(c *fiber.Ctx) error {
	errorHandler := response.NewErrorHandler(c, "Blog Post", "Blog Posts")
	id := c.Params("id")
	if id == "" {
		return errorHandler.SendError(fiber.StatusBadRequest, "Blog post ID is required")
	}

	if err := handler.usecase.Delete(id); err != nil {
		log.Println("Error deleting blog post:", err)
		return errorHandler.SendCustomError(err)
	}

	return errorHandler.SendSuccess(fiber.StatusOK, "Blog post deleted successfully")
}

// @Tags			BlogPost
// @Summary		Update a blog post by ID
// @Description	Update a blog post by its ID with new title, body, author, and descriptions
// @Accept			json
// @Produce		json
// @Param			id			path		string			true	"Blog Post ID"
// @Param			blogPost	body		UpdateBlogPost	true	"Blog Post"
// @Success		200			{object}	blogpost.BlogPost
// @Failure		400			{object}	response.ErrorResponse
// @Failure		404			{object}	response.ErrorResponse
// @Failure		409			{object}	response.ErrorResponse
// @Failure		500			{object}	response.ErrorResponse
// @Router			/api/blog-post/{id} [patch]
func (handler *Handler) Update(c *fiber.Ctx) error {
	errorHandler := response.NewErrorHandler(c, "Blog Post", "Blog Posts")
	id := c.Params("id")
	if id == "" {
		return errorHandler.SendError(fiber.StatusBadRequest, "Blog post ID is required")
	}
	var input UpdateBlogPost
	if err := c.BodyParser(&input); err != nil {
		return errorHandler.SendError(fiber.StatusBadRequest, "Invalid input data")
	}
	err := handler.usecase.Update(id, input.Title, input.Body, input.Author, input.Descriptions, nil)
	if err != nil {
		log.Println("Error updating blog post:", err)
		return errorHandler.SendCustomError(err)
	}

	return errorHandler.SendData(fiber.StatusOK, input)
}
