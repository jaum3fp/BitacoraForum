package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jaum3fp/bitacora-forum/internal/db"
	"github.com/jaum3fp/bitacora-forum/internal/models"
)


func GetAllPosts(c *fiber.Ctx) error {

	type ThisResponse struct {
		models.Post
		OwnerUsername string `json:"owner_username"`
	}

	var data []ThisResponse
	res := db.GetInstance().Model(&models.Post{}).
		Select("posts.*, users.username AS owner_username").
		Joins("JOIN users ON posts.owner_id = users.id").
		Find(&data)

	if err := dbErrorHandler(c, res, "No posts found"); err != nil {
		return err
	}

	return c.JSON(data)
}

func CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Post parse fail: " + err.Error(),
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
	res := db.GetInstance().Find(&post, postId)

	if err := dbErrorHandler(c, res, "Post not found"); err != nil {
		return err
	}

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
	res := db.GetInstance().Model(&models.Post{}).Where("id = ?", postId).Updates(post)

	if err := dbErrorHandler(c, res, "Post not found or not updated"); err != nil {
		return err
	}

	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	postId := c.Params("id")
	var post models.Post
	res := db.GetInstance().Delete(&models.Post{}, postId)
	
	if err := dbErrorHandler(c, res, "Post not found or not deleted"); err != nil {
		return err
	}

	return c.JSON(post)
}
