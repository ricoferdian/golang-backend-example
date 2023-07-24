package usecase

import (
	"context"
	"fmt"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

const (
	sendOtpMsgFormat = "Hi, this is Kora\nYour OTP code is : %s\nPlease do not share this code with anyone !"
)

func (s UserAuthUseCaseImpl) SendOtpRequest(ctx context.Context, request entity.SecureOtpRequest) error {
	otp, err := s.otpModule.GenerateOtp(ctx, request)
	if err != nil {
		return err
	}
	otpMsg := fmt.Sprintf(sendOtpMsgFormat, otp)
	err = s.waModule.SendMessage(otpMsg, request.ReceiverIdentity)
	if err != nil {
		return err
	}
	return nil
}
