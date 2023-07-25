package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func Start(app *fiber.App) error {

	// Create a context with a cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// Listen for OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		fmt.Printf("Server running on port %s\n", os.Getenv("APP_PORT"))
		log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
		cancel()
	}()

	// Wait for OS signal or context cancellation
	select {
	case signal := <-signalChan:
		fmt.Printf("Received OS signal: %v\n", signal)
	case <-ctx.Done():
		fmt.Println("Context canceled")
	}

	//shutdown the server
	if err := app.Shutdown(); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	fmt.Println("Server shutdown complete")

	return nil
}
