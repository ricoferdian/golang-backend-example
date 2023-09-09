package helper

import (
	"database/sql"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func UserModelToEntity(model model.RbacUserModel) entity.UserEntity {
	// Entity will never need password when converted from model
	// as the model comes from db, we shall not put password from db to entity
	return entity.UserEntity{
		UserID:           model.UserID,
		UserIdentity:     model.UserIdentity.String,
		PasslessIdentity: model.PasslessIdentity.String,
		FirstName:        model.FirstName.String,
		LastName:         model.LastName.String,
		UserType:         model.UserType.Int16,
	}
}

func UserEntityToModel(entity entity.UserEntity) model.RbacUserModel {
	return model.RbacUserModel{
		UserID: entity.UserID,
		UserIdentity: sql.NullString{
			String: entity.UserIdentity,
			Valid:  true,
		},
		PasslessIdentity: sql.NullString{
			String: entity.PasslessIdentity,
			Valid:  true,
		},
		FirstName: sql.NullString{
			String: entity.FirstName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: entity.LastName,
			Valid:  true,
		},
		UserType: sql.NullInt16{
			Int16: entity.UserType,
			Valid: true,
		},
	}
}
