package dto

// Request body
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
	UpdatedAt string `json:"updated_at"`
}

// Response body
type NewsResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorId   int    `json:"user_id"`
	AuthorName string `json:"author_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type NewsListResponse struct {
	News  []NewsResponse `json:"news"`
	Total int            `json:"total"`
}
