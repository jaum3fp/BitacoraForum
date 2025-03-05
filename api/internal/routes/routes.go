package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/handlers"
)


func SetUpRoutes(router fiber.Router) {

	// Posts routes
	router.Get("/posts", handlers.GetAllPosts)
	router.Post("/post", handlers.CreatePost)

	router.Get("/post/:id", handlers.GetPost)
	router.Put("/post/:id", handlers.UpdatePost)
	router.Delete("/post/:id", handlers.DeletePost)

	// Countrys routes
	router.Get("/country/:flag/posts", handlers.GetPostCountrys)

}
