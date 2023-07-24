package redis

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/purchase"
	"github.com/redis/go-redis/v9"
)

type RedisChoreoPurchaseRepository struct {
	redisCli *redis.Client
}

func NewRedisChoreoPurchaseRepository(redisCli *redis.Client) purchase.ChoreoPurchaseCacheRepo {
	return &RedisChoreoPurchaseRepository{
		redisCli: redisCli,
	}
}
