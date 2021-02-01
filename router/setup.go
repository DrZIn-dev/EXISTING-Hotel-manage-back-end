package router

import "github.com/gofiber/fiber/v2"

var USER fiber.Router

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello from api")
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", hello)

	USER = api.Group("/user")
	SetupUserRoutes()
}
