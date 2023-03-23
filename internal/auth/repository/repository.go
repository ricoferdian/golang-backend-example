package repository

import "kora-backend/internal/domain/auth"

type UserAuthRepositoryImpl struct {
	auth.UserAuthDatabaseRepo
	auth.UserAuthCacheRepo
}

func NewUserAuthRepository(
	dbRepo auth.UserAuthDatabaseRepo,
	redisRepo auth.UserAuthCacheRepo,
) auth.UserAuthRepository {
	return UserAuthRepositoryImpl{
		dbRepo,
		redisRepo,
	}
}
