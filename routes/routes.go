package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/syahlan1/ecommerce-scrapper.git/controllers"
	"github.com/syahlan1/ecommerce-scrapper.git/middleware"

)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://persentor.online, http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", middleware.AuthRequired, controllers.GetUser)
	app.Post("/api/logout", controllers.Logout)

	app.Post("/api/get-search", controllers.GetFromSearch)
	app.Get("/api/product/:id", controllers.GetProductById)
	app.Get("/api/products/recommend", middleware.AuthRequired, controllers.RecommendProducts)
	app.Get("/api/products/last-viewed", middleware.AuthRequired, controllers.GetLastViewedProducts)
	app.Get("/api/get-story", controllers.GetStory)
	app.Get("/api/get-all-product", controllers.GetProduct)

}
