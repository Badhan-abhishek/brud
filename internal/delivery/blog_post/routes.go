package blogpost

import (
	blogpostRepo "be.blog/internal/infra/repository/blog_post"
	blogpostUsecase "be.blog/internal/usecase/blog_post"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "be.blog/docs"
)

func RegisterBlogPostRoutes(app *fiber.App, db *pgxpool.Pool) {
	repo := blogpostRepo.NewBlogPostRepo(db)
	usecase := blogpostUsecase.NewBlogPostUsecase(repo)
	handler := NewBlogPostHandler(usecase)

	r := app.Group("/api/blog-post")

	handler.RegisterRoutes(r)
}
