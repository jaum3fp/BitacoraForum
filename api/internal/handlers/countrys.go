package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/db"
	"github.com/jaum3fp/bitacora-forum/internal/models"
)

func GetPostCountrys(c *fiber.Ctx) error {
	postId := c.Params("id")
	var totalPosts int64
	db.GetInstance().Model(&models.Post{}).Where("country_id = ?", postId).Count(&totalPosts)
	return c.JSON(fiber.Map{ "total": totalPosts })
}
