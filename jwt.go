package jwt

import (
	"github.com/redis/go-redis/v9"
)

type JwtConfig struct {
	private []byte
}

func NewJwtConfig(privateString string) (*JwtConfig, error) {
	return &JwtConfig{
		private: []byte(privateString),
	}, nil
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
