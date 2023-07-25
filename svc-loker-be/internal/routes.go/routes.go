package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
