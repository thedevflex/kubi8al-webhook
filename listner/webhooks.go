package listener

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/thedevflex/kubi8al-webhook/emitter"
	logs "github.com/thedevflex/kubi8al-webhook/utils/logger"
)

type PayloadResponse struct {
	Message string `json:"message"`
}

func WebhookPOST(c *fiber.Ctx) error {
	logs.InitLogger()
	authHeader := c.Get("Authorization")
	const bearerPrefix = "Bearer "
	var token string
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		token = authHeader[len(bearerPrefix):]
	}
	if token != os.Getenv("WEBHOOK_SECRET") {
		logs.Error("unauthorized access attempt with token: ", token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	var payload emitter.PackagePublishedEvent
	if err := c.BodyParser(&payload); err != nil {
		logs.Error("cannot parse webhook payload", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in parsing the payload",
			"error":   err.Error(),
		})
	}

	if payload.Event == "package" && payload.Payload.Action == "published" {
		logs.Infof("Received %s event for repository: %s", payload.Event, payload.Payload.Repository.FullName)
		if err := emitter.EmitPackagePayload(payload); err != nil {
			logs.Error("failed to emit package payload", err)
		}
		return c.Status(fiber.StatusOK).JSON(PayloadResponse{
			Message: "Webhook Package processed successfully",
		})
	} else {
		logs.Info("Received non-package event: ", payload.Event)
		return c.Status(fiber.StatusOK).JSON(PayloadResponse{
			Message: "Webhook processed successfully Not yet Implemented ",
		})
	}

}

func WebhookGET(c *fiber.Ctx) error {
	logs.InitLogger()
	authHeader := c.Get("Authorization")
	const bearerPrefix = "Bearer "
	var token string
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		token = authHeader[len(bearerPrefix):]
	}
	if token != os.Getenv("WEBHOOK_SECRET") {
		logs.Error("unauthorized access attempt with token: ", token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GET request successful",
		"data":    "Thought of the day: Premature optimization is the root of all evil.",
	})

}
