package ong

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"hello": "world"})
}
