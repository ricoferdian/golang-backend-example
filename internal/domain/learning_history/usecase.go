package learning_history

import (
	"context"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
)

type LearningHistoryUseCase interface {
	GetUserLearningHistory(ctx context.Context, userID int64) ([]entity2.LearningHistoryEntity, error)
	SubmitLearningHistory(ctx context.Context, historyModel entity2.SubmitLearningHistoryEntity) (*entity2.SubmitLearningHistoryEntity, error)
}
