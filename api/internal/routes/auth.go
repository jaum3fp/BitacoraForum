package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/controllers"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

func setAuthenticationRoutes(router fiber.Router, repo repositorys.AuthRepository) {

	authHandler := controllers.NewAuthHandler(repo)

	group := router.Group("/auth")

	group.Post("/login", authHandler.Login)
	group.Post("/register", authHandler.Register)
	group.Get("/logout", authHandler.Logout)
	// router.Get("/refresh-token", authHandler.RefreshToken)
}
