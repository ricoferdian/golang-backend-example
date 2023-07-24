package delivery

import (
	"context"
	"errors"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	"time"
)

func validateLoginParam(c *gin.Context) (*entity.LoginUserEntity, error) {
	userIdentity := c.PostForm("user_identity")
	passIdentifier := c.PostForm("password_identifier")
	if userIdentity == "" || passIdentifier == "" {
		return nil, errors.New("user credentials are incomplete")
	}
	return &entity.LoginUserEntity{
		UserIdentity:       userIdentity,
		PasswordIdentifier: passIdentifier,
	}, nil
}

func (api UserAuthHandler) authUserLoginHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()
	loginData, err := validateLoginParam(c)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	data, err := api.userAuthUC.Login(ctx, *loginData)
	if err != nil || data == nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusAuthFailed)
		return
	}
	http.WriteSuccessResponse(c, startTime, &data)
	return
}
