package delivery

import (
	"github.com/gin-gonic/gin"
)

const (
	basePath = "/kora/learnHistory"

	getLearningHistoryList = basePath + "/getList"
	submitLearningHistory  = basePath + "/submit"
)

func (api LearningHistoryHandler) RegisterPath(router *gin.Engine) {
	router.OPTIONS(getLearningHistoryList, api.middlewareM.CORS())
	router.GET(getLearningHistoryList, api.middlewareM.AuthHandlerMiddleware(api.getHistoryListHandler))

	router.OPTIONS(submitLearningHistory, api.middlewareM.CORS())
	router.POST(submitLearningHistory, api.middlewareM.AuthHandlerMiddleware(api.submitHistoryHandler))
}
