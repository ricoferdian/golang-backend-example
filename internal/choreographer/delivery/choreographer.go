package delivery

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/common/handler"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

var (
	TypeChoreographer = 3

	mapPrimaryKey = map[int]string{
		TypeChoreographer: "choreographer_id",
	}
	mapFileCategory = map[int][]int{
		TypeChoreographer: {constants.FileCategoryThumbnailImage},
	}
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

func (api ChoreographerHandler) upsertChoreographerHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	var choreographerData entity.ChoreographerEntity
	err := c.BindJSON(&choreographerData)
	if err != nil {
		log.Println("[ChoreoHandler] error parse body", err)
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	data, err := api.choreographerUC.UpsertChoreographer(ctx, choreographerData)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func (api ChoreographerHandler) deleteChoreographerByID(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return handler.GenericDeleteHandler(c, api.handlerCfg.Timeout, "choreographer_id", api.choreographerUC.DeleteChoreographerByID)
}

func (api ChoreographerHandler) uploadChoreographerContent(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return handler.GenericFileUploadHandler(c, handler.TypeChoreographer, api.handlerCfg.Timeout, api.choreographerUC.UploadChoreographerContent)
}
