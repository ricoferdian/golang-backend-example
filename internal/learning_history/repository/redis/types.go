package redis

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/learning_history"
	"github.com/redis/go-redis/v9"
)

type RedisLearningHistoryRepository struct {
	redisCli *redis.Client
}

func NewRedisLearningHistoryRepository(redisCli *redis.Client) learning_history.LearningHistoryCacheRepo {
	return &RedisLearningHistoryRepository{
		redisCli: redisCli,
	}
}
