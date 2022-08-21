package domain

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Users Entity
type Users []User

// User Entity
type User struct {
	ID       uint64
	Username string
}

// Address 值对象
type Address struct{}

// UserAggregate 聚合根我们一般用组合的形式
type UserAggregate struct {
	User
	Address []Address
}

// Validate 参数校验
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("用户信息有错误")
	}
	return nil
}

// HashPassword 密码加密
func (u *User) HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	pass := string(b)
	return pass, nil
}

// VerifyPassword 验证密码
func (u *User) VerifyPassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

// UserRepository 与持久层通信的契约
type UserRepository interface {
	Create(ctx context.Context, user *User, password string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (Users, error)
}
