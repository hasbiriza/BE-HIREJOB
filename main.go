package main

import (
	"be_hiring_app/src/config"
	"be_hiring_app/src/helper"
	"be_hiring_app/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type",
	}))

	config.InitDB()
	helper.Migration()
	routes.Router(app)

	app.Listen(":8080")
}
