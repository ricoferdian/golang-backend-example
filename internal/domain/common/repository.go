package common

import (
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/choreo"
)

type BaseRepository interface {
	UserAuthRepository() auth.UserAuthRepository
	ChoreoRepository() choreo.ChoreoRepository
}
