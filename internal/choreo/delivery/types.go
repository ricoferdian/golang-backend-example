package delivery

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/pkg/middleware"
)

type ChoreoHandler struct {
	middlewareM middleware.MiddlewareInterface
	choreoUC    choreo.ChoreoUseCase
	handlerCfg  *helper.HandlerConfig
}

func NewChoreoHandler(middlewareM middleware.MiddlewareInterface, handlerCfg *helper.HandlerConfig, choreoUC choreo.ChoreoUseCase) common.APIPathProvider {
	return &ChoreoHandler{
		middlewareM: middlewareM,
		choreoUC:    choreoUC,
		handlerCfg:  handlerCfg,
	}
}
