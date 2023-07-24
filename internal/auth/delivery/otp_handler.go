package delivery

import (
	"context"
	"errors"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// adjustPhoneNumberWithCountryCode used to adjust phone number with country code prefix if country code provided
func adjustPhoneNumberWithCountryCode(countryCode, phoneNumber string) string {
	if phoneNumber[0] == '0' {
		return countryCode + phoneNumber[1:]
	}
	return countryCode + phoneNumber
}

// getOtpRequestParam used to get otp request param from request
func getOtpRequestParam(c *gin.Context) (*entity.SecureOtpRequest, error) {
	countryCode := c.PostForm("country_code")
	receiver := c.PostForm("phone_number")

	if receiver == "" || countryCode == "" {
		return nil, errors.New("invalid otp request")
	}
	receiver = adjustPhoneNumberWithCountryCode(countryCode, receiver)
	return &entity.SecureOtpRequest{
		ReceiverIdentity:  receiver,
		TransferMediaType: constants.AuthOtpMediaTypeWhatsapp,
	}, nil
}

// requestOtpHandler used to handle request otp request
func (api UserAuthHandler) requestOtpHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	otpReqParam, err := getOtpRequestParam(c)
	if err != nil || otpReqParam == nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}

	err = api.userAuthUC.SendOtpRequest(ctx, *otpReqParam)
	if err != nil {
		log.Println("[UserAuthHandler][requestOtpHandler] err", err.Error())
		http.WriteErrorResponseByCode(c, startTime, http.StatusServerError)
		return
	}
	http.WriteSuccessResponse(c, startTime, nil)
	return
}

// getOtpValidationParam used to get otp validation param from request
func getOtpValidationParam(c *gin.Context) (*entity.SecureOtpRequest, error) {
	param, err := getOtpRequestParam(c)
	if err != nil || param == nil {
		return nil, err
	}
	password := c.PostForm("password")

	if password == "" {
		return nil, errors.New("invalid otp request")
	}
	param.Password = password
	return param, nil
}

// authOtpHandler used to authenticate user with otp
func (api UserAuthHandler) authOtpHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	otpReqParam, err := getOtpValidationParam(c)
	if err != nil || otpReqParam == nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}

	data, err := api.userAuthUC.AuthenticateWithOtp(ctx, *otpReqParam)
	if err != nil {
		if err.Error() == constants.ErrIdentityNotFound {
			http.WriteErrorResponseByCode(c, startTime, http.StatusVerificationFailed)
			return
		}
		log.Println("[UserAuthHandler][authOtpHandler] err", err.Error())
		http.WriteErrorResponseByCode(c, startTime, http.StatusServerError)
		return
	}
	http.WriteSuccessResponse(c, startTime, &data)
	return
}
