package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func (u UserAuthUseCaseImpl) AuthenticateWithOtp(ctx context.Context, request entity2.SecureOtpRequest) (*entity2.AuthUserResponseEntity, error) {
	identity, err := u.otpModule.ValidateOtp(ctx, request, request.Password)
	if err != nil {
		return nil, err
	}
	if identity == "" {
		return nil, errors.New(constants.ErrIdentityNotFound)
	}
	user, canRegister, err := u.loginOtp(ctx, request)
	if err != nil && !canRegister {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	return u.registerOtp(ctx, request)
}

func (u UserAuthUseCaseImpl) loginOtp(ctx context.Context, request entity2.SecureOtpRequest) (*entity2.AuthUserResponseEntity, bool, error) {
	filter := entity2.UserFilterEntity{
		UserIdentity: request.ReceiverIdentity,
		AuthType:     constants.AuthTypePasswordlessOtp,
	}
	userData, err := u.baseRepo.UserAuthRepository().GetSingleUserByUniqueFilter(ctx, filter)
	if err != nil {
		return nil, false, err
	}
	if userData == nil {
		return nil, true, errors.New(fmt.Sprintf("User credential not found for user identity %s", filter.UserIdentity))
	}
	authUser, err := u.generateTokenResponse(*userData)
	if err != nil {
		return nil, false, err
	}
	return authUser, false, nil
}

func (u UserAuthUseCaseImpl) registerOtp(ctx context.Context, request entity2.SecureOtpRequest) (*entity2.AuthUserResponseEntity, error) {
	user := entity2.UserEntity{
		UserIdentity:     request.ReceiverIdentity,
		PasslessIdentity: request.ReceiverIdentity,
	}
	singleUser, err := u.baseRepo.UserAuthRepository().InsertSingleUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return u.generateTokenResponse(*singleUser)
}
