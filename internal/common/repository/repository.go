package repository

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/common"
)

type Repository struct {
	userAuthRepo auth.UserAuthRepository
	choreoRepo   choreo.ChoreoRepository
}

func NewRepository(
	userAuthRepo auth.UserAuthRepository,
	choreoRepo choreo.ChoreoRepository,
) Repository {
	return Repository{
		userAuthRepo: userAuthRepo,
		choreoRepo:   choreoRepo,
	}
}

type BaseRepositoryImpl struct {
	repo   Repository
	config *helper.AppConfig
}

func (repo BaseRepositoryImpl) GetAppConfig() *helper.AppConfig {
	return repo.config
}

func (repo BaseRepositoryImpl) ChoreoRepository() choreo.ChoreoRepository {
	return repo.repo.choreoRepo
}

func (repo BaseRepositoryImpl) UserAuthRepository() auth.UserAuthRepository {
	return repo.repo.userAuthRepo
}

func NewBaseRepository(
	repo Repository,
	config *helper.AppConfig,
) common.BaseRepository {
	return &BaseRepositoryImpl{
		repo:   repo,
		config: config,
	}
}
