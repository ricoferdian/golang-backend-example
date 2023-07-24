package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
)

type RedisChoreoRepository struct {
	redisCli *redis.Client
}

func NewRedisChoreoRepository(redisCli *redis.Client) choreo.ChoreoCacheRepo {
	return &RedisChoreoRepository{
		redisCli: redisCli,
	}
}
