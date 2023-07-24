package usecase

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func (u UserAuthUseCaseImpl) Register(ctx context.Context, user entity.UserEntity) (*entity.AuthUserResponseEntity, error) {
	user.PasswordIdentifier = u.cryptoModule.HashPassword(user.PasswordIdentifier)
	singleUser, err := u.baseRepo.UserAuthRepository().InsertSingleUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return u.generateTokenResponse(*singleUser)
}
