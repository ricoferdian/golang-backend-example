package redis

import (
	"github.com/redis/go-redis/v9"
	"kora-backend/internal/domain/purchase"
)

type RedisChoreoPurchaseRepository struct {
	redisCli *redis.Client
}

func NewRedisChoreoPurchaseRepository(redisCli *redis.Client) purchase.ChoreoPurchaseCacheRepo {
	return &RedisChoreoPurchaseRepository{
		redisCli: redisCli,
	}
}
