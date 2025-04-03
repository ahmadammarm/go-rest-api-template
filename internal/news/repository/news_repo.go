package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/ahmadammarm/go-rest-api-template/internal/news/dto"
)

type NewsRepository interface {
	GetAllNews() (*dto.NewsListResponse, error)
	GetNewsById(id int) (*dto.NewsResponse, error)
	CreateNews(news *dto.NewsCreateRequest) error
	UpdateNews(news *dto.NewsUpdateRequest) error
	DeleteNews(id int) error
}

type newsRepository struct {
	db *sql.DB
}

func (repo *newsRepository) GetAllNews() (*dto.NewsListResponse, error) {
	query := `SELECT id, title, content, user_id, created_at, updated_at FROM news`

	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	newsList := &dto.NewsListResponse{}

	for rows.Next() {
		news := dto.NewsResponse{}
		err := rows.Scan(&news.ID, &news.Title, &news.Content, &news.AuthorId, &news.CreatedAt, &news.UpdatedAt)

		if err != nil {
			return nil, err
		}

		newsList.News = append(newsList.News, news)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

    return &dto.NewsListResponse{
        News:  newsList.News,
    }, nil
}

func (repo *newsRepository) GetNewsById(id int) (*dto.NewsResponse, error) {
	query := "SELECT * FROM news WHERE id = $1"
	news := &dto.NewsResponse{}

	err := repo.db.QueryRow(query, id).Scan(&news.ID, &news.Title, &news.Content, &news.AuthorId, &news.CreatedAt, &news.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("news not found")
		}
	}
	return news, nil
}

func (repo *newsRepository) CreateNews(news *dto.NewsCreateRequest) error {
	query := "INSERT INTO news (title, content, user_id) VALUES ($1, $2, $3) RETURNING id"

	news.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)

	err := repo.db.QueryRow(query, news.Title, news.Content, news.AuthorId).Scan(&news.ID)

	if err != nil {
		return err
	}

	return nil

}

func (repo *newsRepository) UpdateNews(news *dto.NewsUpdateRequest) error {
	query := "UPDATE news SET title = $1, content = $2, user_id = $3, updated_at = $4 WHERE id = $5"

	news.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)

	_, err := repo.db.Exec(query, news.Title, news.Content, news.AuthorId, news.UpdatedAt, news.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *newsRepository) DeleteNews(id int) error {
	query := "DELETE FROM news WHERE id = $1"

	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func NewNewsRepository(db *sql.DB) NewsRepository {
	return &newsRepository{db: db}
}
