package router

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
)

type HandleErrFunc func(ctx *gin.Context) (data interface{}, err error, tags []string)

type internalRouter interface {
	Use(...gin.HandlerFunc) gin.IRoutes
	Handle(string, string, ...gin.HandlerFunc) gin.IRoutes
	Any(string, ...gin.HandlerFunc) gin.IRoutes
	GET(string, ...gin.HandlerFunc) gin.IRoutes
	POST(string, ...gin.HandlerFunc) gin.IRoutes
	DELETE(string, ...gin.HandlerFunc) gin.IRoutes
	PATCH(string, ...gin.HandlerFunc) gin.IRoutes
	PUT(string, ...gin.HandlerFunc) gin.IRoutes
	OPTIONS(string, ...gin.HandlerFunc) gin.IRoutes
	HEAD(string, ...gin.HandlerFunc) gin.IRoutes
	Match([]string, string, ...gin.HandlerFunc) gin.IRoutes
	StaticFile(string, string) gin.IRoutes
	StaticFileFS(string, string, http.FileSystem) gin.IRoutes
	Static(string, string) gin.IRoutes
	StaticFS(string, http.FileSystem) gin.IRoutes
}

type KoraRouter interface {
	Use(HandleErrFunc) gin.IRoutes
	Handle(string, string, HandleErrFunc) gin.IRoutes
	Any(string, HandleErrFunc) gin.IRoutes
	GET(string, HandleErrFunc) gin.IRoutes
	POST(string, HandleErrFunc) gin.IRoutes
	DELETE(string, HandleErrFunc) gin.IRoutes
	PATCH(string, HandleErrFunc) gin.IRoutes
	PUT(string, HandleErrFunc) gin.IRoutes
	OPTIONS(string, HandleErrFunc) gin.IRoutes
	HEAD(string, HandleErrFunc) gin.IRoutes
	Match([]string, string, HandleErrFunc) gin.IRoutes
	StaticFile(string, string) gin.IRoutes
	StaticFileFS(string, string, http.FileSystem) gin.IRoutes
	Static(string, string) gin.IRoutes
	StaticFS(string, http.FileSystem) gin.IRoutes
}

func NewRouter(routerLib internalRouter, metricsClient *newrelic.Application) KoraRouter {
	return &routerInteractor{
		routerLib:     routerLib,
		metricsClient: metricsClient,
	}
}

type routerInteractor struct {
	routerLib     internalRouter
	metricsClient *newrelic.Application
}
