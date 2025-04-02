package dependencyinjection

import (
	"database/sql"

	"github.com/ahmadammarm/go-rest-api-template/internal/user/handler"
	"github.com/ahmadammarm/go-rest-api-template/internal/user/repository"
	"github.com/ahmadammarm/go-rest-api-template/internal/user/service"
	"github.com/go-playground/validator/v10"
)

func InitializeUser(db *sql.DB, validator *validator.Validate) *handler.UserHandler {
    userRepository := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepository)
    userHandler := handler.NewUserHandler(userService, validator)

    return userHandler
}