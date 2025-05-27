package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jaum3fp/bitacora-forum/internal/dtos"
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

func (h *UserController) GetUserOwner(c *fiber.Ctx) error {

	aTkn := c.Cookies("access-token")
	if aTkn == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unautorized"})
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username, ok := claims["username"].(string)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Values not found"})
	}

	data, err := h.UserRepo.GetUser(username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// No debería ocurrir pero...
	if data.Password != "" {
		data.Password = ""
	}

	return c.JSON(data.ParseWithoutPassword())
}

func (h *UserController) GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	data, err := h.UserRepo.GetUser(username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if data.Password != "" {
		data.Password = ""
	}

	return c.JSON(data.ParseWithoutPassword())
}

func (h *UserController) CreateUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserController) UpdateUser(c *fiber.Ctx) error {
	aTkn := c.Cookies("access-token")
	if aTkn == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unautorized"})
	}
	userLocal := c.Locals("user").(*jwt.Token)
	claims := userLocal.Claims.(jwt.MapClaims)
	username, ok := claims["username"].(string)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Values not found"})
	}

	var user dtos.UserWithourPasswordDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parse fail",
		})
	}
	if err := h.UserRepo.UpdateUser(username, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user.ID)
}

func (h *UserController) DeleteUser(c *fiber.Ctx) error {
	return nil
}
