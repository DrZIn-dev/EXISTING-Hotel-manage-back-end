package router

import "github.com/gofiber/fiber/v2"

var USER fiber.Router
var HOTEL fiber.Router
var TICKET fiber.Router

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello from api")
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", hello)

	USER = api.Group("/user")
	HOTEL = api.Group("/hotel")
	TICKET = api.Group("/ticket")
	SetupUserRoutes()
	SetupHotelRoutes()
	SetupTicketRoute()
}
