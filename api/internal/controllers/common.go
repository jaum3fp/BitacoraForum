package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func dbErrorHandler(c *fiber.Ctx, err error, notFoundMsg string) error {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": notFoundMsg})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}
