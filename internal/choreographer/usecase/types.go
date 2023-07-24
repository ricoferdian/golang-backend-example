package usecase

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreographer"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
)

type ChoreographerUseCaseImpl struct {
	baseRepo common.BaseRepository
}

func NewChoreographerUseCase(baseRepo common.BaseRepository) choreographer.ChoreographerUseCase {
	return &ChoreographerUseCaseImpl{
		baseRepo: baseRepo,
	}
}
