package repository

import (
	"kora-backend/internal/domain/learning_history"
)

type LearningHistoryRepositoryImpl struct {
	learning_history.LearningHistoryDatabaseRepo
	learning_history.LearningHistoryCacheRepo
}

func NewLearningHistoryRepository(
	dbRepo learning_history.LearningHistoryDatabaseRepo,
	redisRepo learning_history.LearningHistoryCacheRepo,
) learning_history.LearningHistoryRepository {
	return &LearningHistoryRepositoryImpl{
		dbRepo,
		redisRepo,
	}
}
