package main

import (
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/database"
	"github.com/DrZIn-dev/EXISTING-Hotel-manage-back-end/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
)

func CreateServer() *fiber.App {
	app := fiber.New()
	return app
}

func main() {
	database.ConnectDB()
	app := CreateServer()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
