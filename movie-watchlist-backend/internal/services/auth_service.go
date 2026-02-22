package services

import (
	"errors"

	"movie-watchlist-backend/internal/config"
	"movie-watchlist-backend/internal/models"
	"movie-watchlist-backend/internal/repositories"
	"movie-watchlist-backend/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
	config   *config.Config
}

func NewAuthService(repo *repositories.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo: repo,
		config:   cfg,
	}
}

func (s *AuthService) RegisterUser(email, password, name string) (*models.User, error) {
	// Check if exists
	if _, err := s.userRepo.FindByEmail(email); err == nil {
		return nil, errors.New("user with this email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashed),
		Name:     name,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) LoginUser(email, password string) (string, *models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, s.config)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
