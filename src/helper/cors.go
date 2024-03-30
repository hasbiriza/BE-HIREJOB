package helper

import "github.com/gofiber/fiber/v2"

func EnableCors(c *fiber.Ctx) {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST,PUT,DELETE, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type")
	c.Set("Content-Security-Policy", "default-src 'self'")
}
