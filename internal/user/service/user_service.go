package service

import (
	"os"
	"time"

	"errors"
	userDTO "github.com/ahmadammarm/go-rest-api-template/internal/user/dto"
	userRepo "github.com/ahmadammarm/go-rest-api-template/internal/user/repository"
	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	RegisterUser(user *userDTO.UserRegisterRequest) error
	LoginUser(user *userDTO.UserLoginRequest) (*userDTO.UserJWTResponse, error)
	LogoutUser(user *userDTO.UserLogoutRequest) error
	UpdateUser(user *userDTO.UserUpdateRequest, id int) error
	GetUserByID(userId int) (*userDTO.UserResponse, error)
	UserList() (*userDTO.UserListResponse, error)
}

type userServiceImpl struct {
	userRepo userRepo.UserRepo
}

func NewUserService(userRepo userRepo.UserRepo) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (service *userServiceImpl) RegisterUser(user *userDTO.UserRegisterRequest) error {
	return service.userRepo.RegisterUser(user)
}

func (service *userServiceImpl) LoginUser(user *userDTO.UserLoginRequest) (*userDTO.UserJWTResponse, error) {
	dbUser, err := service.userRepo.LoginUser(user)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"apps":    "go-rest-api-template",
		"user_id": dbUser.ID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	secretToken := os.Getenv("JWT_SECRET")

	if secretToken == "" {
		return nil, errors.New("JWT secret key not found")
	}

	tokenString, err := token.SignedString([]byte(secretToken))
	if err != nil {
		return nil, err
	}

	dbUser.Token = tokenString

	return dbUser, nil
}

func (service *userServiceImpl) LogoutUser(user *userDTO.UserLogoutRequest) error {
	return service.userRepo.LogoutUser(user)
}

func (service *userServiceImpl) UpdateUser(user *userDTO.UserUpdateRequest, id int) error {
	return service.userRepo.UpdateUser(user, id)
}

func (service *userServiceImpl) GetUserByID(userId int) (*userDTO.UserResponse, error) {
    user, err := service.userRepo.GetUserByID(userId)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (service *userServiceImpl) UserList() (*userDTO.UserListResponse, error) {
	users, err := service.userRepo.UserList()
	if err != nil {
		return nil, err
	}

	return users, nil
}
