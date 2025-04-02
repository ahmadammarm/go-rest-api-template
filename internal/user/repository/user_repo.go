package repository

import (
	"database/sql"
	"errors"

	userDTO "github.com/ahmadammarm/go-rest-api-template/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	RegisterUser(user *userDTO.UserRegisterRequest) error
	LoginUser(user *userDTO.UserLoginRequest) (*userDTO.UserJWTResponse, error)
	LogoutUser(user *userDTO.UserLogoutRequest) error
	UpdateUser(user *userDTO.UserUpdateRequest, id int) error
	GetUserByID(userId int) (*userDTO.UserResponse, error)
	UserList() (*userDTO.UserListResponse, error)
	IsUserExist(userId int) (bool, error)
}

type userRepoImpl struct {
	db *sql.DB
}

func (repository *userRepoImpl) RegisterUser(user *userDTO.UserRegisterRequest) error {
	tx, err := repository.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if pan := recover(); pan != nil {
			tx.Rollback()
			panic(pan)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query := `INSERT INTO users (name, email, password, role, created_at) VALUES ($1, $2, $3, $4, $5)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	err = tx.QueryRow(query, user.Name, user.Email, hashedPassword, user.Role, user.CreatedAt).Scan(&user.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil

}

func (repository *userRepoImpl) LoginUser(user *userDTO.UserLoginRequest) (*userDTO.UserJWTResponse, error) {
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	jwtUser := &userDTO.UserJWTResponse{}

	err := repository.db.QueryRow(query, user.Email).Scan(&jwtUser.ID, &jwtUser.Name, &jwtUser.Email, &jwtUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(jwtUser.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return jwtUser, nil

}

func (repository *userRepoImpl) LogoutUser(user *userDTO.UserLogoutRequest) error {
	return nil
}

func (repository *userRepoImpl) UpdateUser(user *userDTO.UserUpdateRequest, id int) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = repository.db.Exec(query, user.Name, user.Email, hashedPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *userRepoImpl) GetUserByID(userId int) (*userDTO.UserResponse, error) {
	query := `SELECT id, name, email, password, created_at, role FROM users WHERE id = $1`
	user := &userDTO.UserResponse{}

	err := repository.db.QueryRow(query, userId).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (repository *userRepoImpl) UserList() (*userDTO.UserListResponse, error) {
	query := `SELECT id, name, email, password, created_at, role FROM users`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []userDTO.UserResponse
	for rows.Next() {
		user := userDTO.UserResponse{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &userDTO.UserListResponse{Users: users}, nil
}

func (repository *userRepoImpl) IsUserExist(userId int) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE id = $1`
	var count int
	err := repository.db.QueryRow(query, userId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func NewUserRepository(db *sql.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}
