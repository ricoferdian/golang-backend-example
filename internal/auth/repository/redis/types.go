package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

type RedisUserAuthRepository struct {
	redisCli *redis.Client
}

func (r RedisUserAuthRepository) GetSingleUserByUniqueFilter(ctx context.Context, entity entity.UserFilterEntity) (*model.RbacUserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisUserAuthRepository) InsertSingleUser(ctx context.Context, entity entity.UserEntity) (*model.RbacUserModel, error) {
	//TODO implement me
	panic("implement me")
}

func NewRedisUserAuthRepository(redisCli *redis.Client) auth.UserAuthCacheRepo {
	return &RedisUserAuthRepository{
		redisCli: redisCli,
	}
}
