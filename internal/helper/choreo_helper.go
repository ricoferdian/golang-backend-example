package helper

import (
	"database/sql"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func ChoreographerEntityToModel(choreographer entity2.ChoreographerEntity) model.ChoreographerModel {
	return model.ChoreographerModel{
		ChoreographerID:   choreographer.ChoreographerID,
		ChoreographerName: choreographer.ChoreographerName,
		Description: sql.NullString{
			String: choreographer.Description,
			Valid:  true,
		},
		ProfileImageURL: sql.NullString{
			String: choreographer.ProfileImageURL,
			Valid:  true,
		},
	}
}
func ChoreographerModelToEntity(choreographer model.ChoreographerModel) entity2.ChoreographerEntity {
	return entity2.ChoreographerEntity{
		ChoreographerID:   choreographer.ChoreographerID,
		ChoreographerName: choreographer.ChoreographerName,
		Description:       choreographer.Description.String,
		ProfileImageURL:   choreographer.ProfileImageURL.String,
	}
}

func ChoreoModelToEntity(model model.ChoreographyModel) entity2.ChoreographyEntity {
	return entity2.ChoreographyEntity{
		ChoreoID:             model.ChoreoID,
		Title:                model.Title.String,
		Description:          model.Description.String,
		Difficulty:           model.Difficulty.Int32,
		Duration:             model.Duration.Float64,
		IsActive:             model.IsActive.Int32,
		VideoPreviewURL:      model.VideoPreviewURL.String,
		VideoThumbnailURL:    model.VideoThumbnailURL.String,
		CDNVideoPreviewURL:   model.CDNVideoPreviewURL.String,
		CDNVideoThumbnailURL: model.CDNVideoThumbnailURL.String,
		ChoreographerID:      model.ChoreographerID.Int64,
		MusicID:              model.MusicID.Int64,
		AdditionalInfo:       model.AdditionalInfo.String,
		Order:                model.Position.Int32,
		ChoreographerData:    nil,
		MusicData:            nil,
	}
}

func ChoreoDetailToEntity(model model.ChoreographyDetailModel) entity2.ChoreographyDetailEntity {
	return entity2.ChoreographyDetailEntity{
		ChoreoDetailID:       model.ChoreoDetailID,
		ChoreoID:             model.ChoreoID.Int64,
		ChoreoData:           nil,
		Title:                model.Title.String,
		Duration:             model.Duration.Float64,
		IsActive:             model.IsActive.Int32,
		VideoURL:             model.VideoURL.String,
		VideoThumbnailURL:    model.VideoThumbnailURL.String,
		CDNVideoURL:          model.CDNVideoURL.String,
		CDNVideoThumbnailURL: model.CDNVideoThumbnailURL.String,
		TestVideoURL:         model.TestVideoURL.String,
		CDNTestVideoURL:      model.CDNTestVideoURL.String,
		Order:                model.Position.Int32,
		VisionAngleThreshold: model.VisionAngleThreshold.Float64,
		VisionTimeOffset:     model.VisionTimeOffset.Float64,
		VisionBodyPose:       model.VisionBodyPose.String,
	}
}

func ChoreoEntityToModel(entity entity2.ChoreographyEntity) model.ChoreographyModel {
	return model.ChoreographyModel{
		ChoreoID: entity.ChoreoID,
		Title: sql.NullString{
			String: entity.Title,
			Valid:  true,
		},
		Description: sql.NullString{
			String: entity.Description,
			Valid:  true,
		},
		Difficulty: sql.NullInt32{
			Int32: entity.Difficulty,
			Valid: true,
		},
		Duration: sql.NullFloat64{
			Float64: entity.Duration,
			Valid:   true,
		},
		IsActive: sql.NullInt32{
			Int32: entity.IsActive,
			Valid: true,
		},
		VideoPreviewURL: sql.NullString{
			String: entity.VideoPreviewURL,
			Valid:  true,
		},
		VideoThumbnailURL: sql.NullString{
			String: entity.VideoThumbnailURL,
			Valid:  true,
		},
		CDNVideoPreviewURL: sql.NullString{
			String: entity.CDNVideoPreviewURL,
			Valid:  true,
		},
		CDNVideoThumbnailURL: sql.NullString{
			String: entity.CDNVideoThumbnailURL,
			Valid:  true,
		},
		ChoreographerID: sql.NullInt64{
			Int64: entity.ChoreographerID,
			Valid: true,
		},
		MusicID: sql.NullInt64{
			Int64: entity.MusicID,
			Valid: true,
		},
		AdditionalInfo: sql.NullString{
			String: entity.AdditionalInfo,
			Valid:  true,
		},
		Position: sql.NullInt32{
			Int32: entity.Order,
			Valid: true,
		},
	}
}

func ChoreoDetailEntityToModel(entity entity2.ChoreographyDetailEntity) model.ChoreographyDetailModel {
	return model.ChoreographyDetailModel{
		ChoreoDetailID: entity.ChoreoDetailID,
		ChoreoID: sql.NullInt64{
			Int64: entity.ChoreoID,
			Valid: true,
		},
		Title: sql.NullString{
			String: entity.Title,
			Valid:  true,
		},
		Duration: sql.NullFloat64{
			Float64: entity.Duration,
			Valid:   true,
		},
		IsActive: sql.NullInt32{
			Int32: entity.IsActive,
			Valid: true,
		},
		VideoURL: sql.NullString{
			String: entity.VideoURL,
			Valid:  true,
		},
		VideoThumbnailURL: sql.NullString{
			String: entity.VideoThumbnailURL,
			Valid:  true,
		},
		CDNVideoURL: sql.NullString{
			String: entity.CDNVideoURL,
			Valid:  true,
		},
		CDNVideoThumbnailURL: sql.NullString{
			String: entity.CDNVideoThumbnailURL,
			Valid:  true,
		},
		TestVideoURL: sql.NullString{
			String: entity.TestVideoURL,
			Valid:  true,
		},
		CDNTestVideoURL: sql.NullString{
			String: entity.CDNTestVideoURL,
			Valid:  true,
		},
		Position: sql.NullInt32{
			Int32: entity.Order,
			Valid: true,
		},
		VisionAngleThreshold: sql.NullFloat64{
			Float64: entity.VisionAngleThreshold,
			Valid:   true,
		},
		VisionTimeOffset: sql.NullFloat64{
			Float64: entity.VisionTimeOffset,
			Valid:   true,
		},
		VisionBodyPose: sql.NullString{
			String: entity.VisionBodyPose,
			Valid:  true,
		},
	}
}
