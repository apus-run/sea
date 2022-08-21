package main

import (
	"context"
	"os"

	"github.com/apus-run/sea/internal/user"
	"github.com/apus-run/sea/internal/user/data"
	"github.com/apus-run/sea/pkg/database"
	"github.com/apus-run/sea/pkg/log"
	"github.com/gin-gonic/gin"
)

var (
	// LOG .
	LOG = log.NewHelper(log.With(log.GetLogger(), "source", "main"))
)

func main() {
	LOG.Infof("服务启动: [%s]", "")
	logger := log.NewStdLogger(os.Stdout)
	logger = log.With(logger,
		"service.name", "hellworld",
		"service.version", "v1.0.0",
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	engine := gin.Default()

	db := database.NewDatabase(database.DSN(""))
	ctx := context.Background()
	instance := db.Get(ctx, "")

	store, err := data.NewData(instance, nil)
	if err != nil {
		panic("failed to connect database")
	}

	// creates new user module
	userModule := &user.Module{}
	userModule.Configure(logger, store, engine)

}
