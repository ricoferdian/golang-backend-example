package redis

import (
	"github.com/redis/go-redis/v9"
	"kora-backend/internal/domain/choreo"
)

type RedisChoreoRepository struct {
	redisCli *redis.Client
}

func NewRedisChoreoRepository(redisCli *redis.Client) choreo.ChoreoCacheRepo {
	return &RedisChoreoRepository{
		redisCli: redisCli,
	}
}
