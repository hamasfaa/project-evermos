package exception

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/hamasfaa/project-evermos/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Bad Request",
			Errors:  messages,
			Data:    nil,
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Not Found",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Unauthorized",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
		Status:  false,
		Message: "Internal Server Error",
		Errors:  []string{err.Error()},
		Data:    nil,
	})
}
