package dto

// Request menyimpan parameter query pagination.
type PaginationRequest struct {
	Page   int
	Limit  int
	Search string
	Sort   string
	Order  string
}

// Response menyimpan informasi pagination yang dikirim ke frontend.
type PaginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}
