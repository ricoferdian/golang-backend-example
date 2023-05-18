package usecase

import (
	"context"
	"kora-backend/internal/entity"
	"kora-backend/internal/learning_history/helper"
	"kora-backend/internal/model"
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
