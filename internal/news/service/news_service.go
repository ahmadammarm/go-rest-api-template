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
    GetNewsByAuthorId(userID int, authorId int) (*dto.NewsByAuthorId, error)
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
    }, nil
}

