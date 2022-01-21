package controllers

import (
	"fis/be/models"
	repo "fis/be/repositories"
	"fis/be/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CreateStore(c *fiber.Ctx) error{
	var store models.Store
	err := c.BodyParser(&store)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	store, err = repo.CreateStore(store)
	return c.Status(fiber.StatusOK).JSON(store)
}

func UpdateStore(c *fiber.Ctx) error{
	var store models.Store
	err := c.BodyParser(&store)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	store.ID = int64(id)
	store, err = repo.UpdateStore(store)
	return c.Status(fiber.StatusOK).JSON(store)
}

func GetStores(c *fiber.Ctx) error{
	order := c.FormValue("_order")
	key := c.FormValue("q")
	key = "%" + key + "%"
	start, err := strconv.Atoi(c.FormValue("_page"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	start = (start - 1) * 10
	if order != "asc"{
		order = "desc"
	}

	stores, total, err := repo.GetStores(start, key, order)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total" : total,
		"stores" : stores,
	})
}

func DeleteStore(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = repo.DeleteStore(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func GetStoreByID(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	store, err := repo.GetStoreByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(store)
}

func ExportStores(c *fiber.Ctx) error {
	filePath := "./csv/customers.csv"
	if err := utils.CreateCSVStores(filePath); err != nil {
		return err
	}
	return c.Download(filePath)
}
