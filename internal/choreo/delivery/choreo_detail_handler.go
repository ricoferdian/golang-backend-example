package delivery

import (
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"kora-backend/internal/entity"
	"strconv"
	"time"
)

func (api ChoreoHandler) getChoreoDetailListHandler(c *gin.Context) {
	startTime := time.Now()
	ctx := c.Request.Context()
	choreoId, err := strconv.Atoi(c.Request.URL.Query().Get("choreo_id"))
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	filter := entity.ChoreoDetailFilterEntity{ChoreoID: int64(choreoId)}
	data, err := api.choreoUC.GetChoreoDetailByChoreoID(ctx, filter)
	if err != nil {
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}
