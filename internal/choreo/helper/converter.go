package helper

import (
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

func MusicModelToEntity(musicModel model.MusicModel) entity.MusicEntity {
	return entity.MusicEntity{
		MusicID:    musicModel.MusicID,
		ArtistName: musicModel.ArtistName,
		Title:      musicModel.Title,
	}
}
func ChoreographerModelToEntity(musicModel model.ChoreographerModel) entity.ChoreographerEntity {
	return entity.ChoreographerEntity{
		ChoreographerID:   musicModel.ChoreographerID,
		ChoreographerName: musicModel.ChoreographerName,
	}
}

func ChoreoModelToEntity(model model.ChoreographyModel) entity.ChoreographyEntity {
	return entity.ChoreographyEntity{
		ChoreoID:          model.ChoreoID,
		Title:             model.Title.String,
		Description:       model.Description.String,
		Difficulty:        model.Difficulty.Int32,
		Duration:          model.Duration.Float64,
		IsActive:          model.IsActive.Int32,
		VideoPreviewURL:   model.VideoPreviewURL.String,
		VideoThumbnailURL: model.VideoThumbnailURL.String,
		ChoreographerID:   model.ChoreographerID.Int64,
		MusicID:           model.MusicID.Int64,
		Order:             model.Position.Int32,
		ChoreographerData: nil,
		MusicData:         nil,
	}
}

func ChoreoDetailToEntity(model model.ChoreographyDetailModel) entity.ChoreographyDetailEntity {
	return entity.ChoreographyDetailEntity{
		ChoreoDetailID:       model.ChoreoDetailID,
		ChoreoID:             model.ChoreoID.Int64,
		ChoreoData:           nil,
		Title:                model.Title.String,
		Duration:             model.Duration.Float64,
		IsActive:             model.IsActive.Int32,
		VideoURL:             model.VideoURL.String,
		VideoThumbnailURL:    model.VideoThumbnailURL.String,
		Order:                model.Position.Int32,
		VisionAngleThreshold: model.VisionAngleThreshold.Float64,
		VisionTimeOffset:     model.VisionTimeOffset.Float64,
		VisionBodyPose:       model.VisionBodyPose.String,
	}
}
