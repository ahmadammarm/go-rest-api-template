package repository_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahmadammarm/go-rest-api-template/internal/news/dto"
	"github.com/ahmadammarm/go-rest-api-template/internal/news/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetAllNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewNewsRepository(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "author_name", "created_at", "updated_at"}).
			AddRow(1, "Title 1", "Content 1", 1, "Author 1", time.Now(), time.Now()).
			AddRow(2, "Title 2", "Content 2", 2, "Author 2", time.Now(), time.Now())

		mock.ExpectQuery("SELECT n.id, n.title, n.content, n.user_id, u.name AS author_name, n.created_at, n.updated_at").
			WillReturnRows(rows)

		result, err := repo.GetAllNews()
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, result.Total)
		assert.Len(t, result.News, 2)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery("SELECT n.id, n.title, n.content, n.user_id, u.name AS author_name, n.created_at, n.updated_at").
			WillReturnError(errors.New("query error"))

		result, err := repo.GetAllNews()
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("scan error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "author_name", "created_at", "updated_at"}).
			AddRow("invalid", "Title 1", "Content 1", 1, "Author 1", time.Now(), time.Now())

		mock.ExpectQuery("SELECT n.id, n.title, n.content, n.user_id, u.name AS author_name, n.created_at, n.updated_at").
			WillReturnRows(rows)

		result, err := repo.GetAllNews()
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetNewsByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewNewsRepository(db)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "author_name", "created_at", "updated_at"}).
			AddRow(1, "Title 1", "Content 1", 1, "Author 1", time.Now(), time.Now())

		mock.ExpectQuery("SELECT n.id, n.title, n.content, n.user_id, u.name AS author_name, n.created_at, n.updated_at").
			WithArgs(1).
			WillReturnRows(rows)

		result, err := repo.GetNewsById(1)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery("SELECT n.id, n.title, n.content, n.user_id, u.name AS author_name, n.created_at, n.updated_at").
			WithArgs(999).
			WillReturnError(sql.ErrNoRows)

		result, err := repo.GetNewsById(999)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("query error", func(t *testing.T) {
		mock.ExpectQuery("SELECT n.id, n.title, n.content, n.user_id, u.name AS author_name, n.created_at, n.updated_at").
			WithArgs(1).
			WillReturnError(errors.New("query error"))

		result, err := repo.GetNewsById(1)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestCreateNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewNewsRepository(db)

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery("INSERT INTO news \\(title, content, user_id\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id").
			WithArgs("Title 1", "Content 1", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		req := &dto.NewsCreateRequest{
			Title:    "Title 1",
			Content:  "Content 1",
			AuthorId: 1,
		}
		err := repo.CreateNews(req)
		assert.NoError(t, err)
	})

	t.Run("exec error", func(t *testing.T) {
		mock.ExpectQuery("INSERT INTO news \\(title, content, user_id\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id").
			WithArgs("Title 1", "Content 1", 1).
			WillReturnError(errors.New("exec error"))

		req := &dto.NewsCreateRequest{
			Title:    "Title 1",
			Content:  "Content 1",
			AuthorId: 1,
		}
		err := repo.CreateNews(req)
		assert.Error(t, err)
	})
}

func TestUpdateNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewNewsRepository(db)

	t.Run("success", func(t *testing.T) {
		// Use ExpectExec instead of ExpectQuery since your code uses Exec
		mock.ExpectExec("UPDATE news SET title = \\$1, content = \\$2, user_id = \\$3, updated_at = \\$4 WHERE id = \\$5").
			WithArgs("Title 1", "Content 1", 0, sqlmock.AnyArg(), 1). // Use AnyArg() for the timestamp
			WillReturnResult(sqlmock.NewResult(1, 1))

		req := &dto.NewsUpdateRequest{
			ID:      1,
			Title:   "Title 1",
			Content: "Content 1",
		}
		err := repo.UpdateNews(req.ID, *req)
		assert.NoError(t, err)
	})

	t.Run("exec error", func(t *testing.T) {
		mock.ExpectExec("UPDATE news SET title = \\$1, content = \\$2, user_id = \\$3, updated_at = \\$4 WHERE id = \\$5").
			WithArgs("Title 1", "Content 1", 0, sqlmock.AnyArg(), 999).
			WillReturnError(errors.New("exec error"))

		req := &dto.NewsUpdateRequest{
			ID:      999,
			Title:   "Title 1",
			Content: "Content 1",
		}
		err := repo.UpdateNews(req.ID, *req)
		assert.Error(t, err)
	})
}

func TestDeleteNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewNewsRepository(db)

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM news WHERE id = \\$1").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DeleteNews(1)
		assert.NoError(t, err)
	})

	t.Run("exec error", func(t *testing.T) {
		mock.ExpectExec("DELETE FROM news WHERE id = \\$1").
			WithArgs(999).
			WillReturnError(errors.New("exec error"))

		err := repo.DeleteNews(999)
		assert.Error(t, err)
	})
}
