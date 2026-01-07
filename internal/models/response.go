package models

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

type UserWithAgeResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age"`
}

type ErrorDetail struct {
	Message   string `json:"message"`
	Code      string `json:"code"`
	RequestID string `json:"request_id,omitempty"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type SimpleErrorResponse struct {
	Error string `json:"error"`
}

type PaginationMeta struct {
	Total       int64 `json:"total"`
	Page        int   `json:"page"`
	Limit       int   `json:"limit"`
	TotalPages  int   `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
	HasPrevious bool  `json:"has_previous"`
}

type PaginatedUsersResponse struct {
	Data       []UserWithAgeResponse `json:"data"`
	Pagination PaginationMeta        `json:"pagination"`
}
