package app

import (
	"context"
	"github.com/apus-run/sea/internal/user/domain"
)

// UserQryExe .
type UserQryExe struct {
	repo domain.UserRepository
}

// NewUserQryExe ,
func NewUserQryExe(repo domain.UserRepository) *UserQryExe {
	return &UserQryExe{repo: repo}
}

// Get .
func (qry *UserQryExe) Get(ctx context.Context, id uint64) (*domain.User, error) {
	return nil, nil
}

// List .
func (qry *UserQryExe) List(ctx context.Context, id uint64) (*domain.Users, error) {
	return nil, nil
}
