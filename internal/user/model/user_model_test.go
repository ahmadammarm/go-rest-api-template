package model_test

import (
	"testing"

	"github.com/ahmadammarm/go-rest-api-template/internal/user/model"
)

func TestUser(t *testing.T) {
	user := model.User{
		ID:       1,
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "securepassword",
	}

	if user.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", user.ID)
	}

	if user.Email != "test@example.com" {
		t.Errorf("expected Email to be 'test@example.com', got %s", user.Email)
	}

	if user.Name != "Test User" {
		t.Errorf("expected Name to be 'Test User', got %s", user.Name)
	}

	if user.Password != "securepassword" {
		t.Errorf("expected Password to be 'securepassword', got %s", user.Password)
	}
}
