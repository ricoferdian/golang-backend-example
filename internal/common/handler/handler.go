package handler

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func GenericDeleteHandler(c *gin.Context, timeout int, key string, handle func(ctx context.Context, id int64) error) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(timeout))
	defer cancel()

	startTime := time.Now()

	id, err := strconv.Atoi(c.Request.URL.Query().Get(key))
	if id == 0 || err != nil {
		log.Println("[DeleteHandler] error parse query param primary key", err)
		http.WriteErrorResponseObj(c, startTime, http.StatusInvalidRequest, http.ErrorResponse{
			Code:       http.StatusInvalidRequest,
			ErrMessage: err.Error(),
			ErrReason:  err.Error(),
		})
	}
	err = handle(ctx, int64(id))
	if err != nil {
		http.WriteErrorResponseObj(c, startTime, http.StatusServerError, http.ErrorResponse{
			Code:       http.StatusServerError,
			ErrMessage: err.Error(),
			ErrReason:  err.Error(),
		})
		return
	}
	http.WriteSuccessResponse(c, startTime, nil)
	return
}
