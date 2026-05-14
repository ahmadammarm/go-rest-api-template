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
	UpdateUser(name string, email string, hashedPassword string, id int) error
	GetUserByID(userId int) (*userDTO.UserResponse, error)
	IsEmailExists(email string) (bool, error)
	IsEmailTakenByOther(email string, id int) (bool, error)
	UserList() (*userDTO.UserListResponse, error)
}

type userRepoImpl struct {
	db *sql.DB
}

func (repository *userRepoImpl) RegisterUser(user *userDTO.UserRegisterRequest) (err error) {
	tx, err := repository.db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if pan := recover(); pan != nil {
			_ = tx.Rollback()
			panic(pan)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query := `INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	err = tx.QueryRow(query, user.Email, user.Name, hashedPassword).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil

}

func (repository *userRepoImpl) LoginUser(user *userDTO.UserLoginRequest) (*userDTO.UserJWTResponse, error) {
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	jwtUser := &userDTO.UserJWTResponse{}
	var hashedPassword string

	err := repository.db.QueryRow(query, user.Email).Scan(&jwtUser.ID, &jwtUser.Name, &jwtUser.Email, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return jwtUser, nil

}

func (repository *userRepoImpl) UpdateUser(name string, email string, hashedPassword string, id int) error {
	query := `UPDATE users SET name = $1, email = $2, password = CASE WHEN $3 <> '' THEN $3 ELSE password END WHERE id = $4`

	_, err := repository.db.Exec(query, name, email, hashedPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *userRepoImpl) GetUserByID(userId int) (*userDTO.UserResponse, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	user := &userDTO.UserResponse{}

	err := repository.db.QueryRow(query, userId).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (repository *userRepoImpl) UserList() (*userDTO.UserListResponse, error) {
	query := `SELECT id, email, name FROM users`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []userDTO.UserResponse
	for rows.Next() {
		user := userDTO.UserResponse{}
		err := rows.Scan(&user.ID, &user.Email, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	total := len(users)

	return &userDTO.UserListResponse{Users: users, Total: total}, nil
}

func (repository *userRepoImpl) IsEmailTakenByOther(email string, id int) (bool, error) {
	query := `SELECT COUNT(1) FROM users WHERE email = $1 AND id <> $2`
	var count int
	err := repository.db.QueryRow(query, email, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repository *userRepoImpl) IsEmailExists(email string) (bool, error) {
    query := `SELECT COUNT(1) FROM users WHERE email = $1`
    var count int
    err := repository.db.QueryRow(query, email).Scan(&count)
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
