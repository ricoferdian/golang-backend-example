package usecase

import (
	"context"
	"kora-backend/internal/entity"
)

func (c ChoreoUseCaseImpl) GetChoreoList(ctx context.Context) (choreoResult []entity.ChoreographyEntity, err error) {
	choreoList, err := c.baseRepo.ChoreoRepository().GetChoreoList(ctx)
	if err != nil {
		return choreoResult, err
	}
	for _, choreoData := range choreoList {
		choreo := entity.ChoreographyEntity{
			ChoreoID:          choreoData.ChoreoID,
			Title:             choreoData.Title.String,
			Description:       choreoData.Description.String,
			Difficulty:        choreoData.Difficulty.Int32,
			Duration:          choreoData.Duration.Float64,
			IsActive:          choreoData.IsActive.Int32,
			VideoPreviewURL:   choreoData.VideoPreviewURL.String,
			VideoThumbnailURL: choreoData.VideoThumbnailURL.String,
			ChoreographerID:   choreoData.ChoreographerID.Int64,
			MusicID:           choreoData.MusicID.Int64,
			Order:             choreoData.Position.Int32,
			ChoreographerData: nil,
			MusicData:         nil,
		}
		choreoResult = append(choreoResult, choreo)
	}
	return choreoResult, nil
}
