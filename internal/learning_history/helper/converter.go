package helper

import (
	"database/sql"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

func LearningHistoryModelToEntity(model model.LearningHistoryModel) entity.LearningHistoryEntity {
	// Entity will never need password when converted from model
	// as the model comes from db, we shall not put password from db to entity
	return entity.LearningHistoryEntity{
		LearningHistoryID: model.LearningHistoryID,
		ChoreoDetailID:    model.ChoreoDetailID,
		UserID:            model.UserID,
		ChoreoDetail:      model.ChoreoDetail.String,
		Device:            model.Device.String,
		Downloaded:        model.Downloaded.Int32,
		Expired:           model.Expired.Int32,
		Progress:          model.Progress.Float64,
		RecordUrl:         model.RecordUrl.String,
		ThumbnailUrl:      model.ThumbnailUrl.String,
	}
}

func SubmitLearningHistoryModelToEntity(model model.SubmitLearningHistoryModel) entity.SubmitLearningHistoryEntity {
	// Entity will never need password when converted from model
	// as the model comes from db, we shall not put password from db to entity
	return entity.SubmitLearningHistoryEntity{
		LearningHistoryID: model.LearningHistoryID,
		ChoreoDetailID:    model.ChoreoDetailID,
		UserID:            model.UserID,
		ChoreoDetail:      model.ChoreoDetail.String,
		Device:            model.Device.String,
		Progress:          model.Progress.Float64,
		RecordUrl:         model.RecordUrl.String,
		ThumbnailUrl:      model.ThumbnailUrl.String,
	}
}

func SubmitLearningHistoryEntityToModel(entity entity.SubmitLearningHistoryEntity) model.SubmitLearningHistoryModel {
	return model.SubmitLearningHistoryModel{
		LearningHistoryID: entity.LearningHistoryID,
		ChoreoDetailID:    entity.ChoreoDetailID,
		UserID:            entity.UserID,
		ChoreoDetail: sql.NullString{
			String: entity.ChoreoDetail,
			Valid:  true,
		},
		Device: sql.NullString{
			String: entity.Device,
			Valid:  true,
		},
		Progress: sql.NullFloat64{
			Float64: entity.Progress,
			Valid:   true,
		},
		RecordUrl: sql.NullString{
			String: entity.RecordUrl,
			Valid:  true,
		},
		ThumbnailUrl: sql.NullString{
			String: entity.ThumbnailUrl,
			Valid:  true,
		},
	}
}
