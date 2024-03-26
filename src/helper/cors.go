package helper

import "github.com/gofiber/fiber/v2"

func EnableCors(c *fiber.Ctx) {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	c.Set("Content-Security-Policy", "default-src 'self'")
}
