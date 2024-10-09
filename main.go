package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/syahlan1/ecommerce-scrapper.git/database"
	"github.com/syahlan1/ecommerce-scrapper.git/routes"

)

func main() {
	app := fiber.New()

	// Initialize database connection
	database.Connect()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
