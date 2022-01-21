package controllers

import (
	"fis/be/models"
	repo "fis/be/repositories"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CreateNewUser(c *fiber.Ctx) error{
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if data["password"] != data["password_confirm"] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "password not match",
		})
	}
	user := models.User{
		Name:        data["name"],
		PhoneNumber: data["phone_number"],
		Email:       data["email"],
	}
	user.SetPassword(data["password"])

	err = repo.CreateNewUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "success",
	})
}

func GetAllUsers(c *fiber.Ctx) error{
	page := c.FormValue("_page")
	start, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	start = (start - 1) * 20
	users := repo.GetAllUsers(start)
	total, err := repo.GetTotalUsers()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total" : total,
		"users" : users,
	})
}

func GetUserById(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user := repo.GetUserById(int64(id))
	return c.Status(fiber.StatusOK).JSON(user)
}

func GetUserByEmail(c *fiber.Ctx) error{
	email := c.FormValue("email")
	user, err := repo.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUserById(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var user models.User
	user.ID = int64(id)

	err = c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	repo.UpdateUserById(int64(id), user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "updated",
	})
}

func DeleteUserById(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	repo.DeleteUserById(int64(id))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "deleted",
	})
}

