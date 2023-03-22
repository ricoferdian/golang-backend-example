package repository

import (
	"kora-backend/internal/domain/choreo"
)

type ChoreoRepositoryImpl struct {
	choreo.ChoreoDatabaseRepo
	choreo.ChoreoCacheRepo
}

func NewChoreoRepository(
	dbRepo choreo.ChoreoDatabaseRepo,
	redisRepo choreo.ChoreoCacheRepo,
) ChoreoRepositoryImpl {
	return ChoreoRepositoryImpl{
		dbRepo,
		redisRepo,
	}
}
