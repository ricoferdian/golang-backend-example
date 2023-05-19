package delivery

import (
	"context"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"kora-backend/app/helper/http"
	"kora-backend/internal/common/constants"
	"kora-backend/internal/entity"
	"log"
	"time"
)

func validateSubmitParam(c *gin.Context) (*entity.VerifyPaymentAppleIAPEntity, error) {
	var paymentData entity.VerifyPaymentAppleIAPEntity
	if err := c.ShouldBind(&paymentData); err != nil {
		return nil, err
	}

	validEng := validator.New()
	err := validEng.Struct(paymentData)
	if err != nil {
		return nil, err
	}
	return &paymentData, nil
}

func (api ChoreoPurchaseHandler) verifyPurchaseHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	// Validate user
	authData, isOk := c.Value(constants.CtxAuthUserData).(*entity.AuthenticatedUserEntity)
	if !isOk {
		http.WriteErrorResponseByCode(c, startTime, http.StatusServerError)
		return
	}

	// Validate param
	paymentData, err := validateSubmitParam(c)
	if err != nil {
		log.Println(err)
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}

	data, err := api.purchaseUC.VerifyPurchaseChoreo(ctx, authData.UserID, *paymentData)
	if err != nil {
		log.Println(err)
		http.WriteErrorResponseByCode(c, startTime, http.StatusServerError)
		return
	}
	http.WriteSuccessResponse(c, startTime, &data)
	return
}
