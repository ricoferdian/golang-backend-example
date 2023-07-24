package usecase

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/internal/domain/learning_history"
)

type LearningHistoryUseCaseImpl struct {
	baseRepo common.BaseRepository
}

func NewLearningHistoryUseCase(baseRepo common.BaseRepository) learning_history.LearningHistoryUseCase {
	return &LearningHistoryUseCaseImpl{
		baseRepo: baseRepo,
	}
}
