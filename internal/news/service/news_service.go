package service

import (
	"fmt"
	"log"

	"github.com/ahmadammarm/go-rest-api-template/internal/news/dto"
	newsRepo "github.com/ahmadammarm/go-rest-api-template/internal/news/repository"
)

type NewsService interface {
	GetAllNews() (*dto.NewsListResponse, error)
	GetNewsByID(id int) (*dto.NewsResponse, error)
	CreateNews(news *dto.NewsCreateRequest) error
	UpdateNews(newsId int, news dto.NewsUpdateRequest) error
	DeleteNews(id int) error
}

type newsServiceImpl struct {
	newsRepo newsRepo.NewsRepository
}

func (service *newsServiceImpl) GetAllNews() (*dto.NewsListResponse, error) {
	log.Println("Fetching all news...")
	news, err := service.newsRepo.GetAllNews()

	if err != nil {
		log.Printf("Error fetching all news: %v", err)
		return nil, fmt.Errorf("error getting all news: %w", err)
	}

	if news == nil {
		log.Println("No news found, returning empty list")
		return &dto.NewsListResponse{
			News:  []dto.NewsResponse{},
			Total: 0,
		}, nil
	}

	log.Println("Successfully fetched all news")
	return news, nil
}

func (service *newsServiceImpl) GetNewsByID(id int) (*dto.NewsResponse, error) {
	log.Printf("Fetching news by ID: %d...", id)
	news, err := service.newsRepo.GetNewsById(id)

	if err != nil {
		log.Printf("Error fetching news by ID %d: %v", id, err)
		return nil, fmt.Errorf("error getting news by ID: %w", err)
	}

	if news == nil {
		log.Printf("News with ID %d not found", id)
		return nil, fmt.Errorf("news with ID %d not found", id)
	}

	log.Printf("Successfully fetched news by ID: %d", id)
	return news, nil
}

func (service *newsServiceImpl) CreateNews(news *dto.NewsCreateRequest) error {
	log.Println("Creating news...")
	err := service.newsRepo.CreateNews(news)

	if err != nil {
		log.Printf("Error creating news: %v", err)
		return fmt.Errorf("error creating news: %w", err)
	}

	log.Println("Successfully created news")
	return nil
}

func (service *newsServiceImpl) UpdateNews(newsId int, news dto.NewsUpdateRequest) error {
	log.Printf("Updating news with ID: %d...", newsId)
	err := service.newsRepo.UpdateNews(newsId, news)

	if err != nil {
		log.Printf("Error updating news with ID %d: %v", newsId, err)
		return fmt.Errorf("error updating news: %w", err)
	}

	log.Printf("Successfully updated news with ID: %d", newsId)
	return nil
}

func (service *newsServiceImpl) DeleteNews(id int) error {
	log.Printf("Deleting news with ID: %d...", id)
	err := service.newsRepo.DeleteNews(id)

	if err != nil {
		log.Printf("Error deleting news with ID %d: %v", id, err)
		return fmt.Errorf("error deleting news: %w", err)
	}

	log.Printf("Successfully deleted news with ID: %d", id)
	return nil
}

func NewNewsService(newsRepo newsRepo.NewsRepository) NewsService {
	log.Println("Initializing NewsService...")
	return &newsServiceImpl{
		newsRepo: newsRepo,
	}
}
