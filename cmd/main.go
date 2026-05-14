package main

import (
	"log"
	"os"

	"github.com/ahmadammarm/go-rest-api-template/config"
	news "github.com/ahmadammarm/go-rest-api-template/internal/news/dependency_injection"
	users "github.com/ahmadammarm/go-rest-api-template/internal/user/dependency_injection"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import middleware CORS
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	db, error := config.PostgresConnect()

	if error != nil {
		log.Printf("Failed to connect to database: %v", error)
		os.Exit(1)
	}

	defer db.Close()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_ALLOW_ORIGINS"),
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	users.InitializeUser(db, validator.New()).UserRouters(app)
	news.InitializeNews(db, validator.New()).NewsRouters(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if error := app.Listen(":" + port); error != nil {
		log.Printf("Failed to start server: %v", error)
		os.Exit(1)
	}
}
