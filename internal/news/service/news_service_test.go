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

func TestGetNewsByID(t *testing.T) {
	mockRepo := new(MockNewsRepository)
	newsService := service.NewNewsService(mockRepo)

	t.Run("success", func(t *testing.T) {
		newsID := 1
		expectedNews := &dto.NewsResponse{
			ID:         int(newsID),
			Title:      "Test News",
			Content:    "Test Content",
			AuthorId:   1,
			AuthorName: "Author 1",
			CreatedAt:  "2023-01-01T00:00:00Z",
			UpdatedAt:  "2023-01-01T00:00:00Z",
		}

		mockRepo.On("GetNewsById", newsID).Return(expectedNews, nil).Once()

		result, err := newsService.GetNewsByID(newsID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedNews, result)

		mockRepo.AssertExpectations(t)
	})

	t.Run("news not found", func(t *testing.T) {
		newsID := 999
		notFoundErr := errors.New("news not found")

		mockRepo.On("GetNewsById", newsID).Return(nil, notFoundErr).Once()

		result, err := newsService.GetNewsByID(newsID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error getting news by ID")

		mockRepo.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		newsID := 1
		dbErr := errors.New("database connection error")

		mockRepo.On("GetNewsById", newsID).Return(nil, dbErr).Once()

		result, err := newsService.GetNewsByID(newsID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "error getting news by ID")

		mockRepo.AssertExpectations(t)
	})

	t.Run("nil news with no error", func(t *testing.T) {

		newsID := 1

		mockRepo.On("GetNewsById", newsID).Return(nil, nil).Once()

		result, err := newsService.GetNewsByID(newsID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "not found")

		mockRepo.AssertExpectations(t)
	})
}

func TestCreateNews(t *testing.T) {
    mockRepo := new(MockNewsRepository)
    newsService := service.NewNewsService(mockRepo)

    t.Run("success", func(t *testing.T) {
        newsRequest := &dto.NewsCreateRequest{
            Title:   "New News",
            Content: "New Content",
            AuthorId: 1,
        }

        mockRepo.On("CreateNews", newsRequest).Return(nil).Once()

        err := newsService.CreateNews(newsRequest)

        assert.NoError(t, err)

        mockRepo.AssertExpectations(t)
    })

    t.Run("repository returns error", func(t *testing.T) {
        newsRequest := &dto.NewsCreateRequest{
            Title:   "New News",
            Content: "New Content",
            AuthorId: 1,
        }
        expectedError := errors.New("database error")

        mockRepo.On("CreateNews", newsRequest).Return(expectedError).Once()

        err := newsService.CreateNews(newsRequest)

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "error creating news")

        mockRepo.AssertExpectations(t)
    })
}

func TestUpdateNews(t *testing.T) {
    mockRepo := new(MockNewsRepository)
    newsService := service.NewNewsService(mockRepo)

    t.Run("success", func(t *testing.T) {
        newsID := 1
        newsUpdateRequest := dto.NewsUpdateRequest{
            ID:       newsID,
            Title:    "Updated Title",
            Content:  "Updated Content",
            AuthorId: 1,
        }

        mockRepo.On("UpdateNews", newsID, newsUpdateRequest).Return(nil).Once()

        err := newsService.UpdateNews(newsID, newsUpdateRequest)

        assert.NoError(t, err)

        mockRepo.AssertExpectations(t)
    })

    t.Run("repository returns error", func(t *testing.T) {
        newsID := 1
        newsUpdateRequest := dto.NewsUpdateRequest{
            ID:       newsID,
            Title:    "Updated Title",
            Content:  "Updated Content",
            AuthorId: 1,
        }
        expectedError := errors.New("database error")

        mockRepo.On("UpdateNews", newsID, newsUpdateRequest).Return(expectedError).Once()

        err := newsService.UpdateNews(newsID, newsUpdateRequest)

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "error updating news")

        mockRepo.AssertExpectations(t)
    })
}

func TestDeleteNews(t *testing.T) {
    mockRepo := new(MockNewsRepository)
    newsService := service.NewNewsService(mockRepo)

    t.Run("success", func(t *testing.T) {
        newsID := 3

        mockRepo.On("DeleteNews", newsID).Return(nil).Once()

        err := newsService.DeleteNews(newsID)

        assert.NoError(t, err)

        mockRepo.AssertExpectations(t)
    })

    t.Run("repository returns error", func(t *testing.T) {
        newsID := 3
        expectedError := errors.New("database error")

        mockRepo.On("DeleteNews", newsID).Return(expectedError).Once()

        err := newsService.DeleteNews(newsID)

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "error deleting news")

        mockRepo.AssertExpectations(t)
    })
}

