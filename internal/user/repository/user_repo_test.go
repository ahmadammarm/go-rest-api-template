package repository_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	userDTO "github.com/ahmadammarm/go-rest-api-template/internal/user/dto"
	"github.com/ahmadammarm/go-rest-api-template/internal/user/repository"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO users \(email, name, password\) VALUES \(\$1, \$2, \$3\) RETURNING id`).
		WithArgs("test@example.com", "Test User", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserRegisterRequest{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "password123",
	}

	err = repo.RegisterUser(request)
	assert.NoError(t, err)
	assert.Equal(t, 1, request.ID)
}

func TestRegisterUser_HashPasswordError(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserRegisterRequest{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: string(make([]byte, bcrypt.MaxCost+1)),
	}

	err = repo.RegisterUser(request)
	assert.Error(t, err)
}

func TestRegisterUser_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO users \(email, name, password\) VALUES \(\$1, \$2, \$3\) RETURNING id`).
		WithArgs("test@example.com", "Test User", sqlmock.AnyArg()).
		WillReturnError(errors.New("query error"))
	mock.ExpectRollback()

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserRegisterRequest{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "password123",
	}

	err = repo.RegisterUser(request)
	assert.Error(t, err)
}

func TestLoginUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	mock.ExpectQuery(`SELECT id, name, email, password FROM users WHERE email = \$1`).
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password"}).
			AddRow(1, "Test User", "test@example.com", hashedPassword))

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserLoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	response, err := repo.LoginUser(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Test User", response.Name)
	assert.Equal(t, "test@example.com", response.Email)
}

func TestLoginUser_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, name, email, password FROM users WHERE email = \$1`).
		WithArgs("nonexistent@example.com").
		WillReturnError(sql.ErrNoRows)

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserLoginRequest{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	response, err := repo.LoginUser(request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "user not found", err.Error())
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	mock.ExpectQuery(`SELECT id, name, email, password FROM users WHERE email = \$1`).
		WithArgs("test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password"}).
			AddRow(1, "Test User", "test@example.com", hashedPassword))

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserLoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	response, err := repo.LoginUser(request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "invalid password", err.Error())
}

func TestLoginUser_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, name, email, password FROM users WHERE email = \$1`).
		WithArgs("test@example.com").
		WillReturnError(errors.New("query error"))

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserLoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	response, err := repo.LoginUser(request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "query error", err.Error())
}

func TestUpdateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewUserRepository(db)

	newPassword := "newpassword123"

	mock.ExpectExec(`UPDATE users SET name = \$1, email = \$2, password = \$3 WHERE id = \$4`).
		WithArgs("Updated User", "updated@example.com", sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	request := &userDTO.UserUpdateRequest{
		Name:     "Updated User",
		Email:    "updated@example.com",
		Password: newPassword,
	}

	err = repo.UpdateUser(request, 1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser_HashPasswordError(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserUpdateRequest{
		Name:     "Updated User",
		Email:    "updated@example.com",
		Password: string(make([]byte, bcrypt.MaxCost+1)),
	}

	err = repo.UpdateUser(request, 1)
	assert.Error(t, err)
}

func TestUpdateUser_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("newpassword123"), bcrypt.DefaultCost)

	mock.ExpectExec(`UPDATE users SET name = \$1, email = \$2, password = \$3 WHERE id = \$4`).
		WithArgs("Updated User", "updated@example.com", hashedPassword, 1).
		WillReturnError(errors.New("query error"))

	repo := repository.NewUserRepository(db)
	request := &userDTO.UserUpdateRequest{
		Name:     "Updated User",
		Email:    "updated@example.com",
		Password: "newpassword123",
	}

	err = repo.UpdateUser(request, 1)
	assert.Error(t, err)
}
func TestGetUserByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, name, email, password FROM users WHERE id = \$1`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password"}).
			AddRow(1, "Test User", "test@example.com", "hashedpassword"))

	repo := repository.NewUserRepository(db)

	response, err := repo.GetUserByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "Test User", response.Name)
	assert.Equal(t, "test@example.com", response.Email)
}

func TestGetUserByID_UserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, name, email, password FROM users WHERE id = \$1`).
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)

	repo := repository.NewUserRepository(db)

	response, err := repo.GetUserByID(999)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "user not found", err.Error())
}

func TestGetUserByID_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, name, email, password FROM users WHERE id = \$1`).
		WithArgs(1).
		WillReturnError(errors.New("query error"))

	repo := repository.NewUserRepository(db)

	response, err := repo.GetUserByID(1)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "query error", err.Error())
}

func TestUserList_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, email, name, password FROM users`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name", "password"}).
			AddRow(1, "test1@example.com", "Test User 1", "hashedpassword1").
			AddRow(2, "test2@example.com", "Test User 2", "hashedpassword2"))

	repo := repository.NewUserRepository(db)

	response, err := repo.UserList()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 2, response.Total)
	assert.Equal(t, "test1@example.com", response.Users[0].Email)
	assert.Equal(t, "Test User 1", response.Users[0].Name)
	assert.Equal(t, "test2@example.com", response.Users[1].Email)
	assert.Equal(t, "Test User 2", response.Users[1].Name)
}

func TestUserList_Empty(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, email, name, password FROM users`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "name", "password"}))

	repo := repository.NewUserRepository(db)

	response, err := repo.UserList()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 0, response.Total)
	assert.Empty(t, response.Users)
}

func TestUserList_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, email, name, password FROM users`).
		WillReturnError(errors.New("query error"))

	repo := repository.NewUserRepository(db)

	response, err := repo.UserList()
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "query error", err.Error())
}
