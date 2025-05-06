package service

import (
	"os"
	"time"

	"errors"

	idgenerate "github.com/ahmadammarm/go-rest-api-template/pkg/id-generate"
	userDTO "github.com/ahmadammarm/go-rest-api-template/internal/user/dto"
	userRepo "github.com/ahmadammarm/go-rest-api-template/internal/user/repository"
	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	RegisterUser(user *userDTO.UserRegisterRequest) error
	LoginUser(user *userDTO.UserLoginRequest) (string, error)
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

	user.ID = idgenerate.GenerateUniqueID()

	return service.userRepo.RegisterUser(user)
}

func (service *userServiceImpl) LoginUser(user *userDTO.UserLoginRequest) (string, error) {
	dbUser, err := service.userRepo.LoginUser(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"apps":    "go-rest-api-template",
		"user_id": dbUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secretToken := os.Getenv("JWT_SECRET_KEY")
	if secretToken == "" {
		return "", errors.New("JWT_SECRET is not set in environment variables")
	}

	stringToken, err := token.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}

	return stringToken, nil
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
