package user

import (
	"github.com/gin-gonic/gin"

	"github.com/apus-run/sea/internal/user/adapter"
	"github.com/apus-run/sea/internal/user/app"
	"github.com/apus-run/sea/internal/user/data"
	"github.com/apus-run/sea/pkg/log"
	"github.com/apus-run/sea/pkg/xgin"
)

// Module is a struct that defines all dependencies inside user module
type Module struct {
}

// Configure setups all dependencies
func (m *Module) Configure(logger log.Logger, db *data.Data, engine *gin.Engine) {
	repository := data.NewUserRepository(db, logger)
	useCase := app.NewUserUseCase(repository, logger)
	controller := adapter.NewUserController(useCase, logger)
	engine.POST("/register", controller.Register)
	engine.POST("/login", xgin.Handle(controller.Login))
}
