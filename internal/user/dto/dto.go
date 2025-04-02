package dto

import "time"

// Request
type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin user" default:"user"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLogoutRequest struct {
	Token string `json:"token" validate:"required"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6,max=20"`
}

// Response
type UserJWTResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
}
