package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	listnere "github.com/thedevflex/kubi8al-webhook/listner"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "http://localhost:3000, https://beta.campuspilot.in",
			AllowMethods: "GET,POST,DELETE",
			AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		},
	))

	app.Post("/webhook", listnere.WebhookPOST)
	app.Get("/webhook", listnere.WebhookGET)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "ok",
			"message": "healthy",
		})
	})
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "you missed the server",
		})
	})
}
