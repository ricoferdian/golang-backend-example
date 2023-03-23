package delivery

import (
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"kora-backend/internal/common/constants"
	"kora-backend/internal/entity"
	"time"
)

func (api UserAuthHandler) userProfileHandler(c *gin.Context) {
	startTime := time.Now()
	data, isOk := c.Value(constants.CtxAuthUserData).(*entity.AuthenticatedUserEntity)
	if !isOk {
		http.WriteErrorResponseByCode(c, startTime, http.StatusServerError)
		return
	}
	http.WriteSuccessResponse(c, startTime, *data)
	return
}
