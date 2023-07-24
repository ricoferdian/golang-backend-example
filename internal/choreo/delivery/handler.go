package delivery

import (
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
)

func (api ChoreoHandler) RegisterPath(router router.KoraRouter) {
	router.OPTIONS(getChoreoList, api.middlewareM.CORS())
	router.GET(getChoreoList, api.middlewareM.OptionalAuthHandlerMiddleware(api.getChoreoListHandler))

	router.OPTIONS(getChoreoParentDetail, api.middlewareM.CORS())
	router.GET(getChoreoParentDetail, api.middlewareM.OptionalAuthHandlerMiddleware(api.getChoreoByIDHandler))

	router.OPTIONS(getChoreoDetailList, api.middlewareM.CORS())
	router.GET(getChoreoDetailList, api.middlewareM.OptionalAuthHandlerMiddleware(api.getChoreoDetailListHandler))

	router.OPTIONS(uploadChoreo, api.middlewareM.CORS())
	router.POST(uploadChoreo, api.middlewareM.CommonHandlerMiddleware(api.uploadChoreoContent))

	router.OPTIONS(uploadChoreoDetail, api.middlewareM.CORS())
	router.POST(uploadChoreoDetail, api.middlewareM.CommonHandlerMiddleware(api.uploadChoreoDetailContent))

	router.OPTIONS(insertChoreo, api.middlewareM.CORS())
	router.POST(insertChoreo, api.middlewareM.CommonHandlerMiddleware(api.insertChoreoHandler))

	router.OPTIONS(insertChoreoDetail, api.middlewareM.CORS())
	router.POST(insertChoreoDetail, api.middlewareM.CommonHandlerMiddleware(api.insertChoreoDetailHandler))

	router.OPTIONS(updateChoreo, api.middlewareM.CORS())
	router.POST(updateChoreo, api.middlewareM.CommonHandlerMiddleware(api.updateChoreoHandler))

	router.OPTIONS(updateChoreoDetail, api.middlewareM.CORS())
	router.POST(updateChoreoDetail, api.middlewareM.CommonHandlerMiddleware(api.updateChoreoDetailHandler))
}
