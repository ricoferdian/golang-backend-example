package common

import (
	"kora-backend/internal/domain/authdomain"
	"kora-backend/internal/domain/choreo"
)

type BaseRepository interface {
	UserAuthRepository() authdomain.UserAuthRepository
	ChoreoRepository() choreo.ChoreoRepository
}
