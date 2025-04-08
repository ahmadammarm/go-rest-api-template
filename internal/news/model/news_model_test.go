package model

import (
    "encoding/json"
    "testing"
)

func TestNewsJSONMarshalling(t *testing.T) {
    news := News{
        ID:        1,
        Title:     "Test Title",
        Content:   "Test Content",
        AuthorId:  6,
        CreatedAt: "2023-01-01T00:00:00Z",
        UpdatedAt: "2023-01-02T00:00:00Z",
    }

    data, err := json.Marshal(news)
    if err != nil {
        t.Fatalf("Failed to marshal News: %v", err)
    }

    var unmarshalledNews News
    err = json.Unmarshal(data, &unmarshalledNews)
    if err != nil {
        t.Fatalf("Failed to unmarshal News: %v", err)
    }

    if news != unmarshalledNews {
        t.Errorf("Expected unmarshalled News to be %+v, got %+v", news, unmarshalledNews)
    }
}

func TestNewsFieldValidation(t *testing.T) {
    news := News{
        ID:        1,
        Title:     "Test Title",
        Content:   "Test Content",
        AuthorId:  123,
        CreatedAt: "2023-01-01T00:00:00Z",
        UpdatedAt: "2023-01-02T00:00:00Z",
    }

    if news.ID != 1 {
        t.Errorf("Expected ID to be 1, got %d", news.ID)
    }
    if news.Title != "Test Title" {
        t.Errorf("Expected Title to be 'Test Title', got '%s'", news.Title)
    }
    if news.Content != "Test Content" {
        t.Errorf("Expected Content to be 'Test Content', got '%s'", news.Content)
    }
    if news.AuthorId != 123 {
        t.Errorf("Expected AuthorId to be 123, got %d", news.AuthorId)
    }
    if news.CreatedAt != "2023-01-01T00:00:00Z" {
        t.Errorf("Expected CreatedAt to be '2023-01-01T00:00:00Z', got '%s'", news.CreatedAt)
    }
    if news.UpdatedAt != "2023-01-02T00:00:00Z" {
        t.Errorf("Expected UpdatedAt to be '2023-01-02T00:00:00Z', got '%s'", news.UpdatedAt)
    }
}