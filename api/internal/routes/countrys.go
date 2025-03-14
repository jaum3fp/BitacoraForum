package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/controllers"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

func setCountryRoutes(router fiber.Router, repo repositorys.CountryRepository) {

	countryHandler := controllers.NewCountryHandler(repo)

	group := router.Group("/country")

	group.Get("/:flag", countryHandler.GetCountry)

}
