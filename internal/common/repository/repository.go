package repository

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/authdomain"
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/common"
)

type Repository struct {
	userAuthRepo authdomain.UserAuthRepository
	choreoRepo   choreo.ChoreoRepository
}

func NewRepository(
	userAuthRepo authdomain.UserAuthRepository,
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

func (repo BaseRepositoryImpl) ChoreoRepository() choreo.ChoreoRepository {
	return repo.repo.choreoRepo
}

func (repo BaseRepositoryImpl) UserAuthRepository() authdomain.UserAuthRepository {
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