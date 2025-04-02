package model

import (
    "testing"
    "time"
)

func TestUserModel(t *testing.T) {
    t.Run("TestUserFields", func(t *testing.T) {
        createdAt := time.Now()
        user := User{
            ID:        1,
            Name:      "Admin Ammar",
            Email:     "admin@mail.com",
            Password:  "securepassword",
            CreatedAt: createdAt,
        }

        if user.ID != 1 {
            t.Errorf("expected ID to be 1, got %d", user.ID)
        }
        if user.Name != "Admin Ammar" {
            t.Errorf("expected Name to be 'Admin Ammar', got %s", user.Name)
        }
        if user.Email != "admin@mail.com" {
            t.Errorf("expected Email to be 'admin@mail.com', got %s", user.Email)
        }
        if user.Password != "securepassword" {
            t.Errorf("expected Password to be 'securepassword', got %s", user.Password)
        }
        if !user.CreatedAt.Equal(createdAt) {
            t.Errorf("expected CreatedAt to be %v, got %v", createdAt, user.CreatedAt)
        }
    })
}