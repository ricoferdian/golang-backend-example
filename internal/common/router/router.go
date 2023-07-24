package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (r *routerInteractor) commonHandlerMiddleware(path string, handle HandleErrFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r.metricsClient.RecordCustomMetric(fmt.Sprintf("HttpRequest_%s", path), 1)
		var errHandler error
		defer func() {
			if errHandler != nil {
				log.Println("Error occurred: ", errHandler)
			}
		}()

		_, errHandler, _ = handle(ctx)
	}
}

func (r *routerInteractor) Use(errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.Use(r.commonHandlerMiddleware("", errFunc))
}

func (r *routerInteractor) Handle(s string, s2 string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.Handle(s, s2, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) Any(s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.Any(s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) GET(s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.GET(s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) POST(s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.POST(s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) DELETE(s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.DELETE(s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) PATCH(s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.PATCH(s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) PUT(s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.PUT(s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) OPTIONS(s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.OPTIONS(s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) HEAD(s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.HEAD(s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) Match(strings []string, s string, errFunc HandleErrFunc) gin.IRoutes {
	return r.routerLib.Match(strings, s, r.commonHandlerMiddleware(s, errFunc))
}

func (r *routerInteractor) StaticFile(s string, s2 string) gin.IRoutes {
	return r.routerLib.StaticFile(s, s2)
}

func (r *routerInteractor) StaticFileFS(s string, s2 string, system http.FileSystem) gin.IRoutes {
	return r.routerLib.StaticFileFS(s, s2, system)
}

func (r *routerInteractor) Static(s string, s2 string) gin.IRoutes {
	return r.routerLib.Static(s, s2)
}

func (r *routerInteractor) StaticFS(s string, system http.FileSystem) gin.IRoutes {
	return r.routerLib.StaticFS(s, system)
}
