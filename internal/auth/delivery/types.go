package delivery

import (
	"kora-backend/internal/domain/authdomain"
	"kora-backend/internal/domain/common"
)

type UserAuthHandler struct {
	middlewareM common.MiddlewareInterface
	userAuthUC  authdomain.UserAuthUseCase
}

func NewUserAuthHandler(middlewareM common.MiddlewareInterface, userAuthUC authdomain.UserAuthUseCase) common.APIPathProvider {
	return &UserAuthHandler{
		middlewareM: middlewareM,
		userAuthUC:  userAuthUC,
	}
}
