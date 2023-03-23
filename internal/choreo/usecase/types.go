package usecase

import (
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/common"
)

type ChoreoUseCaseImpl struct {
	baseRepo common.BaseRepository
}

func NewChoreoUseCase(baseRepo common.BaseRepository) choreo.ChoreoUseCase {
	return &ChoreoUseCaseImpl{
		baseRepo: baseRepo,
	}
}
