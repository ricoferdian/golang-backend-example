package redis

import (
	"github.com/redis/go-redis/v9"
	"kora-backend/internal/domain/learning_history"
)

type RedisLearningHistoryRepository struct {
	redisCli *redis.Client
}

func NewRedisLearningHistoryRepository(redisCli *redis.Client) learning_history.LearningHistoryCacheRepo {
	return &RedisLearningHistoryRepository{
		redisCli: redisCli,
	}
}
