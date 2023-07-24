package auth

import (
	"context"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
)

type UserAuthUseCase interface {
	Login(ctx context.Context, user entity2.LoginUserEntity) (*entity2.AuthUserResponseEntity, error)
	Register(ctx context.Context, user entity2.UserEntity) (*entity2.AuthUserResponseEntity, error)
	SendOtpRequest(ctx context.Context, request entity2.SecureOtpRequest) error
	AuthenticateWithOtp(ctx context.Context, request entity2.SecureOtpRequest) (*entity2.AuthUserResponseEntity, error)
	DeactivateUser(ctx context.Context, userID int64) error
	ReactivateUser(ctx context.Context, userID int64) error
}
