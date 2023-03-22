package common

import "github.com/gin-gonic/gin"

type APIPathProvider interface {
	RegisterPath(router *gin.Engine)
}
