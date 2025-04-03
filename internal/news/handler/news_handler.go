package handler

import (
	"log"
	"strconv"

	"github.com/ahmadammarm/go-rest-api-template/helper/response"
	"github.com/ahmadammarm/go-rest-api-template/internal/middleware"
	"github.com/ahmadammarm/go-rest-api-template/internal/news/dto"
	newsService "github.com/ahmadammarm/go-rest-api-template/internal/news/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type NewsHandler struct {
    newsService newsService.NewsService
    validation *validator.Validate
}

func (handler *NewsHandler) GetAllNews(context *fiber.Ctx) error {
    userId, ok := context.Locals("user_id").(int)
    if !ok {
        return response.JSONResponse(context, 401, "Unauthorized", nil)
    }

    news, err := handler.newsService.GetAllNews(userId)
    if err != nil {
        return response.JSONResponse(context, 500, "Internal Server Error", nil)
    }

    return response.JSONResponse(context, 200, "Success", news)
}

func (handler *NewsHandler) GetNewsByID(context *fiber.Ctx) error {
    userId, ok := context.Locals("user_id").(int)

    if !ok {
        return response.JSONResponse(context, 401, "Unauthorized", nil)
    }

    newsId, err := strconv.Atoi(context.Params("id"))
    if err != nil {
        return response.JSONResponse(context, 400, "Bad Request", nil)
    }

    news, err := handler.newsService.GetNewsByID(userId, newsId)


    if err != nil {
        if err.Error() == "news not found" {
            return response.JSONResponse(context, 404, "Not Found", nil)
        }
        return response.JSONResponse(context, 500, "Internal Server Error", nil)
    }

    return response.JSONResponse(context, 200, "Success", news)
}

func (handler *NewsHandler) CreateNews(context *fiber.Ctx) error {
    userId, ok := context.Locals("user_id").(int)
    if !ok {
        return response.JSONResponse(context, 401, "Unauthorized", nil)
    }

    var news dto.NewsCreateRequest
    if err := context.BodyParser(&news); err != nil {
        return response.JSONResponse(context, 400, "Bad Request", nil)
    }

    if err := handler.validation.Struct(news); err != nil {
        return response.JSONResponse(context, 422, "Validation Error", nil)
    }

    news.AuthorId = userId

    if err := handler.newsService.CreateNews(userId, &news); err != nil {
        return response.JSONResponse(context, 500, "Internal Server Error", nil)
    }

    return response.JSONResponse(context, 201, "Created", nil)
}

func (handler *NewsHandler) UpdateNews(context *fiber.Ctx) error {
    userId, ok := context.Locals("user_id").(int)
    if !ok {
        return response.JSONResponse(context, 401, "Unauthorized", nil)
    }

    id, err := strconv.Atoi(context.Params("id"))
    if err != nil {
        return response.JSONResponse(context, 400, "Bad Request", nil)
    }

    var news dto.NewsUpdateRequest
    if err := context.BodyParser(&news); err != nil {
        return response.JSONResponse(context, 400, "Bad Request", nil)
    }

    if err := handler.validation.Struct(news); err != nil {
        return response.JSONResponse(context, 422, "Validation Error", nil)
    }

    news.AuthorId = userId

    if err := handler.newsService.UpdateNews(userId, id, news); err != nil {
        log.Println("Error updating news:", err)
        if err.Error() == "news not found" {
            return response.JSONResponse(context, 404, "Not Found", nil)
        }
        return response.JSONResponse(context, 500, "Internal Server Error", nil)
    }

    return response.JSONResponse(context, 200, "Success", nil)
}

func (handler *NewsHandler) DeleteNews(context *fiber.Ctx) error {
    userId, ok := context.Locals("user_id").(int)
    if !ok {
        return response.JSONResponse(context, 401, "Unauthorized", nil)
    }

    id, err := strconv.Atoi(context.Params("id"))
    if err != nil {
        return response.JSONResponse(context, 400, "Bad Request", nil)
    }
    if err := handler.newsService.DeleteNews(userId, id); err != nil {
        return response.JSONResponse(context, 500, "Internal Server Error", nil)
    }

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