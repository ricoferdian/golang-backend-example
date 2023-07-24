package usecase

import (
	"context"
)

func (u UserAuthUseCaseImpl) DeactivateUser(ctx context.Context, userID int64) error {
	return u.baseRepo.UserAuthRepository().DeactivateUser(ctx, userID)
}

func (u UserAuthUseCaseImpl) ReactivateUser(ctx context.Context, userID int64) error {
	return u.baseRepo.UserAuthRepository().ReactivateUser(ctx, userID)
}
