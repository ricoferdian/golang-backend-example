package delivery

import "github.com/Kora-Dance/koradance-backend/internal/common/router"

const (
	basePath = "/auth/user"

	authLogin    = basePath + "/login"
	authRegister = basePath + "/register"

	userProfile    = basePath + "/profile"
	userDeactivate = basePath + "/deactivate"
	userReactivate = basePath + "/reactivate"

	baseOtpPath = basePath + "/authOtp"

	otpRequest  = baseOtpPath + "/request"
	otpValidate = baseOtpPath + "/validate"
)

func (api UserAuthHandler) RegisterPath(router router.KoraRouter) {
	router.OPTIONS(authLogin, api.middlewareM.CORS())
	router.POST(authLogin, api.authUserLoginHandler)

	router.OPTIONS(authRegister, api.middlewareM.CORS())
	router.POST(authRegister, api.authUserRegisterHandler)

	router.OPTIONS(userProfile, api.middlewareM.CORS())
	router.GET(userProfile, api.middlewareM.AuthHandlerMiddleware(api.userProfileHandler))

	router.OPTIONS(otpRequest, api.middlewareM.CORS())
	router.POST(otpRequest, api.requestOtpHandler)

	router.OPTIONS(otpValidate, api.middlewareM.CORS())
	router.POST(otpValidate, api.authOtpHandler)

	router.OPTIONS(userDeactivate, api.middlewareM.CORS())
	router.GET(userDeactivate, api.middlewareM.AuthHandlerMiddleware(api.deactivateUserHandler))

	router.OPTIONS(userReactivate, api.middlewareM.CORS())
	router.GET(userReactivate, api.middlewareM.AuthHandlerMiddleware(api.reactivateUserHandler))
}
