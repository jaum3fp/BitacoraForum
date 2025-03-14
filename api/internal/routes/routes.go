package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
	"gorm.io/gorm"
)

func SetUpRoutes(router fiber.Router, db *gorm.DB) {

	postRepository := repositorys.NewPostRepository(db)
	countryRepository := repositorys.NewCountryRepository(db)

	setCountryRoutes(router, countryRepository)
	setPostRoutes(router, postRepository)

}
