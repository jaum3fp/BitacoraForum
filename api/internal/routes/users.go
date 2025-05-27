package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/controllers"
	"github.com/jaum3fp/bitacora-forum/internal/middleware"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

func setUserRoutes(router fiber.Router, repo repositorys.UserRepository) {

	userHandler := controllers.NewUserHandler(repo)

	group := router.Group("/user")

	group.Get("/all", userHandler.GetAllUsers)
	group.Get("/me", middleware.Protected(), userHandler.GetUserOwner)
	group.Get("/:username", userHandler.GetUser)

	//group.Patch("/me/:field", nil)
	group.Post("/", userHandler.CreateUser)
	group.Put("/me", middleware.Protected(), userHandler.UpdateUser)
	group.Delete("/me", middleware.Protected(), userHandler.DeleteUser)

	//group.Patch("/ban/:id", userHandler.BanUser)
}
