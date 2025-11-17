package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type communitiesJSONItem struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Creator     string `json:"creator"`
}

// SeedCommunities creates sample communities and returns them.
func SeedCommunities(communityRepo repository.CommunityRepository, users map[string]*model.User) ([]*model.Community, error) {
    data, err := os.ReadFile("migrations/json/communities.json")
    if err != nil { return nil, err }
    var items []communitiesJSONItem
    if err := json.Unmarshal(data, &items); err != nil { return nil, err }

    comms := make([]*model.Community, 0, len(items))
    for _, it := range items {
        creator, ok := users[it.Creator]
        if !ok { fmt.Println("warning: creator not found for community", it.Name); continue }
        c := &model.Community{Name: it.Name, Description: it.Description, CreatorID: creator.ID}
        if err := communityRepo.Create(c); err != nil { fmt.Println("warning creating community (may already exist):", err); continue }
        fmt.Println("created community:", c.ID, c.Name)
        comms = append(comms, c)
    }
    return comms, nil
}
