package delivery

import "github.com/gin-gonic/gin"

const (
	basePath = "/kora/purchasedContent"

	getPurchasedContent    = basePath + "/getList"
	verifyPurchasedContent = basePath + "/verify"
)

func (api ChoreoPurchaseHandler) RegisterPath(router *gin.Engine) {
	router.OPTIONS(getPurchasedContent, api.middlewareM.CORS())
	router.GET(getPurchasedContent, api.middlewareM.AuthHandlerMiddleware(api.getPurchasedChoreoListHandler))

	router.OPTIONS(verifyPurchasedContent, api.middlewareM.CORS())
	router.POST(verifyPurchasedContent, api.middlewareM.AuthHandlerMiddleware(api.verifyPurchaseHandler))
}
