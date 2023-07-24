package delivery

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (api ChoreoHandler) getChoreoDetailListHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()
	choreoId, err := strconv.Atoi(c.Request.URL.Query().Get("choreo_id"))
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	filter := entity2.ChoreoDetailFilterEntity{ChoreoID: int64(choreoId)}
	data, err := api.getChoreoDetailListWithOptionalAuth(c, ctx, filter)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func (api ChoreoHandler) getChoreoDetailListWithOptionalAuth(c *gin.Context, ctx context.Context, filter entity2.ChoreoDetailFilterEntity) ([]entity2.ChoreographyDetailEntity, error) {
	authData, isOk := c.Value(constants.CtxAuthUserData).(*entity2.AuthenticatedUserEntity)
	if !isOk {
		return api.choreoUC.GetChoreoDetailByChoreoID(ctx, filter)
	}
	return api.choreoUC.GetChoreoDetailByChoreoIDWithUserContent(ctx, authData.UserID, filter)
}
