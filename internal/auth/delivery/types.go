package delivery

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/common"
)

type UserAuthHandler struct {
	middlewareM common.MiddlewareInterface
	userAuthUC  auth.UserAuthUseCase
	handlerCfg  *helper.HandlerConfig
}

func NewUserAuthHandler(middlewareM common.MiddlewareInterface, handlerCfg *helper.HandlerConfig, userAuthUC auth.UserAuthUseCase) common.APIPathProvider {
	return &UserAuthHandler{
		middlewareM: middlewareM,
		userAuthUC:  userAuthUC,
		handlerCfg:  handlerCfg,
	}
}
