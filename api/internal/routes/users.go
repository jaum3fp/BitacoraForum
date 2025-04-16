package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/controllers"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

func setUserRoutes(router fiber.Router, repo repositorys.UserRepository) {

	userHandler := controllers.NewUserHandler(repo)

	group := router.Group("/user")

	group.Get("/all", userHandler.GetAllUsers)
	group.Post("/", userHandler.CreateUser)

	group.Get("/:id", userHandler.GetUser)
	group.Put("/:id", userHandler.UpdateUser)
	group.Delete("/:id", userHandler.DeleteUser)

	//group.Patch("/ban/:id", userHandler.BanUser)
}
