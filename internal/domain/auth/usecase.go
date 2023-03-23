package auth

import (
	"context"
	"kora-backend/internal/entity"
)

type UserAuthUseCase interface {
	Login(ctx context.Context, user entity.LoginUserEntity) (*entity.AuthUserResponseEntity, error)
	Register(ctx context.Context, user entity.UserEntity) (*entity.AuthUserResponseEntity, error)
}
