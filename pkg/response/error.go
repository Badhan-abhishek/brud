package response

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func SendError(c *fiber.Ctx, status int, message string) error {
	response := ErrorResponse{
		Message: message,
	}
	return c.Status(status).JSON(response)
}

func SendData(c *fiber.Ctx, status int, data any) error {
	return c.Status(status).JSON(data)
}

func SendSuccess(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(SuccessResponse{
		Message: message,
	})
}
