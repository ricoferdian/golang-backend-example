package delivery

import (
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/common"
)

type ChoreoHandler struct {
	middlewareM common.MiddlewareInterface
	choreoUC    choreo.ChoreoUseCase
}

func NewChoreoHandler(middlewareM common.MiddlewareInterface, choreoUC choreo.ChoreoUseCase) common.APIPathProvider {
	return &ChoreoHandler{
		middlewareM: middlewareM,
		choreoUC:    choreoUC,
	}
}
