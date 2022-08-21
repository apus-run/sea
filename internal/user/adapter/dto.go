package adapter

import (
	"github.com/apus-run/sea/internal/user/domain"
)

// UserRequest 请求数据结果
type UserRequest struct {
	ID       uint64 `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserResponse 返回数据结构
type UserResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewUser(user *domain.User) *UserRequest {
	return &UserRequest{
		ID:       user.ID,
		Username: user.Username,
	}
}

type RegisterRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

func (register *RegisterRequest) ToEntity() *domain.User {
	return &domain.User{
		Username: "",
	}
}
