package delivery

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	"time"
)

func (api UserAuthHandler) userProfileHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	_, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()
	authUser, isOk := c.Value(constants.CtxAuthUserData).(*entity.AuthenticatedUserEntity)
	if !isOk {
		http.WriteErrorResponseByCode(c, startTime, http.StatusServerError)
		return
	}
	http.WriteSuccessResponse(c, startTime, *authUser)
	return
}
