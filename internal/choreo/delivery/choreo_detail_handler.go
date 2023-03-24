package delivery

import (
	"context"
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"kora-backend/internal/entity"
	"strconv"
	"time"
)

func (api ChoreoHandler) getChoreoDetailListHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()
	choreoId, err := strconv.Atoi(c.Request.URL.Query().Get("choreo_id"))
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	filter := entity.ChoreoDetailFilterEntity{ChoreoID: int64(choreoId)}
	data, err := api.choreoUC.GetChoreoDetailByChoreoID(ctx, filter)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}
