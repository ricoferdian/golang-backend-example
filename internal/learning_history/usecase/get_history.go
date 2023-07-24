package usecase

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/learning_history/helper"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func (c LearningHistoryUseCaseImpl) GetUserLearningHistory(ctx context.Context, userID int64) (resultEntity []entity.LearningHistoryEntity, err error) {
	filter := model.LearningHistoryFilter{
		UserID: userID,
	}
	historyList, err := c.baseRepo.LearningHistoryRepository().GetLearningHistoryList(ctx, filter)
	if err != nil {
		return nil, err
	}
	for _, historyData := range historyList {
		resultEntity = append(resultEntity, helper.LearningHistoryModelToEntity(historyData))
	}
	return resultEntity, nil
}
