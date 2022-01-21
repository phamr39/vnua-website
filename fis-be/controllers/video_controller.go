package controllers

import (
	"fis/be/models"
	repo "fis/be/repositories"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetVideos(c *fiber.Ctx) error {
	page := c.FormValue("_page")
	key := c.FormValue("q")
	start, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	start = (start - 1) * 4
	key = "%" + key + "%"
	videos, total, err := repo.GetVideos(key, start)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"videos" : videos,
		"total" : total,
	})
}

func GetVideoByName(c *fiber.Ctx) error {
	name := c.Params("name")
	video, err := repo.GetVideoByName(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"video" : video,
	})
}

func CreateVideo(c *fiber.Ctx) error {
	var video models.Video
	err := c.BodyParser(&video)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	video, err = repo.CreateVideo(video)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "success",
		"video" : video,
	})
}

func UpdateVideoById(c *fiber.Ctx) error {
	var video models.Video
	err := c.BodyParser(&video)
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

	video.ID = int64(id)
	err = repo.UpdateVideoByID(video)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "success",
		"video" : video,
	})
}

func DeleteVideo(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = repo.DeleteVideo(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted",
	})
}

func GetNewestVideo(c *fiber.Ctx) error{
	vid, err := repo.GetNewestVideo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(vid)
}