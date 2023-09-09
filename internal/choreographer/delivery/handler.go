package delivery

import (
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/common/router"
)

const (
	basePath = "/kora/choreographer"

	getChoreographerList    = basePath + "/getList"
	getChoreographerDetail  = basePath + "/getDetail"
	upsertChoreographerData = basePath + "/upsert"
	deleteChoreographerData = basePath + "/delete"
)

func (api ChoreographerHandler) RegisterPath(router router.KoraRouter) {
	router.OPTIONS(getChoreographerList, api.middlewareM.CORS(nil))
	router.GET(getChoreographerList, api.middlewareM.CommonHandlerMiddleware(api.getChoreographerListHandler))

	router.OPTIONS(getChoreographerDetail, api.middlewareM.CORS(nil))
	router.GET(getChoreographerDetail, api.middlewareM.CommonHandlerMiddleware(api.getChoreographerByIDHandler))

	router.OPTIONS(upsertChoreographerData, api.middlewareM.CORS(nil))
	router.POST(upsertChoreographerData, api.middlewareM.InternalToolMiddleware(api.upsertChoreographerHandler, constants.BackOfficeStatic))

	router.OPTIONS(deleteChoreographerData, api.middlewareM.CORS(nil))
	router.DELETE(deleteChoreographerData, api.middlewareM.InternalToolMiddleware(api.deleteChoreographerByID, constants.BackOfficeStatic))
}
