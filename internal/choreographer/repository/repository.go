package repository

import "kora-backend/internal/domain/choreographer"

type ChoreographerRepositoryImpl struct {
	choreographer.ChoreographerDatabaseRepo
	choreographer.ChoreographerCacheRepo
}

func NewChoreographerRepository(
	dbRepo choreographer.ChoreographerDatabaseRepo,
	redisRepo choreographer.ChoreographerCacheRepo,
) choreographer.ChoreographerRepository {
	return &ChoreographerRepositoryImpl{
		dbRepo,
		redisRepo,
	}
}
