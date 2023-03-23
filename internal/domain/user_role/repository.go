package auth

import (
	"context"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

type UserRoleDatabaseRepo interface {
	GetAllRolesByUserId(ctx context.Context, entity entity.RoleFilterEntity) ([]model.RbacRoleModel, error)
}

type UserRoleCacheRepo interface {
}

type UserRoleRepository interface {
	UserRoleDatabaseRepo
	UserRoleCacheRepo
}
