package app

import (
	"context"
	"github.com/apus-run/sea/internal/user/data"
	"github.com/apus-run/sea/pkg/log"

	"github.com/pkg/errors"

	"github.com/apus-run/sea/internal/user/domain"
)

var ErrRegistrationUseCaseUserAlreadyCreated = errors.New("registration.userAlreadyCreated")

// UserUseCase . 业务逻辑组装的契约
type UserUseCase interface {
	Register(ctx context.Context, user *domain.User, password string) (*domain.User, error)
	Login(ctx context.Context, username, password string) (*domain.User, error)
}

type userUseCase struct {
	userRepo domain.UserRepository
	log      *log.Helper
}

func NewUserUseCase(userRepository domain.UserRepository, logger log.Logger) *userUseCase {
	return &userUseCase{
		userRepo: userRepository,
		log:      log.NewHelper(logger),
	}
}

// func (uc *UserUseCase) Register(ctx context.Context) entity.UserAggregate {
// 	user := entity.CreateNewUser(req)
// 	user.Validate()
// 	// 过程式
// 	// Validate(user)
// 	repo.SaveUser(user)
// 	event.PublishRegisterEvent(user)
// 	return user
// }

func (uc *userUseCase) Register(ctx context.Context, user *domain.User, password string) (*domain.User, error) {
	users, err := uc.userRepo.GetUserByUsername(ctx, user.Username)
	// 对是否有记录进行判断, 根据业务需求, 可进行更多处理
	if err != nil && errors.Is(err, data.ErrRecordNotFound) {
		// ...
		return nil, err
	}

	if len(users) > 0 {
		return nil, errors.New("此用户已存在")
	}

	pwd, _ := user.HashPassword(password)

	newUser, err := uc.userRepo.Create(ctx, user, pwd)

	if err != nil {
		return nil, errors.New("用户注册失败")
	}

	return newUser, nil
}

func (uc *userUseCase) Login(ctx context.Context, username, password string) (*domain.User, error) {

	return nil, nil
}
