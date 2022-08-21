package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/apus-run/sea/internal/user/domain"
	"github.com/apus-run/sea/pkg/log"
)

var ErrRecordNotFound = errors.New("record not found")

type userRepository struct {
	data *Data
	log  *log.Helper
}

func NewUserRepository(data *Data, logger log.Logger) domain.UserRepository {
	return &userRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User, password string) (*domain.User, error) {
	row := newUser(user, password)
	result, err := ur.data.db.InsertContext(ctx, row)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("添加用户数据失败: %v", err))
	}

	row.ID = uint64(id)

	t := row.ToEntity()

	return &t, err
}

func (ur *userRepository) GetUserByUsername(ctx context.Context, username string) (domain.Users, error) {
	var rows []user

	err := ur.data.db.Get(&user{}, "select * from users where username = ?", username)

	if err == sql.ErrNoRows {
		return nil, errors.Wrap(ErrRecordNotFound, fmt.Sprintf("此用户不存在, err: %v", err))
	}

	if err != nil {
		return nil, errors.Wrap(err, "查询用户出错了")
	}

	result := make(domain.Users, 0, len(rows))
	for _, row := range rows {
		result = append(result, row.ToEntity())
	}
	return result, nil
}
