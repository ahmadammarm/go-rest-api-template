package service

import (
	"os"
	"time"

	"errors"

	userDTO "github.com/ahmadammarm/go-rest-api-template/internal/user/dto"
	userRepo "github.com/ahmadammarm/go-rest-api-template/internal/user/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *userDTO.UserRegisterRequest) error
	LoginUser(user *userDTO.UserLoginRequest) (any, error)
	UpdateUser(user *userDTO.UserUpdateRequest, id int) error
	GetUserByID(userId int) (*userDTO.UserResponse, error)
	UserList() (*userDTO.UserListResponse, error)
}

type userServiceImpl struct {
	userRepo  userRepo.UserRepo
	jwtSecret string
}

func NewUserService(userRepo userRepo.UserRepo) UserService {
	return &userServiceImpl{
		userRepo:  userRepo,
		jwtSecret: os.Getenv("JWT_SECRET_KEY"),
	}
}

func (service *userServiceImpl) RegisterUser(user *userDTO.UserRegisterRequest) error {

	if exists, err := service.userRepo.IsEmailExists(user.Email); err != nil {
		return err
	} else if exists {
		return errors.New("email already exists")
	}

	return service.userRepo.RegisterUser(user)
}

func (service *userServiceImpl) LoginUser(user *userDTO.UserLoginRequest) (any, error) {
	dbUser, err := service.userRepo.LoginUser(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"apps":    "go-rest-api-template",
		"user_id": dbUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	if service.jwtSecret == "" {
		return "", errors.New("JWT_SECRET is not set in environment variables")
	}

	stringToken, err := token.SignedString([]byte(service.jwtSecret))
	if err != nil {
		return "", err
	}

	response := userDTO.UserJWTResponse{
		ID:    dbUser.ID,
		Name:  dbUser.Name,
		Email: dbUser.Email,
		Token: stringToken,
	}

	return response, nil
}

func (service *userServiceImpl) UpdateUser(user *userDTO.UserUpdateRequest, id int) error {
	if exists, err := service.userRepo.IsEmailTakenByOther(user.Email, id); err != nil {
		return err
	} else if exists {
		return errors.New("email already exists")
	}

	var hashedPassword string
	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		hashedPassword = string(hash)
	}

	return service.userRepo.UpdateUser(user.Name, user.Email, hashedPassword, id)
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
