package dto

import (
    "testing"

    "github.com/go-playground/validator/v10"
)

func TestNewsCreateRequestValidation(t *testing.T) {
    validate := validator.New()

    tests := []struct {
        name    string
        request NewsCreateRequest
        wantErr bool
    }{
        {
            name: "Valid request",
            request: NewsCreateRequest{
                ID:       1,
                Title:    "Test Title",
                Content:  "Test Content",
                AuthorId: 123,
            },
            wantErr: false,
        },
        {
            name: "Missing title",
            request: NewsCreateRequest{
                ID:       1,
                Content:  "Test Content",
                AuthorId: 123,
            },
            wantErr: true,
        },
        {
            name: "Missing content",
            request: NewsCreateRequest{
                ID:       1,
                Title:    "Test Title",
                AuthorId: 123,
            },
            wantErr: true,
        },
        {
            name: "Missing author ID",
            request: NewsCreateRequest{
                ID:      1,
                Title:   "Test Title",
                Content: "Test Content",
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validate.Struct(tt.request)
            if (err != nil) != tt.wantErr {
                t.Errorf("Validation error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestNewsUpdateRequestValidation(t *testing.T) {
    validate := validator.New()

    tests := []struct {
        name    string
        request NewsUpdateRequest
        wantErr bool
    }{
        {
            name: "Valid request",
            request: NewsUpdateRequest{
                ID:       1,
                Title:    "Test Title",
                Content:  "Test Content",
                AuthorId: 123,
            },
            wantErr: false,
        },
        {
            name: "Missing title",
            request: NewsUpdateRequest{
                ID:      1,
                Content: "Test Content",
                AuthorId: 123,
            },
            wantErr: true,
        },
        {
            name: "Missing content",
            request: NewsUpdateRequest{
                ID:      1,
                Title:   "Test Title",
                AuthorId: 123,
            },
            wantErr: true,
        },
        {
            name: "Missing author ID",
            request: NewsUpdateRequest{
                ID:      1,
                Title:   "Test Title",
                Content: "Test Content",
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validate.Struct(tt.request)
            if (err != nil) != tt.wantErr {
                t.Errorf("Validation error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestNewsResponse(t *testing.T) {
    tests := []struct {
        name     string
        response NewsResponse
        want     NewsResponse
    }{
        {
            name: "Valid response",
            response: NewsResponse{
                ID:       1,
                Title:    "Test Title",
                Content:  "Test Content",
                AuthorId: 123,
            },
            want: NewsResponse{
                ID:       1,
                Title:    "Test Title",
                Content:  "Test Content",
                AuthorId: 123,
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.response != tt.want {
                t.Errorf("NewsResponse = %v, want %v", tt.response, tt.want)
            }
        })
    }
}