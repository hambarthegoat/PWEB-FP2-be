package migrate

import (
	"fmt"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

// SeedCommunities creates sample communities and returns them.
func SeedCommunities(communityRepo repository.CommunityRepository, users map[string]*model.User) ([]*model.Community, error) {
    comms := []*model.Community{
        {Name: "golang", Description: "Gophers unite", CreatorID: users["alice@example.com"].ID},
        {Name: "programming", Description: "General programming", CreatorID: users["bob@example.com"].ID},
    }
    for _, c := range comms {
        if err := communityRepo.Create(c); err != nil {
            fmt.Println("warning creating community (may already exist):", err)
            continue
        }
        fmt.Println("created community:", c.ID, c.Name)
    }
    return comms, nil
}
