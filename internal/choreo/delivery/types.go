package delivery

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/choreo"
	"kora-backend/internal/domain/common"
)

type ChoreoHandler struct {
	middlewareM common.MiddlewareInterface
	choreoUC    choreo.ChoreoUseCase
	handlerCfg  *helper.HandlerConfig
}

func NewChoreoHandler(middlewareM common.MiddlewareInterface, handlerCfg *helper.HandlerConfig, choreoUC choreo.ChoreoUseCase) common.APIPathProvider {
	return &ChoreoHandler{
		middlewareM: middlewareM,
		choreoUC:    choreoUC,
		handlerCfg:  handlerCfg,
	}
}
