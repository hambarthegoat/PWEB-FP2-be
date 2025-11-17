package dto

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Email string `json:"email" binding:"omitempty,email"`
}