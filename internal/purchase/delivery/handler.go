package delivery

import "github.com/Kora-Dance/koradance-backend/internal/common/router"

const (
	basePath = "/kora/purchasedContent"

	getPurchasedContent    = basePath + "/getList"
	verifyPurchasedContent = basePath + "/verify"
)

func (api ChoreoPurchaseHandler) RegisterPath(router router.KoraRouter) {
	router.OPTIONS(getPurchasedContent, api.middlewareM.CORS())
	router.GET(getPurchasedContent, api.middlewareM.AuthHandlerMiddleware(api.getPurchasedChoreoListHandler))

	router.OPTIONS(verifyPurchasedContent, api.middlewareM.CORS())
	router.POST(verifyPurchasedContent, api.middlewareM.AuthHandlerMiddleware(api.verifyPurchaseHandler))
}
