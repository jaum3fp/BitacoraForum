package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/jaum3fp/bitacora-forum/internal/dtos"
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

	totals, err := h.PostRepo.GetTotalPosts(filters)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": data, "total": totals})
}

func (h *PostController) GetCountryPosts(c *fiber.Ctx) error {
	flag := c.Params("cc")
	posts, err := h.PostRepo.GetCountryPosts(flag)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}

func (h *PostController) GetPostCommentsNumber(c *fiber.Ctx) error {
	super := c.Params("super")
	posts, err := h.PostRepo.GetPostCommentsNumber(super)

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

func (h *PostController) GetPostComments(c *fiber.Ctx) error {
	super := c.Params("super")
	posts, err := h.PostRepo.GetPostComments(super)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(posts)
}

func (h *PostController) CreatePost(c *fiber.Ctx) error {
	var post dtos.PostDTO
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post parse fail: " + err.Error(),
		})
	}
	if err := h.PostRepo.CreatePost(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true, "id": post.ID})
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

	post, err := h.PostRepo.GetPost(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var updatedPost dtos.PostUsernameDTO
	if err := c.BodyParser(&updatedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parse fail",
		})
	}

	aTkn := c.Cookies("access-token")
	if aTkn == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unautorized"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username, ok := claims["username"].(string)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No se ha podido identificar al usuario"})
	}

	if username != post.OwnerUsername {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "El usuario no puede realizar esta acción"})
	}

	if err := h.PostRepo.UpdatePost(id, updatedPost); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"id": updatedPost.ID, "success": true})
}

func (h *PostController) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	post, err := h.PostRepo.GetPost(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	aTkn := c.Cookies("access-token")
	if aTkn == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unautorized"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username, ok := claims["username"].(string)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No se ha podido identificar al usuario"})
	}

	if username != post.OwnerUsername {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "El usuario no puede realizar esta acción"})
	}

	if err := h.PostRepo.DeletePost(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"id": id, "success": true})
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
