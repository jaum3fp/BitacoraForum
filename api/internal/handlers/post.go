package handlers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/db"
	"github.com/jaum3fp/bitacora-forum/internal/models"
)


func GetAllPosts(c *fiber.Ctx) error {
	var posts []models.Post
	db.GetInstance().Find(&posts)
	return c.JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Parse fail",
		})
	}
	if err := db.GetInstance().Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(post)
}

func GetPost(c *fiber.Ctx) error {
	postId := c.Params("id")
	var post models.Post
	db.GetInstance().Find(&post, postId)
	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	postId := c.Params("id")
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Parse fail",
		})
	}
	if err := db.GetInstance().Model(&models.Post{}).Where("id = ?", postId).Updates(post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	postId := c.Params("id")
	var post models.Post
	if err := db.GetInstance().Delete(&models.Post{}, postId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(post)
}
