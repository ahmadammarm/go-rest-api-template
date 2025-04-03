package dependency_injection

import (
	"database/sql"

	"github.com/ahmadammarm/go-rest-api-template/internal/news/handler"
    newsRepository "github.com/ahmadammarm/go-rest-api-template/internal/news/repository"
    newsService "github.com/ahmadammarm/go-rest-api-template/internal/news/service"
    userRepository "github.com/ahmadammarm/go-rest-api-template/internal/user/repository"
	"github.com/go-playground/validator/v10"
)

func InitializeNews(db *sql.DB, validator *validator.Validate) *handler.NewsHandler {
    newsRepo := newsRepository.NewNewsRepository(db)
    userRepo := userRepository.NewUserRepository(db)
    newsService := newsService.NewNewsService(newsRepo, userRepo)

    newsHandler := handler.NewNewsHandler(newsService, validator)

    return newsHandler
}