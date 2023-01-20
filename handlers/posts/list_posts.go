package posts

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Oldbury/thefeed/models"
	"github.com/Oldbury/thefeed/utils"
	"github.com/gofiber/fiber/v2"
)

func ListPosts(c *fiber.Ctx) error {

	ctx := context.Background()

	rdb := utils.GetRedis()

	// get post from redis
	posts, err := rdb.LRange(ctx, "posts", 0, -1).Result()

	if err != nil {
		return c.SendStatus(404)
	}

	var postsList []models.Post

	for _, post := range posts {
		var postModel models.Post
		err := json.Unmarshal([]byte(post), &postModel)
		if err != nil {
			log.Println(err)
		}
		postsList = append(postsList, postModel)
	}

	return c.Status(200).JSON(postsList)

}
