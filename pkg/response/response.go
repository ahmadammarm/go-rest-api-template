package response

import "github.com/gofiber/fiber/v2"

type Response struct {
    Message string `json:"message"`
    Data   any    `json:"data"`
}

func JSONResponse(context *fiber.Ctx, statusCode int, message string, data any) error {
    response := Response{
        Message: message,
        Data:    data,
    }

    return context.Status(statusCode).JSON(response)
}