package auth

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
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
