package delivery

import (
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/common/router"
)

var (
	basePath = "/kora/music"

	getAllMusic = basePath + "/getList"
	upsertMusic = basePath + "/upsert"
	deleteMusic = basePath + "/delete"
)

func (api MusicHandler) RegisterPath(router router.KoraRouter) {
	router.OPTIONS(getAllMusic, api.middlewareM.CORS(nil))
	router.GET(getAllMusic, api.middlewareM.CommonHandlerMiddleware(api.getAllMusicHandler))

	router.OPTIONS(upsertMusic, api.middlewareM.CORS(nil))
	router.POST(upsertMusic, api.middlewareM.InternalToolMiddleware(api.upsertMusicHandler, constants.BackOfficeStatic))

	router.OPTIONS(deleteMusic, api.middlewareM.CORS(nil))
	router.DELETE(deleteMusic, api.middlewareM.InternalToolMiddleware(api.deleteMusicByID, constants.BackOfficeStatic))
}
