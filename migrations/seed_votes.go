package migrate

import (
	"fmt"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

// SeedVotes adds sample votes and prints a sample score.
func SeedVotes(voteRepo repository.VoteRepository, users map[string]*model.User, posts []*model.Post) error {
    if len(posts) == 0 { return nil }
    v1 := &model.Vote{UserID: users["bob@example.com"].ID, TargetID: posts[0].ID, TargetType: model.VoteTargetPost, Value: 1}
    if err := voteRepo.Upsert(v1); err != nil { fmt.Println("warning voting:", err) } else { fmt.Println("voted post +1 by", v1.UserID) }
    v2 := &model.Vote{UserID: users["carol@example.com"].ID, TargetID: posts[0].ID, TargetType: model.VoteTargetPost, Value: 1}
    if err := voteRepo.Upsert(v2); err != nil { fmt.Println("warning voting:", err) } else { fmt.Println("voted post +1 by", v2.UserID) }

    score, err := voteRepo.CalculateScore(posts[0].ID, model.VoteTargetPost)
    if err == nil { fmt.Println("post score:", score) }
    return nil
}
