package delivery

import (
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"time"
)

func (api ChoreoHandler) getChoreoListHandler(c *gin.Context) {
	startTime := time.Now()
	ctx := c.Request.Context()
	data, err := api.choreoUC.GetChoreoList(ctx)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusNotFound)
		return
	}
	http.WriteSuccessResponse(c, startTime, data)
	return
}
