package usecase

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/learning_history/helper"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func (c LearningHistoryUseCaseImpl) SubmitLearningHistory(ctx context.Context, submitEntity entity.SubmitLearningHistoryEntity) (*entity.SubmitLearningHistoryEntity, error) {
	submitResult, err := c.baseRepo.LearningHistoryRepository().InsertLearningHistory(ctx, submitEntity)
	if err != nil {
		return nil, err
	}
	submitEntityResult := helper.SubmitLearningHistoryModelToEntity(*submitResult)
	return &submitEntityResult, nil
}
