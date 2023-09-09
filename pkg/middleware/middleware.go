package middleware

import (
	"errors"
	"fmt"
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/app/helper/http"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/common/router"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/Kora-Dance/koradance-backend/pkg/jwtauth"
	"github.com/gin-gonic/gin"
	http2 "net/http"
	"strings"
	"time"
)

type MiddlewareInterface interface {
	CORS(next router.HandleErrFunc) router.HandleErrFunc
	InternalToolMiddleware(next router.HandleErrFunc, staticTokenType string) router.HandleErrFunc
	StaticTokenMiddleware(next router.HandleErrFunc, staticTokenType string) router.HandleErrFunc
	CommonHandlerMiddleware(next router.HandleErrFunc) router.HandleErrFunc
	OptionalAuthHandlerMiddleware(next router.HandleErrFunc) router.HandleErrFunc
	AuthHandlerMiddleware(next router.HandleErrFunc) router.HandleErrFunc
}

// MiddlewareModule represent the data-struct for middleware
type MiddlewareModule struct {
	jwtAuth     *jwtauth.JwtAuthModule
	staticToken *helper.StaticTokenAuth
}

// NewMiddlewareModule initialize the middleware
func NewMiddlewareModule(jwtAuth *jwtauth.JwtAuthModule, staticToken *helper.StaticTokenAuth) MiddlewareInterface {
	return &MiddlewareModule{
		jwtAuth:     jwtAuth,
		staticToken: staticToken,
	}
}

func (m *MiddlewareModule) OptionalAuthHandlerMiddleware(next router.HandleErrFunc) router.HandleErrFunc {
	return func(c *gin.Context) (data interface{}, err error, tags []string) {
		startTime := time.Now()
		var (
			bearerToken = c.Request.Header.Get("Authorization")
		)
		if bearerToken != "" {
			// If token provided and valid, do some auth
			// If not valid, handle as unauthenticated user
			isValid, userEntity, _ := m.getUserAuthData(c, bearerToken, startTime)
			if isValid {
				c.Set(constants.CtxAuthUserData, userEntity)
			}
		}
		return next(c)
	}
}

func (m *MiddlewareModule) AuthHandlerMiddleware(next router.HandleErrFunc) router.HandleErrFunc {
	return func(c *gin.Context) (data interface{}, err error, tags []string) {
		startTime := time.Now()
		var (
			bearerToken = c.Request.Header.Get("Authorization")
		)
		if bearerToken == "" {
			http.WriteErrorResponseByCode(c, startTime, http.StatusUnauthorized)
			return
		}
		isValid, userEntity, err := m.getUserAuthData(c, bearerToken, startTime)
		if err != nil {
			if err.Error() == constants.ErrTokenExpired {
				http.WriteErrorResponseByCode(c, startTime, http.StatusTokenExpired)
				return
			}
			http.WriteErrorResponseByCode(c, startTime, http.StatusUnauthorized)
			return
		}
		if !isValid {
			http.WriteErrorResponseByCode(c, startTime, http.StatusUnauthorized)
			return
		}
		c.Set(constants.CtxAuthUserData, userEntity)
		return next(c)
	}
}

func (m *MiddlewareModule) getUserAuthData(c *gin.Context, bearerToken string, startTime time.Time) (isValid bool, userEntity *entity.AuthenticatedUserEntity, err error) {
	token := strings.SplitN(bearerToken, " ", 2)
	if len(token) < 2 {
		err = errors.New("invalid token")
		return
	}
	return m.jwtAuth.ValidateToken(token[1])
}

func (m *MiddlewareModule) InternalToolMiddleware(next router.HandleErrFunc, staticTokenType string) router.HandleErrFunc {
	return m.CORS(m.StaticTokenMiddleware(next, staticTokenType))
}

func (m *MiddlewareModule) StaticTokenMiddleware(next router.HandleErrFunc, staticTokenType string) router.HandleErrFunc {
	return func(c *gin.Context) (data interface{}, err error, tags []string) {
		startTime := time.Now()
		if staticTokenType == "" {
			http.WriteErrorResponseByCode(c, startTime, http.StatusUnauthorized)
		}
		var (
			staticToken = c.Request.Header.Get("Authorization")
		)
		// split string
		token := strings.SplitN(staticToken, " ", 2)
		if len(token) < 2 {
			err = fmt.Errorf("invalid token: %s", staticToken)
			http.WriteErrorResponseByCode(c, startTime, http.StatusUnauthorized)
			return
		}
		if token[1] != m.staticToken.GetToken(staticTokenType) {
			http.WriteErrorResponseByCode(c, startTime, http.StatusUnauthorized)
			return
		}
		return next(c)
	}
}

func (m *MiddlewareModule) CommonHandlerMiddleware(next router.HandleErrFunc) router.HandleErrFunc {
	return func(c *gin.Context) (data interface{}, err error, tags []string) {
		return next(c)
	}
}

func (m *MiddlewareModule) CORS(next router.HandleErrFunc) router.HandleErrFunc {
	return func(c *gin.Context) (data interface{}, err error, tags []string) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		if next != nil {
			return next(c)
		}

		c.AbortWithStatus(http2.StatusBadRequest)
		return
	}
}
