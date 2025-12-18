package models

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}

type PaginatedResponse[T any] struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
	Data  []T   `json:"data"`
}
