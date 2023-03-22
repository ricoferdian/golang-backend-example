package middleware

import (
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"time"
)

// MiddlewareModule represent the data-struct for middleware
type MiddlewareModule struct {
	// another stuff , may be needed by middleware
}

func (m *MiddlewareModule) AuthenticatedHandlerMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	startTime := time.Now()
	return func(c *gin.Context) {
		var (
			bearerToken = c.Request.Header.Get("Authorization")
		)
		if bearerToken == "" {
			http.WriteErrorResponseByCode(c, startTime, http.StatusForbidden)
			return
		}
		next(c)
	}
}

func (m *MiddlewareModule) CommonHandlerMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		next(c)
	}
}

func (m *MiddlewareModule) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *MiddlewareModule {
	return &MiddlewareModule{}
}
