package application

import (
	"context"

	"github.com/apus-run/sea/internal/user/domain"
)

// UserCmdExe .
type UserCmdExe struct {
	repo domain.UserRepository
}

// NewUserCmdExe .
func NewUserCmdExe(repo domain.UserRepository) *UserCmdExe {
	return &UserCmdExe{repo: repo}
}

// Add .
func (cmd *UserCmdExe) Add(ctx context.Context, user *domain.User) error {
	return nil
}

// Delete .
func (cmd *UserCmdExe) Delete(ctx context.Context, id uint64) error {
	return nil
}

// Update .
func (cmd *UserCmdExe) Update(ctx context.Context, user *domain.User) error {
	return nil
}
