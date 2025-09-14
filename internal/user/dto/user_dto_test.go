package dto_test

import (
	"testing"

	"github.com/ahmadammarm/go-rest-api-template/internal/user/dto"
)

func TestUserRegisterRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    dto.UserRegisterRequest
		expected dto.UserRegisterRequest
	}{
		{
			name: "Valid Input",
			input: dto.UserRegisterRequest{
				ID:       1,
				Email:    "testuser@mail.com",
				Name:     "Saya Test User",
				Password: "password123",
			},
			expected: dto.UserRegisterRequest{
				ID:       1,
				Email:    "testuser@mail.com",
				Name:     "Saya Test User",
				Password: "password123",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input != test.expected {
				t.Errorf("expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func TestUserLoginRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    dto.UserLoginRequest
		expected dto.UserLoginRequest
	}{
		{
			name: "Valid Input",
			input: dto.UserLoginRequest{
				Email:    "testuser@mail.com",
				Password: "testpassword",
			},
			expected: dto.UserLoginRequest{
				Email:    "testuser@mail.com",
				Password: "testpassword",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input != test.expected {
				t.Errorf("expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func UserLogoutRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    dto.UserLogoutRequest
		expected dto.UserLogoutRequest
	}{
		{
			name: "Valid Input",
			input: dto.UserLogoutRequest{
				Token: "testtoken",
			},
			expected: dto.UserLogoutRequest{
				Token: "testtoken",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input != test.expected {
				t.Errorf("expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func TestUserUpdateRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    dto.UserUpdateRequest
		expected dto.UserUpdateRequest
	}{
		{
			name: "Valid Input",
			input: dto.UserUpdateRequest{
				Name:     "Updated User",
				Email:    "updateduser@mail.com",
				Password: "newpassword123",
			},
			expected: dto.UserUpdateRequest{
				Name:     "Updated User",
				Email:    "updateduser@mail.com",
				Password: "newpassword123",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input != test.expected {
				t.Errorf("expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func TestUserJWTResponse(t *testing.T) {
    tests := []struct {
        name     string
        input    dto.UserJWTResponse
        expected dto.UserJWTResponse
    }{
        {
            name: "Valid Input",
            input: dto.UserJWTResponse{
                ID:       1,
                Name:     "Test User",
                Email:    "testuser@mail.com",
                // Password: "password123",
                Token:    "testtoken123",
            },
            expected: dto.UserJWTResponse{
                ID:       1,
                Name:     "Test User",
                Email:    "testuser@mail.com",
                // Password: "password123",
                Token:    "testtoken123",
            },
        },
        {
            name: "Empty Fields",
            input: dto.UserJWTResponse{
                ID:       0,
                Name:     "",
                Email:    "",
                // Password: "",
                Token:    "",
            },
            expected: dto.UserJWTResponse{
                ID:       0,
                Name:     "",
                Email:    "",
                // Password: "",
                Token:    "",
            },
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            if test.input != test.expected {
                t.Errorf("expected %v, got %v", test.expected, test.input)
            }
        })
    }
}

func TestUserResponse(t *testing.T) {
    tests := []struct {
        name     string
        input    dto.UserResponse
        expected dto.UserResponse
    }{
        {
            name: "Valid Input",
            input: dto.UserResponse{
                ID:       1,
                Name:     "Test User",
                Email: "testuser@mail.com",
            },
            expected: dto.UserResponse{
                ID:       1,
                Name:     "Test User",
                Email: "testuser@mail.com",
            },
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            if test.input != test.expected {
                t.Errorf("expected %v, got %v", test.expected, test.input)
            }
        })
    }
}


func TestUserListResponse(t *testing.T) {
    tests := []struct {
        name     string
        input    dto.UserListResponse
        expected dto.UserListResponse
    }{
        {
            name: "Valid Input",
            input: dto.UserListResponse{
                Users: []dto.UserResponse{
                    {
                        ID:       1,
                        Name:     "User One",
                        Email:    "userone@mail.com",
                    },
                    {
                        ID:       2,
                        Name:     "User Two",
                        Email:    "usertwo@mail.com",
                    },
                },
                Total: 2,
            },
            expected: dto.UserListResponse{
                Users: []dto.UserResponse{
                    {
                        ID:       1,
                        Name:     "User One",
                        Email:    "userone@mail.com",
                    },
                    {
                        ID:       2,
                        Name:     "User Two",
                        Email:    "usertwo@mail.com",
                    },
                },
                Total: 2,
            },
        },
        {
            name: "Empty List",
            input: dto.UserListResponse{
                Users: []dto.UserResponse{},
                Total: 0,
            },
            expected: dto.UserListResponse{
                Users: []dto.UserResponse{},
                Total: 0,
            },
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            if len(test.input.Users) != len(test.expected.Users) || test.input.Total != test.expected.Total {
                t.Errorf("expected %v, got %v", test.expected, test.input)
            }
            for i, user := range test.input.Users {
                if user != test.expected.Users[i] {
                    t.Errorf("expected user %v, got %v", test.expected.Users[i], user)
                }
            }
        })
    }
}