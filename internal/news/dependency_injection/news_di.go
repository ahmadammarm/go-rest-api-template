package dependency_injection

import (
	"database/sql"

	"github.com/ahmadammarm/go-rest-api-template/internal/news/handler"
    newsRepository "github.com/ahmadammarm/go-rest-api-template/internal/news/repository"
    newsService "github.com/ahmadammarm/go-rest-api-template/internal/news/service"
	"github.com/go-playground/validator/v10"
)

func InitializeNews(db *sql.DB, validator *validator.Validate) *handler.NewsHandler {
    newsRepo := newsRepository.NewNewsRepository(db)
    newsService := newsService.NewNewsService(newsRepo)

    newsHandler := handler.NewNewsHandler(newsService, validator)

    return newsHandler
}