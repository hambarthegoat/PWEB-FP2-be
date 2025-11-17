package dto

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,max=300"`
	Content string `json:"content" binding:"required"`
}

type PostResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content,omitempty"`
	CommunityID string `json:"community_id"`
	AuthorID    string `json:"author_id"`
	Score       int    `json:"score"`
}

type PostListItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Score int    `json:"score"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"omitempty,max=300"`
	Content string `json:"content" binding:"omitempty"`
}