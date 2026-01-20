package service

import (
	"context"
	"time"

	"github.com/leirbagxis/example-bot/internal/database/models"
	"github.com/leirbagxis/example-bot/internal/database/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpSertUser(ctx context.Context, user *models.User) error {
	existing, err := s.repo.FindById(ctx, user.UserId)

	if err != nil && err.Error() != "record not found" {
		return err
	}

	now := time.Now()

	if existing == nil {
		user.CreatedAt = now
		user.UpdatedAt = now
		return s.repo.Create(ctx, user)
	}

	existing.FirstName = user.FirstName
	existing.UpdatedAt = now

	return s.repo.Update(ctx, existing)
}

func (s *UserService) GetUserById(ctx context.Context, userID int64) (*models.User, error) {
	return s.repo.FindById(ctx, userID)
}
