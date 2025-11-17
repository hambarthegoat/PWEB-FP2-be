package migrate

import (
	"fmt"
	"log"

	"github.com/hambarthegoat/PWEB-FP2-be/config"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
	seedpkg "github.com/hambarthegoat/PWEB-FP2-be/migrations/seed"
)

// Seed orchestrates per-entity seeders.
func Seed() {
    cfg := config.Load()
    db := config.ConnectSupabase(cfg)

    userRepo := repository.NewUserRepository(db)
    communityRepo := repository.NewCommunityRepository(db)
    postRepo := repository.NewPostRepository(db)
    commentRepo := repository.NewCommentRepository(db)
    voteRepo := repository.NewVoteRepository(db)

    users, err := seedpkg.SeedUsers(userRepo)
    if err != nil { log.Fatal("seed users:", err) }

    comms, err := seedpkg.SeedCommunities(communityRepo, users)
    if err != nil { log.Fatal("seed communities:", err) }

    posts, err := seedpkg.SeedPosts(postRepo, users, comms)
    if err != nil { log.Fatal("seed posts:", err) }

    _, err = seedpkg.SeedComments(commentRepo, users, posts)
    if err != nil { log.Fatal("seed comments:", err) }

    if err := seedpkg.SeedVotes(voteRepo, users, posts); err != nil { log.Fatal("seed votes:", err) }

    fmt.Println("seeding completed")
}

