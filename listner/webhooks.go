package listener

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/thedevflex/kubi8al-webhook/emitter"
	"github.com/thedevflex/kubi8al-webhook/model"
	logs "github.com/thedevflex/kubi8al-webhook/utils/logger"
)

type PayloadResponse struct {
	Message string `json:"message"`
}

func WebhookPOST(c *fiber.Ctx) error {
	logs.InitLogger()

	token := c.Query("token")
	if token != os.Getenv("WEBHOOK_SECRET") {
		logs.Error("unauthorized access attempt with token: ", token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	var payload model.ParsedWebHookPayload
	if err := c.BodyParser(&payload); err != nil {
		logs.Error("cannot parse webhook payload", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error in parsing the payload",
			"error":   err.Error(),
		})
	}
	logs.Infof("Received %s event for repository: %s", payload.Event, payload.Payload.Repository.FullName)
	if err := emitter.EmitWebhookPayload(payload); err != nil {
		logs.Error("failed to emit webhook payload", err)
	}

	return c.Status(fiber.StatusOK).JSON(PayloadResponse{
		Message: "Webhook processed successfully",
	})
}

func WebhookGET(c *fiber.Ctx) error {
	logs.InitLogger()
	token := c.Query("token")
	if token != os.Getenv("WEBHOOK_SECRET") {
		logs.Error("unauthorized access attempt with token: ", token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "GET request successful",
		"data":    "This is a sample response for GET request",
	})

}
