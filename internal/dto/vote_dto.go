package dto

type VoteRequest struct {
	Value int8 `json:"value" binding:"required,oneof=1 -1"`
}

type VoteResult struct {
	PostID    *string `json:"post_id,omitempty"`
	CommentID *string `json:"comment_id,omitempty"`
	Score     int     `json:"score"`
}