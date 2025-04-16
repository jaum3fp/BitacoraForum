package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaum3fp/bitacora-forum/internal/repositorys"
)

type UserController struct {
	UserRepo repositorys.UserRepository
}

func NewUserHandler(userRepo repositorys.UserRepository) *UserController {
	return &UserController{userRepo}
}

func (h *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserRepo.GetAllUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func (h *UserController) GetUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserController) CreateUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserController) UpdateUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserController) DeleteUser(c *fiber.Ctx) error {
	return nil
}
