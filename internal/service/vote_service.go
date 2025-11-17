package service

import (
	"errors"

	"github.com/hambarthegoat/PWEB-FP2-be/internal/dto"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type VoteService interface {
	UpsertPostVote(postID, userID string, value int8) (*dto.VoteResult, error)
	UpsertCommentVote(commentID, userID string, value int8) (*dto.VoteResult, error)
}
type voteService struct{ repo repository.VoteRepository }

func NewVoteService(r repository.VoteRepository) VoteService { return &voteService{repo: r} }

// ON CONFLICT (user_id, target_id, target_type) DO UPDATE
// Recalculate & update Post/Comment.Score

func (s *voteService) UpsertPostVote(postID, userID string, value int8) (*dto.VoteResult, error) {
	return nil, errors.New("not implemented")
}

func (s *voteService) UpsertCommentVote(commentID, userID string, value int8) (*dto.VoteResult, error) {
	return nil, errors.New("not implemented")
}