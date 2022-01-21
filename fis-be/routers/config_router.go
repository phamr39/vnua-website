package routers

import (
	"fis/be/controllers"
	"fis/be/middlewares"
	"fis/be/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthConfig(router *fiber.Router) {
	(*router).Post("/login", controllers.Login)
	(*router).Post("/logout", controllers.Logout)
}

func UploadConfig(router *fiber.Router) {
	(*router).Static("/uploads", "./uploads")
}

func ProductsConfig(router *fiber.Router) {
	(*router).Get("/products", controllers.GetAllProducts)
	(*router).Get("/new-products", controllers.GetNewProducts)
	(*router).Get("/random-products", controllers.GetRandomProducts)
	(*router).Get("/products/:id", controllers.GetProductById)
	(*router).Get("/products/name/:name", controllers.GetProductByName)
	(*router).Get("/products/handle/:handle", controllers.GetProductByHandle)
	(*router).Get("/products/categories/:category", controllers.GetProductsByCategory)
}

func CategoriesConfig(router *fiber.Router) {
	(*router).Get("/categories", controllers.GetAllCategories)
}

func VideosConfig(router *fiber.Router){
	(*router).Get("/videos", controllers.GetVideos)
	(*router).Get("/videos/new", controllers.GetNewestVideo)
}

func NewsConfig(router *fiber.Router) {
	(*router).Get("/news", controllers.GetAllNews)
	(*router).Get("/events", controllers.GetAllEvents)
	(*router).Get("/posts/:id", controllers.GetNewsByID)
	(*router).Get("/new-posts", controllers.GetListNewestNews)
	(*router).Get("/topics", controllers.GetTopics)
	(*router).Get("/random-posts", controllers.GetRandomPost)
}

func StoreConfig(router *fiber.Router){
	(*router).Post("/stores", controllers.CreateStore)
}

func CustomersConfig(router *fiber.Router) {
	(*router).Post("/customers", controllers.CreateCustomerInfo)
}

func UsersConfig(router *fiber.Router) {
	(*router).Use(middlewares.CommonMiddleWare)
	(*router).Use(middlewares.IsAuthenticated)
	(*router).Get("/users", controllers.GetAllUsers)
	(*router).Post("/users", controllers.CreateNewUser)
	(*router).Get("/users/:id", controllers.GetUserById)
	(*router).Get("/users/email/:email", controllers.GetUserByEmail)
	(*router).Put("/users/:id", controllers.UpdateUserById)
	(*router).Delete("/users/:id", controllers.DeleteUserById)
}

func AdminConfig(router *fiber.Router) {
	(*router).Use(middlewares.CommonMiddleWare)
	(*router).Use(middlewares.IsAuthenticated)

	(*router).Post("/upload", utils.Upload)

	(*router).Put("/profiles", controllers.UpdatePassword)
	(*router).Get("/profiles", controllers.GetUserInfo)
	(*router).Put("/update-info", controllers.UpdateInfo)

	(*router).Post("/products", controllers.CreateNewProduct)
	(*router).Put("/products/:id", controllers.UpdateProductById)
	(*router).Delete("/products/:id", controllers.DeleteProductById)

	(*router).Post("/videos", controllers.CreateVideo)
	(*router).Put("/videos/:id", controllers.UpdateVideoById)
	(*router).Delete("/videos/:id", controllers.DeleteVideo)
	(*router).Get("/videos/:name", controllers.GetVideoByName)

	(*router).Post("/posts", controllers.CreateNews)
	(*router).Delete("/posts/:id", controllers.DeletePostById)

	(*router).Get("/stores", controllers.GetStores)
	(*router).Get("/stores/export", controllers.ExportStores)

	(*router).Get("/customers/export", controllers.ExportCustomers)
	(*router).Get("/customers", controllers.GetCustomers)
}