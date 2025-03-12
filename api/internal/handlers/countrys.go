package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/db"
	"github.com/jaum3fp/bitacora-forum/internal/models"
)

func GetPostCountrys(c *fiber.Ctx) error {
	countryFlag := c.Params("flag")
	var totalPosts int64
	db.GetInstance().Model(&models.Post{}).Joins("JOIN countries ON posts.country_id = countries.id").Where("countries.flag = ?", countryFlag).Count(&totalPosts)
	return c.JSON(fiber.Map{ "total": totalPosts })
}

func GetCountry(c *fiber.Ctx) error {
	type ThisResponse struct {
		Name 		string 	`json:"name"`
		TotalPosts 	int64 	`json:"total_posts"`
	}
	countryFlag := c.Params("flag")
	var data ThisResponse
	res := db.GetInstance().Model(&models.Country{}).
		Select("countries.name, COUNT(posts.id) AS total_posts").
		Joins("LEFT JOIN posts ON posts.country_id = countries.id").
		Where("countries.flag = ?", countryFlag).
		Group("countries.id").
		First(&data)
	
	if err := dbErrorHandler(c, res, "Country not found"); err != nil {
		return err
	}

	return c.JSON(data)
}
