package controllers

import (
	repo "fis/be/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(c *fiber.Ctx) error {
	root := c.FormValue("root")
	categories, err := repo.GetAllCategories(root)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}
