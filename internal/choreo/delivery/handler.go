package delivery

import (
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/common/router"
)

const (
	basePath = "/kora/choreo"

	getChoreoList         = basePath + "/getList"
	getChoreoParentDetail = basePath + "/getDetail"
	getChoreoDetailList   = basePath + "/detail/getList"

	writePath          = "/upload/choreo"
	uploadChoreo       = writePath + "/parent/file"
	uploadChoreoDetail = writePath + "/detail/file"

	insertPath         = "/insert/choreo"
	insertChoreo       = insertPath + "/parent"
	insertChoreoDetail = insertPath + "/detail"

	updatePath         = "/update/choreo"
	updateChoreo       = updatePath + "/parent"
	updateChoreoDetail = updatePath + "/detail"

	deletePath         = "/delete/choreo"
	deleteChoreo       = deletePath + "/parent"
	deleteChoreoDetail = deletePath + "/detail"
)

func (api ChoreoHandler) RegisterPath(router router.KoraRouter) {
	router.OPTIONS(getChoreoList, api.middlewareM.CORS(nil))
	router.GET(getChoreoList, api.middlewareM.OptionalAuthHandlerMiddleware(api.getChoreoListHandler))

	router.OPTIONS(getChoreoParentDetail, api.middlewareM.CORS(nil))
	router.GET(getChoreoParentDetail, api.middlewareM.OptionalAuthHandlerMiddleware(api.getChoreoByIDHandler))

	router.OPTIONS(getChoreoDetailList, api.middlewareM.CORS(nil))
	router.GET(getChoreoDetailList, api.middlewareM.OptionalAuthHandlerMiddleware(api.getChoreoDetailListHandler))

	router.OPTIONS(uploadChoreo, api.middlewareM.CORS(nil))
	router.POST(uploadChoreo, api.middlewareM.InternalToolMiddleware(api.uploadChoreoContent, constants.BackOfficeStatic))

	router.OPTIONS(uploadChoreoDetail, api.middlewareM.CORS(nil))
	router.POST(uploadChoreoDetail, api.middlewareM.InternalToolMiddleware(api.uploadChoreoDetailContent, constants.BackOfficeStatic))

	router.OPTIONS(insertChoreo, api.middlewareM.CORS(nil))
	router.POST(insertChoreo, api.middlewareM.InternalToolMiddleware(api.insertChoreoHandler, constants.BackOfficeStatic))

	router.OPTIONS(insertChoreoDetail, api.middlewareM.CORS(nil))
	router.POST(insertChoreoDetail, api.middlewareM.InternalToolMiddleware(api.insertChoreoDetailHandler, constants.BackOfficeStatic))

	router.OPTIONS(updateChoreo, api.middlewareM.CORS(nil))
	router.POST(updateChoreo, api.middlewareM.InternalToolMiddleware(api.updateChoreoHandler, constants.BackOfficeStatic))

	router.OPTIONS(updateChoreoDetail, api.middlewareM.CORS(nil))
	router.POST(updateChoreoDetail, api.middlewareM.InternalToolMiddleware(api.updateChoreoDetailHandler, constants.BackOfficeStatic))

	router.OPTIONS(deleteChoreo, api.middlewareM.CORS(nil))
	router.DELETE(deleteChoreo, api.middlewareM.InternalToolMiddleware(api.deleteChoreoByID, constants.BackOfficeStatic))

	router.OPTIONS(deleteChoreoDetail, api.middlewareM.CORS(nil))
	router.DELETE(deleteChoreoDetail, api.middlewareM.InternalToolMiddleware(api.deleteChoreoDetailByID, constants.BackOfficeStatic))
}
