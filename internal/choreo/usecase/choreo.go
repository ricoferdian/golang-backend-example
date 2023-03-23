package usecase

import (
	"context"
	"kora-backend/internal/choreo/helper"
	"kora-backend/internal/entity"
)

func (c ChoreoUseCaseImpl) GetChoreoList(ctx context.Context) (choreoResult []entity.ChoreographyEntity, err error) {
	choreoList, err := c.baseRepo.ChoreoRepository().GetChoreoList(ctx)
	if err != nil {
		return choreoResult, err
	}
	for _, choreoData := range choreoList {
		choreoResult = append(choreoResult, helper.ChoreoModelToEntity(choreoData))
	}
	return choreoResult, nil
}
