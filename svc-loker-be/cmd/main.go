package main

import (
	"log"
	"svc-loker-be/internal/database"
	"svc-loker-be/internal/routes"
	"svc-loker-be/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize database
	database.Init()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "svc-loker",
		AppName:       "service backend loker",
	})

	// Initialize routes from the "routes" package
	routes.Setup(app)

	if err := server.Start(app); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
