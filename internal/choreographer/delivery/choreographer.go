package delivery

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func (api ChoreographerHandler) getChoreographerListHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	data, err := api.choreographerUC.GetChoreographerList(ctx)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func (api ChoreographerHandler) getChoreographerByIDHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	choreographerID, err := strconv.Atoi(c.Request.URL.Query().Get("choreographer_id"))
	if err != nil {
		log.Println("[ChoreoHandler] error parse query param choreographer_id", err)
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
	}
	filter := entity.ChoreographerFilter{ChoreographerID: int64(choreographerID)}
	data, err := api.choreographerUC.GetChoreographerByID(ctx, filter)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}
