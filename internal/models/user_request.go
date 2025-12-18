package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}
type DeleteUserRequest struct {
	ID int64 `json:"id" validate:"required,gt=0"`
}
type GetUserRequest struct {
	ID int64 `json:"id" validate:"required,gt=0"`
}
type ListUsersRequest struct{}

type UpdateUserRequest struct {
	ID   int64  `json:"id" validate:"required,gt=0"`
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}
