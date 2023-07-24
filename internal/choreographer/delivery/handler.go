package delivery

import (
	"github.com/Kora-Dance/koradance-backend/internal/common/router"
)

const (
	basePath = "/kora/choreographer"

	getChoreographerList   = basePath + "/getList"
	getChoreographerDetail = basePath + "/getDetail"
)

func (api ChoreographerHandler) RegisterPath(router router.KoraRouter) {
	router.OPTIONS(getChoreographerList, api.middlewareM.CORS())
	router.GET(getChoreographerList, api.middlewareM.CommonHandlerMiddleware(api.getChoreographerListHandler))

	router.OPTIONS(getChoreographerDetail, api.middlewareM.CORS())
	router.GET(getChoreographerDetail, api.middlewareM.CommonHandlerMiddleware(api.getChoreographerByIDHandler))
}
