package utils

import "github.com/gofiber/fiber/v2"

type ResponseData struct {
	Message string      `json:"message,omitempty"`
	Code    string      `json:"code,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func Response(ctx *fiber.Ctx, status int, data interface{}) error {
	var response ResponseData

	if data == nil {
		response.Result = data
	}
	return ctx.Status(status).JSON(response)
}

func ResponseError(ctx *fiber.Ctx, status int, message string, code string) error {
	return ctx.Status(status).JSON(&ResponseData{
		Message: message,
		Code:    code,
	})
}
