package dto

type CreateCommentRequest struct {
	Content  string  `json:"content" binding:"required"`
	ParentID *string `json:"parent_id" binding:"omitempty,uuid"`
}

type CommentResponse struct {
	ID        string  `json:"id"`
	Content   string  `json:"content"`
	PostID    string  `json:"post_id"`
	ParentID  *string `json:"parent_id,omitempty"`
	AuthorID  string  `json:"author_id"`
	Score     int     `json:"score"`
	Replies   []*CommentResponse `json:"replies,omitempty"` // for tree
}

type CommentTree struct {
	CommentResponse
	Replies []CommentTree `json:"replies,omitempty"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}