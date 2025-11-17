package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type usersJSONItem struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// SeedUsers creates sample users and returns a map by email.
func SeedUsers(userRepo repository.UserRepository) (map[string]*model.User, error) {
    data, err := os.ReadFile("migrations/json/users.json")
    if err != nil { return nil, err }
    var items []usersJSONItem
    if err := json.Unmarshal(data, &items); err != nil { return nil, err }

    users := make(map[string]*model.User)
    for _, it := range items {
        existing, _ := userRepo.FindByEmail(it.Email)
        if existing != nil {
            users[it.Email] = existing
            fmt.Println("user exists:", existing.ID)
            continue
        }
        hashed, err := bcrypt.GenerateFromPassword([]byte(it.Password), bcrypt.DefaultCost)
        if err != nil { return nil, err }
        hs := string(hashed)
        u := &model.User{Email: it.Email, Password: &hs}
        if err := userRepo.Create(u); err != nil { return nil, err }
        users[it.Email] = u
        fmt.Println("created user:", u.ID)
    }
    return users, nil
}
