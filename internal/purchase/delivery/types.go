package delivery

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/common"
	"kora-backend/internal/domain/purchase"
)

type ChoreoPurchaseHandler struct {
	middlewareM common.MiddlewareInterface
	purchaseUC  purchase.ChoreoPurchaseUseCase
	handlerCfg  *helper.HandlerConfig
}

func NewChoreoPurchaseHandler(middlewareM common.MiddlewareInterface, handlerCfg *helper.HandlerConfig, purchaseUC purchase.ChoreoPurchaseUseCase) common.APIPathProvider {
	return &ChoreoPurchaseHandler{
		middlewareM: middlewareM,
		purchaseUC:  purchaseUC,
		handlerCfg:  handlerCfg,
	}
}
