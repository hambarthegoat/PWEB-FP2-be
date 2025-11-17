package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type postsJSONItem struct {
    Title     string `json:"title"`
    Content   string `json:"content"`
    Author    string `json:"author"`
    Community string `json:"community"`
}

// SeedPosts creates sample posts and returns them.
func SeedPosts(postRepo repository.PostRepository, users map[string]*model.User, comms []*model.Community) ([]*model.Post, error) {
    data, err := os.ReadFile("migrations/json/posts.json")
    if err != nil { return nil, err }
    var items []postsJSONItem
    if err := json.Unmarshal(data, &items); err != nil { return nil, err }

    // build map community name -> *Community
    commMap := make(map[string]*model.Community)
    for _, c := range comms { commMap[c.Name] = c }

    posts := make([]*model.Post, 0, len(items))
    for _, it := range items {
        author, ok := users[it.Author]
        if !ok { fmt.Println("warning: author not found for post", it.Title); continue }
        comm, ok := commMap[it.Community]
        if !ok { fmt.Println("warning: community not found for post", it.Title); continue }
        p := &model.Post{Title: it.Title, Content: it.Content, AuthorID: author.ID, CommunityID: comm.ID}
        if err := postRepo.Create(p); err != nil { fmt.Println("warning creating post:", err); continue }
        fmt.Println("created post:", p.ID, p.Title)
        posts = append(posts, p)
    }
    return posts, nil
}
