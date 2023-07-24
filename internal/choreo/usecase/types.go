package usecase

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
)

type ChoreoUseCaseImpl struct {
	baseRepo common.BaseRepository
}

func NewChoreoUseCase(baseRepo common.BaseRepository) choreo.ChoreoUseCase {
	return &ChoreoUseCaseImpl{
		baseRepo: baseRepo,
	}
}
