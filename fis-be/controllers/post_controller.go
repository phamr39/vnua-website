package controllers

import (
	"fis/be/models"
	repo "fis/be/repositories"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

const (
	postTypeNews  = "news"
	postTypeEvent = "event"
)

func GetAllNews(c *fiber.Ctx) error {
	topic := c.FormValue("topic_id")
	orders := c.FormValue("_order")
	_page := c.FormValue("_page")
	page, _ := strconv.Atoi(_page)
	page = (page - 1) * 4
	title := c.FormValue("q")
	title = "%" + title + "%"
	news, err := repo.GetAllPosts(topic, orders, page, title, postTypeNews)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	total, err := repo.GetTotalPostRequest(topic, title, postTypeNews)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total": total,
		"news":  news,
	})
}

func GetAllEvents(c *fiber.Ctx) error {
	topic := c.FormValue("topic_id")
	orders := c.FormValue("_order")
	_page := c.FormValue("_page")
	page, _ := strconv.Atoi(_page)
	page = (page - 1) * 4
	title := c.FormValue("q")
	title = "%" + title + "%"
	events, err := repo.GetAllPosts(topic, orders, page, title, postTypeEvent)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	total, err := repo.GetTotalPostRequest(topic, title, postTypeEvent)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total": total,
		"events":  events,
	})
}

func CreateNews(c *fiber.Ctx) error {
	var news models.Post
	err := c.BodyParser(&news)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	post, err := repo.CreatePost(news)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

func GetNewsByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(repo.GetPostById(int64(id)))
}

func GetListNewestNews(c *fiber.Ctx) error {
	postType := c.FormValue("post_type")
	return c.Status(fiber.StatusOK).JSON(repo.GetNewestPosts(postType))
}

func GetRandomPost(c *fiber.Ctx) error {
	postType := c.FormValue("post_type")
	return c.Status(fiber.StatusOK).JSON(repo.GetRandomPosts(postType))
}

func GetTopics(c *fiber.Ctx) error{
	return c.Status(fiber.StatusOK).JSON(repo.GetTopics())
}

func DeletePostById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = repo.DeletePostById(int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted",
	})
}