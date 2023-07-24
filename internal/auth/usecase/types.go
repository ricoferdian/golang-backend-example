package usecase

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/auth"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/pkg/cryptography"
	"github.com/Kora-Dance/koradance-backend/pkg/jwtauth"
	"github.com/Kora-Dance/koradance-backend/pkg/secure_otp"
	"github.com/Kora-Dance/koradance-backend/pkg/whatsapp"
)

type UserAuthUseCaseImpl struct {
	baseRepo     common.BaseRepository
	jwtModule    *jwtauth.JwtAuthModule
	cryptoModule *cryptography.CryptographyModule
	otpModule    *secure_otp.SecureOtpModule
	waModule     *whatsapp.WhatsappModule
}

func NewUserAuthUseCase(baseRepo common.BaseRepository, jwtModule *jwtauth.JwtAuthModule, cryptoModule *cryptography.CryptographyModule, otpModule *secure_otp.SecureOtpModule, waModule *whatsapp.WhatsappModule) auth.UserAuthUseCase {
	return &UserAuthUseCaseImpl{
		baseRepo:     baseRepo,
		jwtModule:    jwtModule,
		cryptoModule: cryptoModule,
		otpModule:    otpModule,
		waModule:     waModule,
	}
}
