package controllers

import (
	"fis/be/models"
	repo "fis/be/repositories"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllProducts(c *fiber.Ctx) error {
	cat := c.FormValue("category_id")
	orders := c.FormValue("_order")
	_page := c.FormValue("_page")
	page, _ := strconv.Atoi(_page)
	page = (page - 1) * 12
	name := c.FormValue("q")
	name = "%" + name + "%"
	if orders != "desc"{
		orders = "asc"
	}
	products, err := repo.GetAllProducts(cat, orders, page, name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	total, err := repo.GetTotalProductByRequest(cat, name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total":    total,
		"products": products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(repo.GetProductByID(int64(id)))
}

func GetProductByName(c *fiber.Ctx) error {
	name := c.Params("name")
	return c.Status(fiber.StatusOK).JSON(repo.GetProductByName(name))
}

func GetProductByHandle(c *fiber.Ctx) error {
	handle := c.Params("handle")
	return c.Status(fiber.StatusOK).JSON(repo.GetProductByHandle(handle))
}

func GetNewProducts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(repo.GetNewProducts())
}

func GetRandomProducts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(repo.GetRandomProducts())
}

func GetProductsByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	return c.Status(fiber.StatusOK).JSON(repo.GetProductsByCategory(category))
}

func UpdateProductById(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
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

	product.ID = int64(id)
	err = repo.UpdateProductByID(product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(repo.GetProductByID(product.ID))
}

func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = repo.DeleteProduct(int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted",
	})
}

func CreateNewProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(repo.CreateProduct(product))
}
