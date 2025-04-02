package main

import (
	"log"
	"os"

	"github.com/ahmadammarm/go-rest-api-template/config"
	users "github.com/ahmadammarm/go-rest-api-template/internal/user/dependency_injection"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

    api := app.Group("/api/v1")

    users.InitializeUser(db, validator.New()).UserRouters(api)

    if error := app.Listen(":8080"); error != nil {
        log.Printf("Failed to start server: %v", error)
        os.Exit(1)
    }
    
    log.Println("Server started on port 8080")

}