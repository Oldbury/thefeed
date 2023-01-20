package main

import (
	"log"

	"github.com/Oldbury/thefeed/handlers/posts"
	"github.com/Oldbury/thefeed/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// initialise redis connection
	utils.InitRedis()

	api := app.Group("/api")
	api.Get("/post/:id", posts.GetPost)
	api.Get("/posts", posts.ListPosts)
	api.Post("/post", posts.AddPost)

	log.Fatal(app.Listen(":3000"))
}
