package delivery

import (
	"context"
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"kora-backend/internal/common/constants"
	"kora-backend/internal/entity"
	"time"
)

func (api ChoreoHandler) getChoreoListHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()
	data, err := api.getChoreoListWithOptionalAuth(c, ctx)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func (api ChoreoHandler) getChoreoListWithOptionalAuth(c *gin.Context, ctx context.Context) ([]entity.ChoreographyEntity, error) {
	authData, isOk := c.Value(constants.CtxAuthUserData).(*entity.AuthenticatedUserEntity)
	if !isOk {
		return api.choreoUC.GetChoreoList(ctx)
	}
	return api.choreoUC.GetChoreoListWithUserContent(ctx, authData.UserID)
}
