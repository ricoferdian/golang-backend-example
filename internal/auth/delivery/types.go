package delivery

import (
	"kora-backend/internal/domain/auth"
	"kora-backend/internal/domain/common"
)

type UserAuthHandler struct {
	middlewareM common.MiddlewareInterface
	userAuthUC  auth.UserAuthUseCase
}

func NewUserAuthHandler(middlewareM common.MiddlewareInterface, userAuthUC auth.UserAuthUseCase) common.APIPathProvider {
	return &UserAuthHandler{
		middlewareM: middlewareM,
		userAuthUC:  userAuthUC,
	}
}
