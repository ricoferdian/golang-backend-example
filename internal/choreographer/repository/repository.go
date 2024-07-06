package repository

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreographer"
)

type ChoreographerRepositoryImpl struct {
	choreographer.ChoreographerDatabaseRepo
	choreographer.ChoreographerCacheRepo
	choreographer.S3ChoreographerContentRepo
}

func NewChoreographerRepository(
	dbRepo choreographer.ChoreographerDatabaseRepo,
	redisRepo choreographer.ChoreographerCacheRepo,
	s3Repo choreographer.S3ChoreographerContentRepo,
) choreographer.ChoreographerRepository {
	return &ChoreographerRepositoryImpl{
		dbRepo,
		redisRepo,
		s3Repo,
	}
}
