package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	"github.com/apus-run/sea/pkg/database"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepository)

// Data .
type Data struct {
	db  *database.DB
	rdb *redis.Client
}

func NewData(db *database.DB, rdb *redis.Client) (*Data, error) {
	return &Data{
		db:  db,
		rdb: rdb,
	}, nil
}
