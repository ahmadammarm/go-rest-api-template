package handler

import (
	"log"
	"strconv"

	"github.com/ahmadammarm/go-rest-api-template/pkg/response"
	"github.com/ahmadammarm/go-rest-api-template/internal/middleware"
	"github.com/ahmadammarm/go-rest-api-template/internal/news/dto"
	newsService "github.com/ahmadammarm/go-rest-api-template/internal/news/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type NewsHandler struct {
	newsService newsService.NewsService
	validation  *validator.Validate
}

func (handler *NewsHandler) GetAllNews(context *fiber.Ctx) error {
	news, err := handler.newsService.GetAllNews()
	if err != nil {
		log.Println("Error fetching all news:", err)
		return response.JSONResponse(context, 500, "Internal Server Error", nil)
	}

	log.Println("Successfully fetched all news")
	return response.JSONResponse(context, 200, "Success", news)
}

func (handler *NewsHandler) GetNewsByID(context *fiber.Ctx) error {
	newsId, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		log.Println("Error parsing news ID:", err)
		return response.JSONResponse(context, 400, "Bad Request", nil)
	}

	news, err := handler.newsService.GetNewsByID(newsId)
	if err != nil {
		if err.Error() == "news not found" {
			log.Println("News not found with ID:", newsId)
			return response.JSONResponse(context, 404, "Not Found", nil)
		}
		log.Println("Error fetching news by ID:", err)
		return response.JSONResponse(context, 500, "Internal Server Error", nil)
	}

	log.Println("Successfully fetched news with ID:", newsId)
	return response.JSONResponse(context, 200, "Success", news)
}

func (handler *NewsHandler) CreateNews(context *fiber.Ctx) error {
	var news dto.NewsCreateRequest
	if err := context.BodyParser(&news); err != nil {
		log.Println("Error parsing request body for creating news:", err)
		return response.JSONResponse(context, 400, "Bad Request", nil)
	}

	if err := handler.validation.Struct(news); err != nil {
		log.Println("Validation error for creating news:", err)
		return response.JSONResponse(context, 422, "Validation Error", nil)
	}

	if err := handler.newsService.CreateNews(&news); err != nil {
		log.Println("Error creating news:", err)
		return response.JSONResponse(context, 500, "Internal Server Error", nil)
	}

	log.Println("Successfully created news")
	return response.JSONResponse(context, 201, "Created", nil)
}

func (handler *NewsHandler) UpdateNews(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		log.Println("Error parsing news ID for update:", err)
		return response.JSONResponse(context, 400, "Bad Request", nil)
	}

	var news dto.NewsUpdateRequest
	if err := context.BodyParser(&news); err != nil {
		log.Println("Error parsing request body for updating news:", err)
		return response.JSONResponse(context, 400, "Bad Request", nil)
	}

	if err := handler.validation.Struct(news); err != nil {
		log.Println("Validation error for updating news:", err)
		return response.JSONResponse(context, 422, "Validation Error", nil)
	}

	if err := handler.newsService.UpdateNews(id, news); err != nil {
		log.Println("Error updating news with ID:", id, "Error:", err)
		if err.Error() == "news not found" {
			return response.JSONResponse(context, 404, "Not Found", nil)
		}
		return response.JSONResponse(context, 500, "Internal Server Error", nil)
	}

	log.Println("Successfully updated news with ID:", id)
	return response.JSONResponse(context, 200, "Success", nil)
}

func (handler *NewsHandler) DeleteNews(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		log.Println("Error parsing news ID for deletion:", err)
		return response.JSONResponse(context, 400, "Bad Request", nil)
	}
	if err := handler.newsService.DeleteNews(id); err != nil {
		log.Println("Error deleting news with ID:", id, "Error:", err)
		return response.JSONResponse(context, 500, "Internal Server Error", nil)
	}

	log.Println("Successfully deleted news with ID:", id)
	return response.JSONResponse(context, 200, "Success", nil)
}

func (handler *NewsHandler) NewsRouters(router fiber.Router) {
	router.Use(middleware.JWTAuth())
	router.Get("/news", handler.GetAllNews)
	router.Get("/news/:id", handler.GetNewsByID)
	router.Post("/news", handler.CreateNews)
	router.Put("/news/:id", handler.UpdateNews)
	router.Delete("/news/:id", handler.DeleteNews)
}

func NewNewsHandler(newsService newsService.NewsService, validation *validator.Validate) *NewsHandler {
	return &NewsHandler{
		newsService: newsService,
		validation:  validation,
	}
}
