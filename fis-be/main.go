package main

import (
	"fis/be/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	v1 := app.Group("api/v1")
	routers.AuthConfig(&v1)
	routers.ProductsConfig(&v1)
	routers.CategoriesConfig(&v1)
	routers.NewsConfig(&v1)
	routers.VideosConfig(&v1)
	routers.UploadConfig(&v1)
	routers.CustomersConfig(&v1)
	routers.UsersConfig(&v1)
	routers.StoreConfig(&v1)
	routers.AdminConfig(&v1)


	app.Listen(":3000")
}
