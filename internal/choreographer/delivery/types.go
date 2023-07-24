package delivery

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreographer"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/pkg/middleware"
)

type ChoreographerHandler struct {
	middlewareM     middleware.MiddlewareInterface
	choreographerUC choreographer.ChoreographerUseCase
	handlerCfg      *helper.HandlerConfig
}

func NewChoreographerHandler(middlewareM middleware.MiddlewareInterface, handlerCfg *helper.HandlerConfig, choreographerUC choreographer.ChoreographerUseCase) common.APIPathProvider {
	return &ChoreographerHandler{
		middlewareM:     middlewareM,
		choreographerUC: choreographerUC,
		handlerCfg:      handlerCfg,
	}
}
