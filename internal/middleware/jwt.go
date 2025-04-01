package middleware

import (
	"os"
	"strings"

	"github.com/ahmadammarm/go-rest-api-template/helper/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func JWTAuth() fiber.Handler {
	return func(context *fiber.Ctx) error {
		stringToken := context.Get("Authorization")
		if stringToken == "" {
			return response.JSONResponse(context, 401, "Unauthorized", nil)
		}

		stringToken = strings.TrimPrefix(stringToken, "Bearer ")

		token, err := jwt.ParseWithClaims(stringToken, *&JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			secretKey := os.Getenv("JWT_SECRET_KEY")
			if secretKey == "" {
				return nil, response.JSONResponse(context, 500, "Internal Server Error", nil)
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return response.JSONResponse(context, 401, "Unauthorized", nil)
		}

		claims, ok := token.Claims.(*JWTClaims)
		if !ok {
			return response.JSONResponse(context, 401, "Unauthorized", nil)
		}

		context.Locals("user_id", claims.UserID)

		return context.Next()
	}

}
