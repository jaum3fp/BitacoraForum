package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
	"gorm.io/gorm"
)

func SetUpRoutes(router fiber.Router, db *gorm.DB) {

	postRepository := repositorys.NewPostRepository(db)
	userRepository := repositorys.NewUserRepository(db)
	authRepository := repositorys.NewAuthRepository(db, userRepository)

	setPostRoutes(router, postRepository)
	setUserRoutes(router, userRepository)
	setAuthenticationRoutes(router, authRepository)
}
