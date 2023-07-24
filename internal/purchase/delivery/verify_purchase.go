package delivery

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"log"
	"time"
)

func validateSubmitParam(c *gin.Context) (*entity2.VerifyPaymentAppleIAPEntity, error) {
	var paymentData entity2.VerifyPaymentAppleIAPEntity
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

func (api ChoreoPurchaseHandler) verifyPurchaseHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	// Validate user
	authData, isOk := c.Value(constants.CtxAuthUserData).(*entity2.AuthenticatedUserEntity)
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
