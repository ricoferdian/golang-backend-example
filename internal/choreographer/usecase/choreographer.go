package usecase

import (
	"context"
	"errors"
	"github.com/Kora-Dance/koradance-backend/internal/choreo/helper"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func (c ChoreographerUseCaseImpl) GetChoreographerList(ctx context.Context) ([]entity.ChoreographerEntity, error) {
	choreographerList, err := c.baseRepo.ChoreographerRepository().GetChoreographerList(ctx)
	if err != nil {
		return nil, err
	}
	var result []entity.ChoreographerEntity
	for _, choreographerData := range choreographerList {
		data := helper.ChoreographerModelToEntity(choreographerData)
		result = append(result, data)
	}
	return result, nil
}

func (c ChoreographerUseCaseImpl) GetChoreographerByID(ctx context.Context, filter entity.ChoreographerFilter) (*entity.ChoreographerEntity, error) {
	choreographerData, err := c.baseRepo.ChoreographerRepository().GetChoreographerById(ctx, filter.ChoreographerID)
	if err != nil {
		return nil, err
	}
	if choreographerData == nil {
		return nil, errors.New("nil choreographer data")
	}
	result := helper.ChoreographerModelToEntity(*choreographerData)
	return &result, nil
}
