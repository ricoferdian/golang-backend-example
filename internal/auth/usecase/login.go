package usecase

import (
	"context"
	"errors"
	"fmt"
	"kora-backend/internal/entity"
)

func (u UserAuthUseCaseImpl) Login(ctx context.Context, user entity.LoginUserEntity) (*entity.AuthUserResponseEntity, error) {
	filter := entity.UserFilterEntity{
		UserIdentity: user.UserIdentity,
	}
	userData, err := u.baseRepo.UserAuthRepository().GetSingleUserByUniqueFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, errors.New(fmt.Sprintf("User credential not found for user identity %s", user.UserIdentity))
	}
	isValid := u.cryptoModule.ComparePassword(userData.HashPasswordIdentifier, user.PasswordIdentifier)
	if !isValid {
		return nil, errors.New(fmt.Sprintf("User password does not match for identity %s", user.UserIdentity))
	}
	return u.generateTokenResponse(*userData)
}
