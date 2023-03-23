package delivery

import (
	"github.com/gin-gonic/gin"
)

const (
	basePath = "/auth/user"

	authLogin    = basePath + "/login"
	authRegister = basePath + "/register"

	userProfile = basePath + "/profile"
)

func (api UserAuthHandler) RegisterPath(router *gin.Engine) {
	router.OPTIONS(authLogin, api.middlewareM.CORS())
	router.POST(authLogin, api.authUserLoginHandler)

	router.OPTIONS(authRegister, api.middlewareM.CORS())
	router.POST(authRegister, api.authUserRegisterHandler)

	router.OPTIONS(userProfile, api.middlewareM.CORS())
	router.GET(userProfile, api.middlewareM.AuthHandlerMiddleware(api.userProfileHandler))
}
