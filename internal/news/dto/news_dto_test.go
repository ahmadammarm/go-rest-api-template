package dto_test

import (
	"testing"

	"github.com/ahmadammarm/go-rest-api-template/internal/news/dto"
	"github.com/go-playground/validator/v10"
)

func compareNewsListResponse(a, b dto.NewsListResponse) bool {
	if a.Total != b.Total {
		return false
	}
	if len(a.News) != len(b.News) {
		return false
	}
	for i := range a.News {
		if a.News[i] != b.News[i] {
			return false
		}
	}
	return true
}

func TestNewsCreateRequestValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		request dto.NewsCreateRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			request: dto.NewsCreateRequest{
				ID:       1,
				Title:    "Test Title",
				Content:  "Test Content",
				AuthorId: 123,
			},
			wantErr: false,
		},
		{
			name: "Missing title",
			request: dto.NewsCreateRequest{
				ID:       1,
				Content:  "Test Content",
				AuthorId: 123,
			},
			wantErr: true,
		},
		{
			name: "Missing content",
			request: dto.NewsCreateRequest{
				ID:       1,
				Title:    "Test Title",
				AuthorId: 123,
			},
			wantErr: true,
		},
		{
			name: "Missing author ID",
			request: dto.NewsCreateRequest{
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
		request dto.NewsUpdateRequest
		wantErr bool
	}{
		{
			name: "Valid request",
			request: dto.NewsUpdateRequest{
				ID:       1,
				Title:    "Test Title",
				Content:  "Test Content",
				AuthorId: 123,
			},
			wantErr: false,
		},
		{
			name: "Missing title",
			request: dto.NewsUpdateRequest{
				ID:       1,
				Content:  "Test Content",
				AuthorId: 123,
			},
			wantErr: true,
		},
		{
			name: "Missing content",
			request: dto.NewsUpdateRequest{
				ID:       1,
				Title:    "Test Title",
				AuthorId: 123,
			},
			wantErr: true,
		},
		{
			name: "Missing author ID",
			request: dto.NewsUpdateRequest{
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
		response dto.NewsResponse
		want     dto.NewsResponse
	}{
		{
			name: "Valid response",
			response: dto.NewsResponse{
				ID:       1,
				Title:    "Test Title",
				Content:  "Test Content",
				AuthorId: 123,
			},
			want: dto.NewsResponse{
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

func TestNewsListResponse(t *testing.T) {
	tests := []struct {
		name     string
		response dto.NewsListResponse
		want     dto.NewsListResponse
	}{
		{
			name: "Valid response",
			response: dto.NewsListResponse{
				News: []dto.NewsResponse{
					{
						ID:       1,
						Title:    "Test Title 1",
						Content:  "Test Content 1",
						AuthorId: 123,
					},
					{
						ID:       2,
						Title:    "Test Title 2",
						Content:  "Test Content 2",
						AuthorId: 456,
					},
				},
				Total: 2,
			},
			want: dto.NewsListResponse{
				News: []dto.NewsResponse{
					{
						ID:       1,
						Title:    "Test Title 1",
						Content:  "Test Content 1",
						AuthorId: 123,
					},
					{
						ID:       2,
						Title:    "Test Title 2",
						Content:  "Test Content 2",
						AuthorId: 456,
					},
				},
				Total: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !compareNewsListResponse(tt.response, tt.want) {
				t.Errorf("NewsListResponse = %v, want %v", tt.response, tt.want)
			}
		})
	}
}
