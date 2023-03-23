package usecase

import (
	"kora-backend/internal/common/cryptography"
	"kora-backend/internal/common/jwtauth"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/common"
)

type UserAuthUseCaseImpl struct {
	baseRepo     common.BaseRepository
	jwtModule    *jwtauth.JwtAuthModule
	cryptoModule *cryptography.CryptographyModule
}

func NewUserAuthUseCase(baseRepo common.BaseRepository, jwtModule *jwtauth.JwtAuthModule, cryptoModule *cryptography.CryptographyModule) auth.UserAuthUseCase {
	return UserAuthUseCaseImpl{
		baseRepo:     baseRepo,
		jwtModule:    jwtModule,
		cryptoModule: cryptoModule,
	}
}
