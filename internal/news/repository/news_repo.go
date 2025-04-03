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
	UpdateNews(id int, news dto.NewsUpdateRequest) error
	DeleteNews(id int) error
}

type newsRepository struct {
	db *sql.DB
}

func (repo *newsRepository) GetAllNews() (*dto.NewsListResponse, error) {
	query := `SELECT n.id, n.title, n.content, n.user_id, u.name AS author_name, n.created_at, n.updated_at
              FROM news n
              JOIN users u ON n.user_id = u.id`

	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var news []dto.NewsResponse

	for rows.Next() {
		var n dto.NewsResponse
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.AuthorId, &n.AuthorName, &n.CreatedAt, &n.UpdatedAt)
		if err != nil {
			return nil, err
		}
		news = append(news, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	total := len(news)

	return &dto.NewsListResponse{
		News:  news,
		Total: total,
	}, nil
}

func (repo *newsRepository) GetNewsById(id int) (*dto.NewsResponse, error) {
	query := `SELECT n.id, n.title, n.content, n.user_id, u.name AS author_name, n.created_at, n.updated_at
              FROM news n
              JOIN users u ON n.user_id = u.id WHERE n.id = $1`

	rows, err := repo.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var n dto.NewsResponse
	if rows.Next() {
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.AuthorId, &n.AuthorName, &n.CreatedAt, &n.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return &n, nil
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nil, errors.New("news not found")
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

func (repo *newsRepository) UpdateNews(id int, news dto.NewsUpdateRequest) error {
    query := "UPDATE news SET title = $1, content = $2, user_id = $3, updated_at = $4 WHERE id = $5"

    updatedAt := time.Now()

    _, err := repo.db.Exec(query, news.Title, news.Content, news.AuthorId, updatedAt, id)

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
