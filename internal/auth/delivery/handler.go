package delivery

import (
	"github.com/gin-gonic/gin"
)

const (
	basePath = "/auth/user"

	getUserProfile = basePath + "/profile"
)

func (api UserAuthHandler) RegisterPath(router *gin.Engine) {
	router.OPTIONS(getUserProfile, api.middlewareM.CORS())
	router.GET(getUserProfile, api.getUserProfile)
}
