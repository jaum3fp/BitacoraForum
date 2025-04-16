package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

type AuthController struct {
	AuthRepo repositorys.AuthRepository
}

func NewAuthHandler(authRepo repositorys.AuthRepository) *AuthController {
	return &AuthController{authRepo}
}

func (h *AuthController) Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	tkn, err := h.AuthRepo.Login(username, password)
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
	username := c.FormValue("username")
	password := c.FormValue("password")

	tkn, err := h.AuthRepo.Register(username, password)
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
