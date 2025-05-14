package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/dtos"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

type AuthController struct {
	AuthRepo repositorys.AuthRepository
}

func NewAuthHandler(authRepo repositorys.AuthRepository) *AuthController {
	return &AuthController{authRepo}
}

func (h *AuthController) Login(c *fiber.Ctx) error {
	var userCredentials dtos.UserDTO
	if err := c.BodyParser(&userCredentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post parse fail: " + err.Error(),
		})
	}

	tkn, err := h.AuthRepo.Login(userCredentials.Username, userCredentials.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token": tkn,
	})
}

func (h *AuthController) Register(c *fiber.Ctx) error {
	var user dtos.UserDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post parse fail: " + err.Error(),
		})
	}

	tkn, err := h.AuthRepo.Register(user)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": tkn,
	})
}

func (h *AuthController) Logout(c *fiber.Ctx) error {
	return nil
}
