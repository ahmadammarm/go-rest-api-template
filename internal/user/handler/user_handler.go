package handler

import (
	"strconv"

	"github.com/ahmadammarm/go-rest-api-template/internal/user/dto"
	userService "github.com/ahmadammarm/go-rest-api-template/internal/user/service"
	formvalidation "github.com/ahmadammarm/go-rest-api-template/pkg/form-validation"
	"github.com/ahmadammarm/go-rest-api-template/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService userService.UserService
	validation  *validator.Validate
}

func (handler *UserHandler) RegisterUser(context *fiber.Ctx) error {
	user := new(dto.UserRegisterRequest)
	if err := context.BodyParser(user); err != nil {
		return response.JSONResponse(context, 400, "Invalid Request", nil)
	}

	if err := handler.validation.Struct(user); err != nil {
		errorValidations := formvalidation.FormValidationError(err)
		return response.JSONResponse(context, 400, "Invalid Request", errorValidations)
	}

	if err := handler.userService.RegisterUser(user); err != nil {
		if err.Error() == "email already exists" {
			return response.JSONResponse(context, 409, "Email Already Exists", nil)
		}
	} else {
		return response.JSONResponse(context, 500, "Register User Failed", nil)
	}

	return response.JSONResponse(context, 200, "Register User Success", user)
}

func (handler *UserHandler) LoginUser(context *fiber.Ctx) error {

	loginRequest := new(dto.UserLoginRequest)

	if err := context.BodyParser(loginRequest); err != nil {
		return response.JSONResponse(context, 400, "Invalid Request", nil)
	}

	token, err := handler.userService.LoginUser(loginRequest)

	if err != nil {
		return response.JSONResponse(context, 401, "Login User Failed", nil)
	}

	return response.JSONResponse(context, 200, "Login User Success", token)
}

func (handler *UserHandler) UpdateUser(context *fiber.Ctx) error {
	user := new(dto.UserUpdateRequest)
	if err := context.BodyParser(user); err != nil {
		return response.JSONResponse(context, 400, "Invalid Request", nil)
	}

	if err := handler.validation.Struct(user); err != nil {
		errorValidations := formvalidation.FormValidationError(err)
		return response.JSONResponse(context, 400, "Invalid Request", errorValidations)
	}

	userId := context.Locals("user_id").(int)

	if err := handler.userService.UpdateUser(user, userId); err != nil {
		return response.JSONResponse(context, 500, "Update User Failed", nil)
	}

	return response.JSONResponse(context, 200, "Update User Success", nil)
}

func (handler *UserHandler) GetUserByID(context *fiber.Ctx) error {
	userIdString := context.Params("id")
	userId, err := strconv.Atoi(userIdString)

	if err != nil || userId < 1 {
		return response.JSONResponse(context, 400, "Invalid Request", nil)
	}

	user, err := handler.userService.GetUserByID(userId)
	if err != nil {
		return response.JSONResponse(context, 404, "User Not Found", nil)
	}

	return response.JSONResponse(context, 200, "Get User Success", user)
}

func (handler *UserHandler) UserList(context *fiber.Ctx) error {
	userList, err := handler.userService.UserList()
	if err != nil {
		return response.JSONResponse(context, 500, "User List Failed", nil)
	}

	return response.JSONResponse(context, 200, "Get User List Success", userList)
}

func (handler *UserHandler) UserRouters(router fiber.Router) {
	router.Post("/user/register", handler.RegisterUser)
	router.Post("/user/login", handler.LoginUser)
	router.Get("/users", handler.UserList)
	router.Get("/user/:id", handler.GetUserByID)
}

func NewUserHandler(userService userService.UserService, validation *validator.Validate) *UserHandler {
	return &UserHandler{
		userService: userService,
		validation:  validation,
	}
}
