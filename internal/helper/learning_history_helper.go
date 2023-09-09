package helper

import (
	"database/sql"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func LearningHistoryModelToEntity(model model.LearningHistoryModel) entity2.LearningHistoryEntity {
	// Entity will never need password when converted from model
	// as the model comes from db, we shall not put password from db to entity
	return entity2.LearningHistoryEntity{
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

func SubmitLearningHistoryModelToEntity(model model.SubmitLearningHistoryModel) entity2.SubmitLearningHistoryEntity {
	// Entity will never need password when converted from model
	// as the model comes from db, we shall not put password from db to entity
	return entity2.SubmitLearningHistoryEntity{
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

func SubmitLearningHistoryEntityToModel(entity entity2.SubmitLearningHistoryEntity) model.SubmitLearningHistoryModel {
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
