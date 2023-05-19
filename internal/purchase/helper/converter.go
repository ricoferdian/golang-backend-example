package helper

import (
	"database/sql"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

func ChoreoPurchaseModelToEntity(model model.ChoreoPurchaseModel) entity.ChoreoPurchaseEntity {
	// Entity will never need password when converted from model
	// as the model comes from db, we shall not put password from db to entity
	return entity.ChoreoPurchaseEntity{
		ChoreoPurchaseID: model.ChoreoPurchaseID,
		ChoreoID:         model.ChoreoID,
		UserID:           model.UserID,
		Receipt:          model.Receipt.String,
		Status:           model.Status,
	}
}

func ChoreoPurchaseEntityToModel(entity entity.ChoreoPurchaseEntity) model.ChoreoPurchaseModel {
	// Entity will never need password when converted from model
	// as the model comes from db, we shall not put password from db to entity
	return model.ChoreoPurchaseModel{
		ChoreoPurchaseID: entity.ChoreoPurchaseID,
		ChoreoID:         entity.ChoreoID,
		UserID:           entity.UserID,
		Receipt: sql.NullString{
			String: entity.Receipt,
			Valid:  true,
		},
		Status: entity.Status,
	}
}
