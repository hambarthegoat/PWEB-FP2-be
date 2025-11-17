package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type votesJSONItem struct {
    User      string `json:"user"`
    PostTitle string `json:"post_title"`
    Value     int8   `json:"value"`
}

// SeedVotes adds sample votes and prints a sample score.
func SeedVotes(voteRepo repository.VoteRepository, users map[string]*model.User, posts []*model.Post) error {
    data, err := os.ReadFile("migrations/json/votes.json")
    if err != nil { return err }
    var items []votesJSONItem
    if err := json.Unmarshal(data, &items); err != nil { return err }

    postMap := make(map[string]*model.Post)
    for _, p := range posts { postMap[p.Title] = p }

    for _, it := range items {
        user, ok := users[it.User]
        if !ok { fmt.Println("warning: vote user not found:", it.User); continue }
        post, ok := postMap[it.PostTitle]
        if !ok { fmt.Println("warning: vote post not found:", it.PostTitle); continue }
        v := &model.Vote{UserID: user.ID, TargetID: post.ID, TargetType: model.VoteTargetPost, Value: it.Value}
        if err := voteRepo.Upsert(v); err != nil { fmt.Println("warning voting:", err); continue }
        fmt.Println("voted post", it.Value, "by", v.UserID)
    }

    // optionally print score for first post
    if len(posts) > 0 {
        if score, err := voteRepo.CalculateScore(posts[0].ID, model.VoteTargetPost); err == nil {
            fmt.Println("post score:", score)
        }
    }
    return nil
}
