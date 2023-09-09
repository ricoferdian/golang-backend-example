package usecase

import (
	"context"
	"errors"
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func (c ChoreographerUseCaseImpl) UpsertChoreographer(ctx context.Context, choreographerData entity.ChoreographerEntity) (entity.ChoreographerEntity, error) {
	choreographerModel := helper.ChoreographerEntityToModel(choreographerData)
	choreographerResult, err := c.baseRepo.ChoreographerRepository().UpsertChoreographerByIds(ctx, choreographerModel)
	if err != nil || choreographerResult == nil {
		return choreographerData, err
	}
	choreographerData = helper.ChoreographerModelToEntity(*choreographerResult)
	return choreographerData, nil
}

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

func (m ChoreographerUseCaseImpl) DeleteChoreographerByID(ctx context.Context, choreographerID int64) error {
	err := m.baseRepo.ChoreographerRepository().DeleteChoreographerByID(ctx, choreographerID)
	if err != nil {
		return err
	}
	return nil
}
