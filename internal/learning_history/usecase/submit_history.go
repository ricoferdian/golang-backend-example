package usecase

import (
	"context"
	"kora-backend/internal/entity"
	"kora-backend/internal/learning_history/helper"
)

func (c LearningHistoryUseCaseImpl) SubmitLearningHistory(ctx context.Context, submitEntity entity.SubmitLearningHistoryEntity) (*entity.SubmitLearningHistoryEntity, error) {
	submitResult, err := c.baseRepo.LearningHistoryRepository().InsertLearningHistory(ctx, submitEntity)
	if err != nil {
		return nil, err
	}
	submitEntityResult := helper.SubmitLearningHistoryModelToEntity(*submitResult)
	return &submitEntityResult, nil
}
