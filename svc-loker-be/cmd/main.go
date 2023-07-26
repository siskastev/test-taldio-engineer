package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"svc-loker-be/internal/database"
	"svc-loker-be/internal/routes"
	"syscall"

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

	routes.Setup(app)

	ctx, cancel := context.WithCancel(context.Background())

	// add goroutine for running server in background
	go func() {
		fmt.Printf("Server running on port %s\n", os.Getenv("APP_PORT"))

		if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	//channel for catch signal SIGINT (Ctrl+C) dan SIGTERM
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// wait signal to stop server
	sig := <-sigCh
	fmt.Printf("Received signal: %v\n", sig)

	// call ctx cancel
	cancel()

	// waiting server to stop
	<-ctx.Done()

	fmt.Println("Server shutdown complete")
}
