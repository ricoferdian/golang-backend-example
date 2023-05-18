package learning_history

import (
	"context"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
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
