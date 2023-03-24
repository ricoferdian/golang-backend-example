package delivery

import (
	"context"
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"time"
)

func (api ChoreoHandler) getChoreoListHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()
	data, err := api.choreoUC.GetChoreoList(ctx)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}
