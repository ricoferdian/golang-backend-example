package delivery

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"time"
)

func validateSubmitParam(c *gin.Context) (*entity.SubmitLearningHistoryEntity, error) {
	var history entity.SubmitLearningHistoryEntity
	if err := c.ShouldBind(&history); err != nil {
		return nil, err
	}

	validEng := validator.New()
	err := validEng.Struct(history)
	if err != nil {
		return nil, err
	}
	return &history, nil
}

func (api LearningHistoryHandler) submitHistoryHandler(c *gin.Context) (metricsData interface{}, metricsErr error, metricsTags []string) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Millisecond*time.Duration(api.handlerCfg.Timeout))
	defer cancel()

	startTime := time.Now()
	historyData, err := validateSubmitParam(c)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusInvalidRequest)
		return
	}
	data, err := api.learnHistoryUC.SubmitLearningHistory(ctx, *historyData)
	if err != nil {
		http.WriteErrorResponseByCode(c, startTime, http.StatusServerError)
		return
	}
	http.WriteSuccessResponse(c, startTime, &data)
	return
}
