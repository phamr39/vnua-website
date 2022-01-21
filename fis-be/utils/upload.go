package utils

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["file"]

	for _, file := range files {

		if err := c.SaveFile(file, "./uploads/"+ strconv.Itoa(int(time.Now().Unix())) + ".jpg") ; err != nil {
			return err
		}
	}
	return c.JSON(fiber.Map{
		"url": "http://vnuapharma.com.vn:3000/api/v1/uploads/" + strconv.Itoa(int(time.Now().Unix())) + ".jpg",
	})
}
