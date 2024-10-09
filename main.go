package main

import (
	"log"
	"os"

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
	port := os.Getenv("PORT")
    if port == "" {
        // Default ke port 8080 jika tidak disetel oleh Heroku
        port = "8080"
        log.Printf("INFO: No PORT environment variable detected, defaulting to %s", port)
    }
}
