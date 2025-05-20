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

	db := db.Connect()

	app := fiber.New(fiber.Config{
		AppName:               "Bitacora Forum REST API v0.0.1",
		BodyLimit:             4 * 1024 * 1024,
		CaseSensitive:         true,
		StrictRouting:         true,
		DisableStartupMessage: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if e, ok := err.(*fiber.Error); ok {
				return c.Status(e.Code).JSON(fiber.Map{"error": e.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://127.0.0.1:3000,https://nuxt-client:3000,https://localhost:3000,https://bitacoraforum.es:3000",
		AllowCredentials: true,
	}))
	app.Use(logger.New())

	router := app.Group("/api/v1")
	routes.SetUpRoutes(router, db)

	app.Get("/health", func (c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Fatal(app.ListenTLS(":8080", "/certs/fiber-api.pem", "/certs/fiber-api-key.pem"))
}
