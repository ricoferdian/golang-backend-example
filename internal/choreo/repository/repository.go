package repository

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
)

type ChoreoRepositoryImpl struct {
	choreo.ChoreoDatabaseRepo
	choreo.ChoreoCacheRepo
	choreo.S3ChoreoContentRepo
}

func NewChoreoRepository(
	dbRepo choreo.ChoreoDatabaseRepo,
	redisRepo choreo.ChoreoCacheRepo,
	s3Repo choreo.S3ChoreoContentRepo,
) choreo.ChoreoRepository {
	return &ChoreoRepositoryImpl{
		dbRepo,
		redisRepo,
		s3Repo,
	}
}
