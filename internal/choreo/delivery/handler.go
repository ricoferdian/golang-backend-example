package delivery

import (
	"github.com/gin-gonic/gin"
)

const (
	basePath = "/kora/choreo"

	getChoreoList       = basePath + "/getList"
	getChoreoDetailList = basePath + "/detail/getList"
)

func (api ChoreoHandler) RegisterPath(router *gin.Engine) {
	router.OPTIONS(getChoreoList, api.middlewareM.CORS())
	router.GET(getChoreoList, api.middlewareM.OptionalAuthHandlerMiddleware(api.getChoreoListHandler))

	router.OPTIONS(getChoreoDetailList, api.middlewareM.CORS())
	router.GET(getChoreoDetailList, api.middlewareM.OptionalAuthHandlerMiddleware(api.getChoreoDetailListHandler))
}
