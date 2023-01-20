package posts

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Oldbury/thefeed/utils"
	"github.com/gofiber/fiber/v2"
)

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func AddPost(c *fiber.Ctx) error {

	c.Accepts("application/json")
	ctx := context.Background()

	newPost := new(Post)

	err := c.BodyParser(newPost)

	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	rdb := utils.GetRedis()

	marshalPost, err := json.Marshal(newPost)

	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	err = rdb.LPush(ctx, "posts", marshalPost).Err()

	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	return c.Status(201).JSON(newPost)
}
