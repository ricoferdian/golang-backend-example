package repository

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/choreographer"
	"kora-backend/internal/domain/common"
	"kora-backend/internal/domain/learning_history"
	"kora-backend/internal/domain/music"
)

type Repository struct {
	userAuthRepo        auth.UserAuthRepository
	choreoRepo          choreo.ChoreoRepository
	musicRepo           music.MusicRepository
	choreographerRepo   choreographer.ChoreographerRepository
	learningHistoryRepo learning_history.LearningHistoryRepository
}

func NewRepository(
	userAuthRepo auth.UserAuthRepository,
	choreoRepo choreo.ChoreoRepository,
	musicRepo music.MusicRepository,
	choreographerRepo choreographer.ChoreographerRepository,
	learningHistoryRepo learning_history.LearningHistoryRepository,
) Repository {
	return Repository{
		userAuthRepo:        userAuthRepo,
		choreoRepo:          choreoRepo,
		musicRepo:           musicRepo,
		choreographerRepo:   choreographerRepo,
		learningHistoryRepo: learningHistoryRepo,
	}
}

type BaseRepositoryImpl struct {
	repo   Repository
	config *helper.AppConfig
}

func (repo BaseRepositoryImpl) GetAppConfig() *helper.AppConfig {
	return repo.config
}

func (repo BaseRepositoryImpl) ChoreographerRepository() choreographer.ChoreographerRepository {
	return repo.repo.choreographerRepo
}

func (repo BaseRepositoryImpl) MusicRepository() music.MusicRepository {
	return repo.repo.musicRepo
}

func (repo BaseRepositoryImpl) ChoreoRepository() choreo.ChoreoRepository {
	return repo.repo.choreoRepo
}

func (repo BaseRepositoryImpl) UserAuthRepository() auth.UserAuthRepository {
	return repo.repo.userAuthRepo
}

func (repo BaseRepositoryImpl) LearningHistoryRepository() learning_history.LearningHistoryRepository {
	return repo.repo.learningHistoryRepo
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
