package auth

import (
	"context"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

type UserAuthDatabaseRepo interface {
	GetSingleUserByUniqueFilter(ctx context.Context, entity entity.UserFilterEntity) (*model.RbacUserModel, error)
	InsertSingleUser(ctx context.Context, entity entity.UserEntity) (*model.RbacUserModel, error)
}

type UserAuthCacheRepo interface {
}

type UserAuthRepository interface {
	UserAuthDatabaseRepo
	UserAuthCacheRepo
}
