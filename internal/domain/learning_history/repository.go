package learning_history

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

type LearningHistoryDatabaseRepo interface {
	GetLearningHistoryList(ctx context.Context, filter model.LearningHistoryFilter) ([]model.LearningHistoryModel, error)
	InsertLearningHistory(ctx context.Context, history entity.SubmitLearningHistoryEntity) (*model.SubmitLearningHistoryModel, error)
}

type LearningHistoryCacheRepo interface {
}

type LearningHistoryRepository interface {
	LearningHistoryDatabaseRepo
	LearningHistoryCacheRepo
}
