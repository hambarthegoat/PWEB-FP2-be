package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type commentsJSONItem struct {
    Content   string `json:"content"`
    Author    string `json:"author"`
    PostTitle string `json:"post_title"`
}

// SeedComments creates sample comments (including a reply) and returns them.
func SeedComments(commentRepo repository.CommentRepository, users map[string]*model.User, posts []*model.Post) ([]*model.Comment, error) {
    data, err := os.ReadFile("migrations/json/comments.json")
    if err != nil { return nil, err }
    var items []commentsJSONItem
    if err := json.Unmarshal(data, &items); err != nil { return nil, err }

    // build map post title -> *Post
    postMap := make(map[string]*model.Post)
    for _, p := range posts { postMap[p.Title] = p }

    comments := make([]*model.Comment, 0, len(items))
    for _, it := range items {
        author, ok := users[it.Author]
        if !ok { fmt.Println("warning: author not found for comment", it.Content); continue }
        post, ok := postMap[it.PostTitle]
        if !ok { fmt.Println("warning: post not found for comment", it.Content); continue }
        cm := &model.Comment{Content: it.Content, AuthorID: author.ID, PostID: post.ID}
        if err := commentRepo.Create(cm); err != nil { fmt.Println("warning creating comment:", err); continue }
        fmt.Println("created comment:", cm.ID)
        comments = append(comments, cm)
    }
    // optionally create a reply to the first comment (keeps same logic as before if available)
    if len(comments) > 0 {
        reply := &model.Comment{Content: "Thanks!", AuthorID: users["alice@example.com"].ID, PostID: comments[0].PostID, ParentID: &comments[0].ID}
        if err := commentRepo.Create(reply); err != nil { fmt.Println("warning creating reply:", err) } else { fmt.Println("created reply:", reply.ID) }
    }
    return comments, nil
}
