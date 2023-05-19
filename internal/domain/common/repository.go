package common

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/choreographer"
	"kora-backend/internal/domain/learning_history"
	"kora-backend/internal/domain/music"
	"kora-backend/internal/domain/purchase"
)

type BaseRepository interface {
	UserAuthRepository() auth.UserAuthRepository
	ChoreoRepository() choreo.ChoreoRepository
	MusicRepository() music.MusicRepository
	ChoreographerRepository() choreographer.ChoreographerRepository
	LearningHistoryRepository() learning_history.LearningHistoryRepository
	ChoreoPurchaseRepository() purchase.ChoreoPurchaseRepository
	GetAppConfig() *helper.AppConfig
}
