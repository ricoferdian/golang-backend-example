package auth

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

type UserAuthDatabaseRepo interface {
	GetSingleUserByUniqueFilter(ctx context.Context, entity entity.UserFilterEntity) (*model.RbacUserModel, error)
	InsertSingleUser(ctx context.Context, entity entity.UserEntity) (*model.RbacUserModel, error)
	DeactivateUser(ctx context.Context, userID int64) error
	ReactivateUser(ctx context.Context, userID int64) error
}

type UserAuthCacheRepo interface {
}

type UserAuthRepository interface {
	UserAuthDatabaseRepo
	UserAuthCacheRepo
}
