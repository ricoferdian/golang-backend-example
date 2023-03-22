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
	router.GET(getChoreoList, api.middlewareM.AuthenticatedHandlerMiddleware(api.getChoreoList))

	router.OPTIONS(getChoreoDetailList, api.middlewareM.CORS())
	router.GET(getChoreoDetailList, api.middlewareM.AuthenticatedHandlerMiddleware(api.getChoreoDetailList))
}
