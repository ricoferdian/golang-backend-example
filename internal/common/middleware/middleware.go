package middleware

import (
	"github.com/gin-gonic/gin"
	"kora-backend/app/helper/http"
	"kora-backend/internal/common/constants"
	"kora-backend/internal/common/jwtauth"
	"kora-backend/internal/entity"
	"strings"
	"time"
)

// MiddlewareModule represent the data-struct for middleware
type MiddlewareModule struct {
	jwtAuth *jwtauth.JwtAuthModule
}

// NewMiddlewareModule initialize the middleware
func NewMiddlewareModule(jwtAuth *jwtauth.JwtAuthModule) *MiddlewareModule {
	return &MiddlewareModule{
		jwtAuth: jwtAuth,
	}
}

func (m *MiddlewareModule) OptionalAuthHandlerMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		var (
			bearerToken = c.Request.Header.Get("Authorization")
		)
		if bearerToken != "" {
			// If token provided and valid, do some auth
			// If not valid, handle as unauthenticated user
			isValid, userEntity := m.getUserAuthData(c, bearerToken, startTime)
			if isValid {
				c.Set(constants.CtxAuthUserData, userEntity)
			}
		}
		next(c)
		return
	}
}

func (m *MiddlewareModule) AuthHandlerMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		var (
			bearerToken = c.Request.Header.Get("Authorization")
		)
		if bearerToken == "" {
			http.WriteErrorResponseByCode(c, startTime, http.StatusForbidden)
			return
		}
		isValid, userEntity := m.getUserAuthData(c, bearerToken, startTime)
		if !isValid {
			http.WriteErrorResponseByCode(c, startTime, http.StatusForbidden)
			return
		}
		c.Set(constants.CtxAuthUserData, userEntity)
		next(c)
		return
	}
}

func (m *MiddlewareModule) getUserAuthData(c *gin.Context, bearerToken string, startTime time.Time) (isValid bool, userEntity *entity.AuthenticatedUserEntity) {
	token := strings.SplitN(bearerToken, " ", 2)
	if len(token) < 2 {
		http.WriteErrorResponseByCode(c, startTime, http.StatusForbidden)
		return
	}
	isValid, userEntity, err := m.jwtAuth.ValidateToken(token[1])
	if err != nil {
		if err.Error() == constants.ErrTokenExpired {
			http.WriteErrorResponseByCode(c, startTime, http.StatusTokenExpired)
			return
		}
		http.WriteErrorResponseByCode(c, startTime, http.StatusForbidden)
		return
	}

	return isValid, userEntity
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
