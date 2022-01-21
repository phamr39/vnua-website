package controllers

import (
	"fis/be/models"
	repo "fis/be/repositories"
	"fis/be/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetCustomers(c *fiber.Ctx) error {
	page := c.FormValue("_page")
	start, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	start = (start - 1) * 20
	customers, err := repo.GetAllCustomers(start)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	total, err := repo.GetTotalCustomers()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total" : total,
		"customers" : customers,
	})
}

func CreateCustomerInfo(c *fiber.Ctx) error {
	var customer models.Customer
	err := c.BodyParser(&customer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	customer, err = repo.CreateCustomerInfo(customer)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "success",
		"customer" : customer,
	})
}

func ExportCustomers(c *fiber.Ctx) error {
	filePath := "./csv/customers.csv"
	if err := utils.CreateFile(filePath); err != nil {
		return err
	}
	return c.Download(filePath)
}

