package controllers

import (
	"time"

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

	cookie := getAccessCookie(tkn, false)
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"success": true,
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

	cookie := getAccessCookie(tkn, false)
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func (h *AuthController) Logout(c *fiber.Ctx) error {

	c.Cookie(getAccessCookie("", true))

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func getAccessCookie(tkn string, expired bool) (cookie *fiber.Cookie) {

	cookie = &fiber.Cookie{
		Name: "access-token",
		Value: tkn,
		Domain: "bitacoraforum.es",
		Path: "/",
		Expires: time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure: true,
		SameSite: fiber.CookieSameSiteNoneMode,
	}

	if (expired) {
		cookie.Expires = time.Unix(0, 0)
	}

	return
}
