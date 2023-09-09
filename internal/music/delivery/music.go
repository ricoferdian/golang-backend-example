package delivery

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/handler"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func (api *MusicHandler) getAllMusicHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	data, err := api.musicUC.GetAllMusic(ctx)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func (api *MusicHandler) upsertMusicHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()

	var musicData entity.MusicEntity
	err := c.BindJSON(&musicData)
	if err != nil {
		log.Println("[MusicHandler] error parse body", err)
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	data, err := api.musicUC.UpsertMusic(ctx, musicData)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}

func (api *MusicHandler) deleteMusicByID(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	return handler.GenericDeleteHandler(c, api.handlerCfg.Timeout, "music_id", api.musicUC.DeleteMusicByID)
}
