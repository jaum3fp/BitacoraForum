package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func dbErrorHandler(c *fiber.Ctx, result *gorm.DB, notFoundMsg string) error {
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": notFoundMsg})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": notFoundMsg})
	}
	return nil
}