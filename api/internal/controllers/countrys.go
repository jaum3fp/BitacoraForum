package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

type CountryController struct {
	CountryRepo repositorys.CountryRepository
}

func NewCountryHandler(countryRepo repositorys.CountryRepository) *CountryController {
	return &CountryController{countryRepo}
}

func (h *CountryController) GetCountry(c *fiber.Ctx) error {

	countryFlag := c.Params("flag")

	country, err := h.CountryRepo.GetCountry(countryFlag)

	if err != nil {
		return dbErrorHandler(c, err, "Country not found")
	}

	return c.JSON(country)
}
