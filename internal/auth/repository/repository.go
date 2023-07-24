package repository

import "github.com/Kora-Dance/koradance-backend/internal/domain/auth"

type UserAuthRepositoryImpl struct {
	auth.UserAuthDatabaseRepo
	auth.UserAuthCacheRepo
}

func NewUserAuthRepository(
	dbRepo auth.UserAuthDatabaseRepo,
	redisRepo auth.UserAuthCacheRepo,
) auth.UserAuthRepository {
	return &UserAuthRepositoryImpl{
		dbRepo,
		redisRepo,
	}
}
