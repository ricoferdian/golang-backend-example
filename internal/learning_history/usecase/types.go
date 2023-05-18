package usecase

import (
	"kora-backend/internal/domain/common"
	"kora-backend/internal/domain/learning_history"
)

type LearningHistoryUseCaseImpl struct {
	baseRepo common.BaseRepository
}

func NewLearningHistoryUseCase(baseRepo common.BaseRepository) learning_history.LearningHistoryUseCase {
	return &LearningHistoryUseCaseImpl{
		baseRepo: baseRepo,
	}
}
