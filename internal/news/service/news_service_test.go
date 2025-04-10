package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ahmadammarm/go-rest-api-template/internal/news/dto"
	"github.com/ahmadammarm/go-rest-api-template/internal/news/service"
)

type MockNewsRepository struct {
	mock.Mock
}

func (m *MockNewsRepository) GetAllNews() (*dto.NewsListResponse, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.NewsListResponse), args.Error(1)
}

func (m *MockNewsRepository) GetNewsById(id int) (*dto.NewsResponse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.NewsResponse), args.Error(1)
}

func (m *MockNewsRepository) CreateNews(news *dto.NewsCreateRequest) error {
	args := m.Called(news)
	return args.Error(0)
}

func (m *MockNewsRepository) UpdateNews(id int, news dto.NewsUpdateRequest) error {
	args := m.Called(id, news)
	return args.Error(0)
}

func (m *MockNewsRepository) DeleteNews(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllNews(t *testing.T) {
	mockRepo := new(MockNewsRepository)
	newsService := service.NewNewsService(mockRepo)

	t.Run("success with news", func(t *testing.T) {
		expectedNews := &dto.NewsListResponse{
			News: []dto.NewsResponse{
				{
					ID:         1,
					Title:      "Test News 1",
					Content:    "Content 1",
					AuthorId:   1,
					AuthorName: "Author 1",
				},
				{
					ID:         2,
					Title:      "Test News 2",
					Content:    "Content 2",
					AuthorId:   2,
					AuthorName: "Author 2",
				},
			},
			Total: 2,
		}

		mockRepo.On("GetAllNews").Return(expectedNews, nil).Once()

		result, err := newsService.GetAllNews()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, result.Total)
		assert.Len(t, result.News, 2)
		assert.Equal(t, expectedNews, result)

		mockRepo.AssertExpectations(t)
	})

	t.Run("repository returns error", func(t *testing.T) {
		expectedError := errors.New("database error")
		mockRepo.On("GetAllNews").Return(nil, expectedError).Once()

		result, err := newsService.GetAllNews()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error getting all news")

		mockRepo.AssertExpectations(t)
	})

	t.Run("news is nil", func(t *testing.T) {
		mockRepo.On("GetAllNews").Return(nil, nil).Once()

		result, err := newsService.GetAllNews()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.Total)
		assert.Empty(t, result.News)

		mockRepo.AssertExpectations(t)
	})

	t.Run("empty news list", func(t *testing.T) {
		emptyNews := &dto.NewsListResponse{
			News:  []dto.NewsResponse{},
			Total: 0,
		}

		mockRepo.On("GetAllNews").Return(emptyNews, nil).Once()

		result, err := newsService.GetAllNews()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.Total)
		assert.Empty(t, result.News)

		mockRepo.AssertExpectations(t)
	})
}
