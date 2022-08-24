package data

import (
	"github.com/go-redis/redis/v8"

	"github.com/apus-run/sea/pkg/database"
)

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
