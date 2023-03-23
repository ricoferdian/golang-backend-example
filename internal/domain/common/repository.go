package common

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/choreo"
)

type BaseRepository interface {
	UserAuthRepository() auth.UserAuthRepository
	ChoreoRepository() choreo.ChoreoRepository
	GetAppConfig() *helper.AppConfig
}
