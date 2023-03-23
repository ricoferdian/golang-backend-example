package common

import (
	"github.com/gin-gonic/gin"
)

type MiddlewareInterface interface {
	CORS() gin.HandlerFunc
	CommonHandlerMiddleware(next gin.HandlerFunc) gin.HandlerFunc
	AuthHandlerMiddleware(next gin.HandlerFunc) gin.HandlerFunc
}
