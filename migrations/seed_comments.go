package migrate

import (
	"fmt"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

// SeedComments creates sample comments (including a reply) and returns them.
func SeedComments(commentRepo repository.CommentRepository, users map[string]*model.User, posts []*model.Post) ([]*model.Comment, error) {
    comments := []*model.Comment{
        {Content: "Great post!", AuthorID: users["bob@example.com"].ID, PostID: posts[0].ID},
        {Content: "I agree", AuthorID: users["carol@example.com"].ID, PostID: posts[0].ID},
    }
    for _, cm := range comments {
        if err := commentRepo.Create(cm); err != nil { fmt.Println("warning creating comment:", err); continue }
        fmt.Println("created comment:", cm.ID)
    }
    // reply to first comment
    if len(comments) > 0 {
        reply := &model.Comment{Content: "Thanks!", AuthorID: users["alice@example.com"].ID, PostID: posts[0].ID, ParentID: &comments[0].ID}
        if err := commentRepo.Create(reply); err != nil { fmt.Println("warning creating reply:", err) } else { fmt.Println("created reply:", reply.ID) }
    }
    return comments, nil
}
