package jwt

import (
	"github.com/redis/go-redis/v9"
)

type JwtConfig struct {
	private []byte
}

type JwtStorage struct {
	redis  *redis.Client
	config *JwtConfig
}

func NewJwtStorage(redis *redis.Client, config *JwtConfig) (*JwtStorage, error) {
	return &JwtStorage{
		redis:  redis,
		config: config,
	}, nil
}
