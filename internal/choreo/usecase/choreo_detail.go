package usecase

import (
	"context"
	"kora-backend/internal/choreo/helper"
	"kora-backend/internal/entity"
)

func (c ChoreoUseCaseImpl) GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) (choreoResult []entity.ChoreographyDetailEntity, err error) {
	choreoList, err := c.baseRepo.ChoreoRepository().GetChoreoDetailByChoreoID(ctx, filter)
	if err != nil {
		return choreoResult, err
	}
	for _, choreoData := range choreoList {
		choreoResult = append(choreoResult, helper.ChoreoDetailToEntity(choreoData))
	}
	return choreoResult, nil
}
