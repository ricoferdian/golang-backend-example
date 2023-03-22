package usecase

import (
	"context"
	"kora-backend/internal/entity"
)

func (c ChoreoUseCaseImpl) GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) (choreoResult []entity.ChoreographyDetailEntity, err error) {
	choreoList, err := c.baseRepo.ChoreoRepository().GetChoreoDetailByChoreoID(ctx, filter)
	if err != nil {
		return choreoResult, err
	}
	for _, choreoData := range choreoList {
		choreo := entity.ChoreographyDetailEntity{
			ChoreoDetailID:       choreoData.ChoreoDetailID,
			ChoreoID:             choreoData.ChoreoID.Int64,
			ChoreoData:           nil,
			Title:                choreoData.Title.String,
			Duration:             choreoData.Duration.Float64,
			IsActive:             choreoData.IsActive.Int32,
			VideoURL:             choreoData.VideoURL.String,
			VideoThumbnailURL:    choreoData.VideoThumbnailURL.String,
			Order:                choreoData.Position.Int32,
			VisionAngleThreshold: choreoData.VisionAngleThreshold.Float64,
			VisionTimeOffset:     choreoData.VisionTimeOffset.Float64,
			VisionBodyPose:       choreoData.VisionBodyPose.String,
		}
		choreoResult = append(choreoResult, choreo)
	}
	return choreoResult, nil
}
