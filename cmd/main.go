package main

import (
	"log"

	"be.blog/configs"
	_ "be.blog/docs"
	blogpost "be.blog/internal/delivery/blog_post"
	"be.blog/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title			BRUD
// @version		1.0
// @description	BRUD, a Blog CRUD application.
// @host		brud.onrender.com
// @BasePath		/
func main() {
	app := fiber.New()

	// For development purposes, enable CORS for all origins.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	cfg := configs.Load()

	app.Get("/swagger/*", swagger.HandlerDefault)

	db := database.ConnectPG(cfg)

	blogpost.RegisterBlogPostRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
