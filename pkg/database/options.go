package database

import (
	"time"

	"github.com/apus-run/sea/pkg/log"
)

// Option is database option
type Option func(*options)

type options struct {
	dsn             string        // 数据库连接地址
	maxOpenConns    int           // default: 100
	maxIdleConns    int           // default: 10
	connMaxLifetime time.Duration // default: 1h
	logging         bool          // default: "false"

	logger log.Logger
}

// DefaultOptions .
func DefaultOptions() *options {
	return &options{
		dsn:             "",
		maxOpenConns:    100,
		maxIdleConns:    10,
		connMaxLifetime: 10 * time.Minute,
		logging:         false,

		logger: log.DefaultLogger,
	}
}

// DSN .
func DSN(dsn string) Option {
	return func(o *options) {
		o.dsn = dsn
	}
}

// MaxOpenConns .
func MaxOpenConns(moc int) Option {
	return func(o *options) {
		o.maxOpenConns = moc
	}
}

// MaxIdleConns .
func MaxIdleConns(mic int) Option {
	return func(o *options) {
		o.maxIdleConns = mic
	}
}

// ConnMaxLifetime .
func ConnMaxLifetime(cml time.Duration) Option {
	return func(o *options) {
		o.connMaxLifetime = cml
	}
}

// Logging .
func Logging(logging bool) Option {
	return func(o *options) {
		o.logging = logging
	}
}

// WithLogger .
func WithLogger(l log.Logger) Option {
	return func(o *options) {
		o.logger = l
	}
}
