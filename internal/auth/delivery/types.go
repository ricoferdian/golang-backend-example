package delivery

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/auth"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/pkg/middleware"
)

type UserAuthHandler struct {
	middlewareM middleware.MiddlewareInterface
	userAuthUC  auth.UserAuthUseCase
	handlerCfg  *helper.HandlerConfig
}

func NewUserAuthHandler(middlewareM middleware.MiddlewareInterface, handlerCfg *helper.HandlerConfig, userAuthUC auth.UserAuthUseCase) common.APIPathProvider {
	return &UserAuthHandler{
		middlewareM: middlewareM,
		userAuthUC:  userAuthUC,
		handlerCfg:  handlerCfg,
	}
}
