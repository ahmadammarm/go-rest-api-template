package service

import (
	"fmt"

	"github.com/ahmadammarm/go-rest-api-template/internal/news/dto"
	newsRepo "github.com/ahmadammarm/go-rest-api-template/internal/news/repository"
	userRepo "github.com/ahmadammarm/go-rest-api-template/internal/user/repository"
)

type NewsService interface {
	GetAllNews(userID int) (*dto.NewsListResponse, error)
    GetNewsByID(userID int, id int) (*dto.NewsResponse, error)
    CreateNews(userID int, news *dto.NewsCreateRequest) error
    UpdateNews(userID int, news *dto.NewsUpdateRequest) error
    DeleteNews(userID int, id int) error
}

type newsServiceImpl struct {
    newsRepo newsRepo.NewsRepository
    userRepo userRepo.UserRepo
}

func (service *newsServiceImpl) GetAllNews(userID int) (*dto.NewsListResponse, error) {
    userExists, err := service.userRepo.IsUserExist(userID)

    if err != nil {
        return nil, fmt.Errorf("error checking user existence: %w", err)
    }

    if !userExists {
        return nil, fmt.Errorf("user with ID %d does not exist", userID)
    }

    news, err := service.newsRepo.GetAllNews()

    if err != nil {
        return nil, fmt.Errorf("error getting news: %w", err)
    }

    filteredNews := []dto.NewsResponse{}

    for _, n := range news.News {
        if n.AuthorId == userID {
            filteredNews = append(filteredNews, n)
        }
    }

    return &dto.NewsListResponse{
        News:  filteredNews,
        Total: len(filteredNews),
    }, nil
}

func (service *newsServiceImpl) GetNewsByID(userID int, id int) (*dto.NewsResponse, error) {
    userExists, err := service.userRepo.IsUserExist(userID)

    if err != nil {
        return nil, fmt.Errorf("error checking user existence: %w", err)
    }

    if !userExists {
        return nil, fmt.Errorf("user with ID %d does not exist", userID)
    }

    news, err := service.newsRepo.GetNewsById(id)

    if err != nil {
        return nil, fmt.Errorf("error getting news: %w", err)
    }

    return news, nil
}

func (service *newsServiceImpl) CreateNews(userID int, news *dto.NewsCreateRequest) error {
    userExists, err := service.userRepo.IsUserExist(userID)

    if err != nil {
        return fmt.Errorf("error checking user existence: %w", err)
    }

    if !userExists {
        return fmt.Errorf("user with ID %d does not exist", userID)
    }

    news.AuthorId = userID

    err = service.newsRepo.CreateNews(news)

    if err != nil {
        return fmt.Errorf("error creating news: %w", err)
    }

    return nil
}

func (service *newsServiceImpl) UpdateNews(userID int, news *dto.NewsUpdateRequest) error {
    userExists, err := service.userRepo.IsUserExist(userID)

    if err != nil {
        return fmt.Errorf("error checking user existence: %w", err)
    }

    if !userExists {
        return fmt.Errorf("user with ID %d does not exist", userID)
    }

    news.AuthorId = userID

    err = service.newsRepo.UpdateNews(news)

    if err != nil {
        return fmt.Errorf("error updating news: %w", err)
    }

    return nil
}

func (service *newsServiceImpl) DeleteNews(userID int, id int) error {
    userExists, err := service.userRepo.IsUserExist(userID)

    if err != nil {
        return fmt.Errorf("error checking user existence: %w", err)
    }

    if !userExists {
        return fmt.Errorf("user with ID %d does not exist", userID)
    }

    err = service.newsRepo.DeleteNews(id)

    if err != nil {
        return fmt.Errorf("error deleting news: %w", err)
    }

    return nil
}

func NewNewsService(newsRepo newsRepo.NewsRepository, userRepo userRepo.UserRepo) NewsService {
    return &newsServiceImpl{
        newsRepo: newsRepo,
        userRepo: userRepo,
    }
}