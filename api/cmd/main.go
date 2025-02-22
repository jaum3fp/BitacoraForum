package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jaum3fp/bitacora-forum/internal/db"
	"github.com/jaum3fp/bitacora-forum/internal/routes"
)


func main() {

	db.Init()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	router := app.Group("/api/v1")
	routes.SetUpRoutes(router)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Fatal(app.Listen(":8080"))
}