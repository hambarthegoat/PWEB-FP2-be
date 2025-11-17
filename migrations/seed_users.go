package migrate

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

// SeedUsers creates sample users and returns a map by email.
func SeedUsers(userRepo repository.UserRepository) (map[string]*model.User, error) {
    passwords := map[string]string{
        "alice@example.com": "password123",
        "bob@example.com":   "secret456",
        "carol@example.com": "hunter2",
    }
    users := make(map[string]*model.User)
    for email, pass := range passwords {
        existing, _ := userRepo.FindByEmail(email)
        if existing != nil {
            users[email] = existing
            fmt.Println("user exists:", existing.ID)
            continue
        }
        hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
        if err != nil { return nil, err }
        hs := string(hashed)
        u := &model.User{Email: email, Password: &hs}
        if err := userRepo.Create(u); err != nil { return nil, err }
        users[email] = u
        fmt.Println("created user:", u.ID)
    }
    return users, nil
}
