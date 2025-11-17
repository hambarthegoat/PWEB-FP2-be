package dto

type CreateCommunityRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

type CommunityResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	CreatorID   string `json:"creator_id"`
}

type UpdateCommunityRequest struct {
	Description string `json:"description" binding:"omitempty,max=500"`
}