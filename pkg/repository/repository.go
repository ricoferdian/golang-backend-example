package repository

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/auth"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreographer"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/internal/domain/learning_history"
	"github.com/Kora-Dance/koradance-backend/internal/domain/like_save"
	"github.com/Kora-Dance/koradance-backend/internal/domain/music"
	"github.com/Kora-Dance/koradance-backend/internal/domain/purchase"
)

type Repository struct {
	userAuthRepo        auth.UserAuthRepository
	choreoRepo          choreo.ChoreoRepository
	musicRepo           music.MusicRepository
	choreographerRepo   choreographer.ChoreographerRepository
	learningHistoryRepo learning_history.LearningHistoryRepository
	choreoPurchaseRepo  purchase.ChoreoPurchaseRepository
	likeSaveRepo        like_save.LikeSaveRepository
}

func NewRepository(
	userAuthRepo auth.UserAuthRepository,
	choreoRepo choreo.ChoreoRepository,
	musicRepo music.MusicRepository,
	choreographerRepo choreographer.ChoreographerRepository,
	learningHistoryRepo learning_history.LearningHistoryRepository,
	choreoPurchaseRepo purchase.ChoreoPurchaseRepository,
	likeSaveRepo like_save.LikeSaveRepository,
) Repository {
	return Repository{
		userAuthRepo:        userAuthRepo,
		choreoRepo:          choreoRepo,
		musicRepo:           musicRepo,
		choreographerRepo:   choreographerRepo,
		learningHistoryRepo: learningHistoryRepo,
		choreoPurchaseRepo:  choreoPurchaseRepo,
		likeSaveRepo:        likeSaveRepo,
	}
}

type BaseRepositoryImpl struct {
	repo   Repository
	config *helper.AppConfig
}

func (repo BaseRepositoryImpl) ChoreoLikeSaveRepository() like_save.LikeSaveRepository {
	return repo.repo.likeSaveRepo
}

func (repo BaseRepositoryImpl) ChoreoPurchaseRepository() purchase.ChoreoPurchaseRepository {
	return repo.repo.choreoPurchaseRepo
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
