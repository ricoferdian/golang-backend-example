package redis

import (
	"github.com/redis/go-redis/v9"
	"kora-backend/internal/domain/authdomain"
)

type RedisUserAuthRepository struct {
	redisCli *redis.Client
}

func NewRedisUserAuthRepository(redisCli *redis.Client) authdomain.UserAuthDatabaseRepo {
	return RedisUserAuthRepository{
		redisCli: redisCli,
	}
}
