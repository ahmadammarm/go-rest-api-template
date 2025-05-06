package middleware

import (
	"fmt"
	"os"
	"strings"
	"log"
	"github.com/ahmadammarm/go-rest-api-template/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func JWTAuth() fiber.Handler {
	return func(context *fiber.Ctx) error {
		authHeader := context.Get("Authorization")
		if authHeader == "" {
			return response.JSONResponse(context, 401, "Unauthorized: No Token Provided", nil)
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.JSONResponse(context, 401, "Unauthorized: Invalid Token Format", nil)
		}

		stringToken := strings.TrimPrefix(authHeader, "Bearer ")

		if stringToken == "" {
			return response.JSONResponse(context, 401, "Unauthorized: Empty Token", nil)
		}

		secret := os.Getenv("JWT_SECRET_KEY")
		if secret == "" {
			log.Println("Error: JWT_SECRET_KEY is missing in environment variables")
			return response.JSONResponse(context, 500, "Internal Server Error", nil)
		}

		token, err := jwt.ParseWithClaims(stringToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			log.Printf("Token parse error: %v", err)
			return response.JSONResponse(context, 401, "Unauthorized: Token Invalid", nil)
		}

		if !token.Valid {
			return response.JSONResponse(context, 401, "Unauthorized: Token Not Valid", nil)
		}

		claims, ok := token.Claims.(*JWTClaims)
		if !ok {
			return response.JSONResponse(context, 401, "Unauthorized: Failed to Parse Token Claims", nil)
		}

		context.Locals("user_id", claims.UserID)

		return context.Next()
	}
}