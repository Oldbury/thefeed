package posts

import (
	"context"

	"github.com/Oldbury/thefeed/utils"
	"github.com/gofiber/fiber/v2"
)

func GetPost(c *fiber.Ctx) error {
	ctx := context.Background()

	rdb := utils.GetRedis()

	// get post from redis
	post, err := rdb.Get(ctx, c.Params("id")).Result()

	if err != nil {
		return c.SendStatus(404)
	}

	return c.SendString(post)

}
