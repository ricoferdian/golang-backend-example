package learning_history

import (
	"context"
	"kora-backend/internal/entity"
)

type LearningHistoryUseCase interface {
	GetUserLearningHistory(ctx context.Context, userID int64) ([]entity.LearningHistoryEntity, error)
	SubmitLearningHistory(ctx context.Context, historyModel entity.SubmitLearningHistoryEntity) (*entity.SubmitLearningHistoryEntity, error)
}
