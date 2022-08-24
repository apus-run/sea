package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/apus-run/sea/internal/user/application"
	"github.com/apus-run/sea/pkg/log"
	"github.com/apus-run/sea/pkg/xgin"
)

type UserController struct {
	user application.UserUseCase

	log *log.Helper
}

func NewUserController(user application.UserUseCase, logger log.Logger) *UserController {
	return &UserController{
		user: user,
		log:  log.NewHelper(logger),
	}
}

func (c *UserController) Login(ctx *xgin.Context) {
	var loginData UserRequest
	err := ctx.Bind(&loginData)
	if err != nil {
		ctx.JSONE(1, err.Error(), nil)
	}
	if len(loginData.Username) < 2 || len(loginData.Password) > 20 {
		ctx.JSONE(1, "username length should between 2 ~ 20", "")
		return
	}

	log.Infof("用户: %v, %v", loginData.Password, loginData.Username)
	// 数据库操作

	ctx.JSONOK("")

	return
}

func (c *UserController) Register(ctx *gin.Context) {
	var registerData RegisterRequest
	if err := ctx.ShouldBindJSON(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := c.user.Register(ctx, registerData.ToEntity(), registerData.Password)

	// 对是否有记录进行判断, 根据业务需求, 可进行更多处理
	if err != nil {
		if errors.Is(err, application.ErrRegistrationUseCaseUserAlreadyCreated) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	ctx.JSON(http.StatusOK, NewUser(result))
}
