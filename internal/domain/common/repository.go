package common

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/auth"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreographer"
	"github.com/Kora-Dance/koradance-backend/internal/domain/learning_history"
	"github.com/Kora-Dance/koradance-backend/internal/domain/like_save"
	"github.com/Kora-Dance/koradance-backend/internal/domain/music"
	"github.com/Kora-Dance/koradance-backend/internal/domain/purchase"
)

type BaseRepository interface {
	UserAuthRepository() auth.UserAuthRepository
	ChoreoRepository() choreo.ChoreoRepository
	MusicRepository() music.MusicRepository
	ChoreographerRepository() choreographer.ChoreographerRepository
	LearningHistoryRepository() learning_history.LearningHistoryRepository
	ChoreoPurchaseRepository() purchase.ChoreoPurchaseRepository
	GetAppConfig() *helper.AppConfig
	ChoreoLikeSaveRepository() like_save.LikeSaveRepository
}
