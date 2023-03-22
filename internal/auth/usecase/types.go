package usecase

import (
	"kora-backend/internal/domain/authdomain"
	"kora-backend/internal/domain/common"
)

type UserAuthUseCaseImpl struct {
	baseRepo common.BaseRepository
}

func NewUserAuthUseCase(baseRepo common.BaseRepository) authdomain.UserAuthUseCase {
	return UserAuthUseCaseImpl{
		baseRepo: baseRepo,
	}
}
