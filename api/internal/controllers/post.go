package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jaum3fp/bitacora-forum/internal/models"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

type PostController struct {
	PostRepo repositorys.PostRepository
}

func NewPostHandler(postRepo repositorys.PostRepository) *PostController {
	return &PostController{postRepo}
}


func (h *PostController) GetAllPosts(c *fiber.Ctx) error {

	filters := c.Queries()
	data, err := h.PostRepo.GetAllPosts(filters)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *PostController) GetCountryPosts(c *fiber.Ctx) error {
	flag := c.Params("cc")
	posts, err := h.PostRepo.GetCountryPosts(flag)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}

func (h *PostController) GetCountryPostsNumber(c *fiber.Ctx) error {
	flag := c.Params("flag")
	posts, err := h.PostRepo.GetCountryPostsNumber(flag)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}

func (h *PostController) CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post parse fail: " + err.Error(),
		})
	}
	if err := h.PostRepo.CreatePost(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(post.ID)
}

func (h *PostController) GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := h.PostRepo.GetPost(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(post)
}

func (h *PostController) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parse fail",
		})
	}
	if err := h.PostRepo.UpdatePost(id, post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(post.ID)
}

func (h *PostController) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.PostRepo.DeletePost(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"ID": id})
}

func (h *PostController) IncrementPostViews(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.PostRepo.IncrementPostViews(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"ID": id})
}

func (h *PostController) GetPostsByTag(c *fiber.Ctx) error {
	id := c.Params("id")

	posts, err := h.PostRepo.GetPostsByTag(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}
