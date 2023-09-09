package delivery

import "github.com/Kora-Dance/koradance-backend/internal/common/router"

const (
	basePath = "/kora/learnHistory"

	getLearningHistoryList = basePath + "/getList"
	submitLearningHistory  = basePath + "/submit"
)

func (api LearningHistoryHandler) RegisterPath(router router.KoraRouter) {
	router.OPTIONS(getLearningHistoryList, api.middlewareM.CORS(nil))
	router.GET(getLearningHistoryList, api.middlewareM.AuthHandlerMiddleware(api.getHistoryListHandler))

	router.OPTIONS(submitLearningHistory, api.middlewareM.CORS(nil))
	router.POST(submitLearningHistory, api.middlewareM.AuthHandlerMiddleware(api.submitHistoryHandler))
}
