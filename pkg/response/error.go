package response

import (
	"fmt"

	customerrors "be.blog/pkg/custom_errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorHandler struct {
	ctx        *fiber.Ctx
	name       string
	pluralName string
}

func NewErrorHandler(ctx *fiber.Ctx, name, pluralName string) *ErrorHandler {
	return &ErrorHandler{
		ctx:        ctx,
		name:       name,
		pluralName: pluralName,
	}
}

func (e *ErrorHandler) SendCustomError(err error) error {
	if err == nil {
		return nil
	}

	var status int
	var message string
	errorMessage := err.Error()

	switch errorMessage {
	case customerrors.RepoItemNotFound, customerrors.RepoIDNotFound:
		status = fiber.StatusNotFound
		message = fmt.Sprintf("%s not found", e.name)

	case customerrors.RepoItemAlreadyExists:
		status = fiber.StatusConflict
		message = fmt.Sprintf("%s already exists", e.name)

	case customerrors.RepoForeignKeyViolation:
		status = fiber.StatusBadRequest
		message = fmt.Sprintf("Cannot delete %s, it is referenced by another %s", e.name, e.pluralName)

	default:
		status = fiber.StatusInternalServerError
		message = fmt.Sprintf("An error occurred while processing %s: %s", e.name, errorMessage)
	}

	return e.SendError(status, message)
}

func (e *ErrorHandler) SendError(status int, message string) error {
	response := ErrorResponse{
		Message: message,
	}
	return e.ctx.Status(status).JSON(response)
}

func (e *ErrorHandler) SendData(status int, data any) error {
	return e.ctx.Status(status).JSON(data)
}

func (e *ErrorHandler) SendSuccess(status int, message string) error {
	return e.ctx.Status(status).JSON(SuccessResponse{
		Message: message,
	})
}
