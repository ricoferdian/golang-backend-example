package repository

import "kora-backend/internal/domain/authdomain"

type UserAuthRepositoryImpl struct {
	dbRepo    authdomain.UserAuthDatabaseRepo
	redisRepo authdomain.UserAuthCacheRepo
}

func NewUserAuthRepository(
	dbRepo authdomain.UserAuthDatabaseRepo,
	redisRepo authdomain.UserAuthCacheRepo,
) authdomain.UserAuthRepository {
	return UserAuthRepositoryImpl{
		dbRepo:    dbRepo,
		redisRepo: redisRepo,
	}
}
