package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/hambarthegoat/PWEB-FP2-be/config"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/dto"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/model"
	"github.com/hambarthegoat/PWEB-FP2-be/internal/repository"
)

type AuthService interface {
	Register(email, password string) (*dto.UserResponse, error)
	Login(email, password string) (*dto.LoginResponse, error)
	OAuthStart(provider string) (string, error)
	OAuthCallback(provider, code string) (*dto.LoginResponse, error)
	Logout(userID string) error
}

type authService struct {
	cfg      *config.Config
	userRepo repository.UserRepository
}

func NewAuthService(cfg *config.Config, userRepo repository.UserRepository) AuthService {
	return &authService{cfg: cfg, userRepo: userRepo}
}

func (s *authService) Register(email, password string) (*dto.UserResponse, error) {
	// check existing
	if u, _ := s.userRepo.FindByEmail(email); u != nil {
		return nil, errors.New("email already registered")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil { return nil, err }
	hs := string(hashed)
	user := &model.User{Email: email, Password: &hs}
	if err := s.userRepo.Create(user); err != nil { return nil, err }
	return &dto.UserResponse{ID: user.ID, Email: user.Email}, nil
}

func (s *authService) Login(email, password string) (*dto.LoginResponse, error) {
	u, err := s.userRepo.FindByEmail(email)
	if err != nil { return nil, err }
	if u == nil || u.Password == nil { return nil, errors.New("invalid credentials") }
	if err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	// sign JWT
	claims := jwt.RegisteredClaims{
		Subject:   u.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil { return nil, err }
	return &dto.LoginResponse{Token: signed, UserID: u.ID}, nil
}

func (s *authService) OAuthStart(provider string) (string, error) { return "", errors.New("oauth disabled") }
func (s *authService) OAuthCallback(provider, code string) (*dto.LoginResponse, error) { return nil, errors.New("oauth disabled") }
func (s *authService) Logout(userID string) error { return nil }