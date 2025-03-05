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
