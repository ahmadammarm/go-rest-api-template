package dto

// Request
type NewsCreateRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	AuthorId  int    `json:"user_id" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type NewsUpdateRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	AuthorId  int    `json:"user_id" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type NewsDeleteRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorId  int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


// Response
type NewsResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorId  int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type NewsByAuthorId struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    AuthorId  int    `json:"user_id"`
}

type NewsListResponse struct {
	News  []NewsResponse `json:"news"`
	Total int            `json:"total"`
}
