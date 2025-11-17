package migrate

import (
	"fmt"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

// SeedPosts creates sample posts and returns them.
func SeedPosts(postRepo repository.PostRepository, users map[string]*model.User, comms []*model.Community) ([]*model.Post, error) {
    posts := []*model.Post{
        {Title: "Welcome to Gophers", Content: "Hello from Alice", AuthorID: users["alice@example.com"].ID, CommunityID: comms[0].ID},
        {Title: "Programming tips", Content: "Use functions", AuthorID: users["bob@example.com"].ID, CommunityID: comms[1].ID},
        {Title: "Concurrency in Go", Content: "Goroutines are lightweight", AuthorID: users["carol@example.com"].ID, CommunityID: comms[0].ID},
    }
    for _, p := range posts {
        if err := postRepo.Create(p); err != nil {
            fmt.Println("warning creating post:", err)
            continue
        }
        fmt.Println("created post:", p.ID, p.Title)
    }
    return posts, nil
}
