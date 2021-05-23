package cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type Config struct {
	MaxActive       int
	MaxIdle         int
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration
}

func Init(cfg Config) redis.Conn {
	redispool := redis.Pool{
		MaxIdle:         cfg.MaxIdle,
		MaxActive:       cfg.MaxActive,
		IdleTimeout:     cfg.IdleTimeout,
		MaxConnLifetime: cfg.MaxConnLifetime,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	return redispool.Get()
}
