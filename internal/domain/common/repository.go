package common

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/choreographer"
	"kora-backend/internal/domain/music"
)

type BaseRepository interface {
	UserAuthRepository() auth.UserAuthRepository
	ChoreoRepository() choreo.ChoreoRepository
	MusicRepository() music.MusicRepository
	ChoreographerRepository() choreographer.ChoreographerRepository
	GetAppConfig() *helper.AppConfig
}
