package cache

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Redis is an object to hold the lifetime reusable pool object
type Redis struct {
	pool *redis.Pool
}

// Config holds configuration values
type Config struct {
	Address         string
	MaxActive       int
	MaxIdle         int
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration
}

// InitializeRedis initializes redis cache
func InitializeRedis(cfg Config) *Redis {
	return &Redis{pool: &redis.Pool{
		MaxIdle:         cfg.MaxIdle,
		MaxActive:       cfg.MaxActive,
		IdleTimeout:     cfg.IdleTimeout,
		MaxConnLifetime: cfg.MaxConnLifetime,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Address)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}}
}

// Do does operation to redis with given command & arguments
func (r *Redis) Do(ctx context.Context, command string, args ...interface{}) (interface{}, error) {
	conn, err := r.pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close() // close connection after we use it

	return conn.Do(command, args...)
}

// Close closes redis pool
func (r *Redis) Close() error {
	// cleaning up pool when it's called explicitly
	return r.pool.Close()
}
