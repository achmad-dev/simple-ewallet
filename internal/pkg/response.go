package pkg

import (
	"github.com/achmad-dev/simple-ewallet/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func SendResponse(c *fiber.Ctx, res *dto.Response) error {
	response := fiber.Map{
		"success": res.Success,
		"message": res.Message,
		"status":  res.Status,
		// "data":   res.Data,
	}
	if res.Data != nil {
		response["data"] = res.Data
	}
	return c.Status(res.Status).JSON(response)
}

func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	res := &dto.Response{
		Success: true,
		Message: message,
		Status:  fiber.StatusOK,
		Data:    data,
	}
	return SendResponse(c, res)
}

func SuccessResponseWithStatus(c *fiber.Ctx, data interface{}, message string, status int) error {
	res := &dto.Response{
		Success: true,
		Message: message,
		Status:  status,
		Data:    data,
	}
	return SendResponse(c, res)
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	res := &dto.Response{
		Success: false,
		Message: message,
		Status:  status,
	}
	return SendResponse(c, res)
}
