package delivery

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/internal/domain/purchase"
	"github.com/Kora-Dance/koradance-backend/pkg/middleware"
)

type ChoreoPurchaseHandler struct {
	middlewareM middleware.MiddlewareInterface
	purchaseUC  purchase.ChoreoPurchaseUseCase
	handlerCfg  *helper.HandlerConfig
}

func NewChoreoPurchaseHandler(middlewareM middleware.MiddlewareInterface, handlerCfg *helper.HandlerConfig, purchaseUC purchase.ChoreoPurchaseUseCase) common.APIPathProvider {
	return &ChoreoPurchaseHandler{
		middlewareM: middlewareM,
		purchaseUC:  purchaseUC,
		handlerCfg:  handlerCfg,
	}
}
