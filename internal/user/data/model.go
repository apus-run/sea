package data

import (
	"encoding/json"

	"github.com/apus-run/sea/internal/user/domain"
)

type user struct {
	ID       uint64 `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func newUser(u *domain.User, password string) *user {
	return &user{
		Username: u.Username,
		Password: password,
	}
}

func (u *user) TableName() string { return "users" }
func (u *user) KeyName() string   { return "id" }

func (u *user) ToEntity() domain.User {
	return domain.User{
		ID:       u.ID,
		Username: u.Username,
	}
}

// MarshalBinary 实现BinaryMarshaler 接口
func (u *user) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary 实现 BinaryUnMarshaler 接口
func (u *user) UnmarshalBinary(bt []byte) error {
	return json.Unmarshal(bt, u)
}
