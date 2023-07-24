package delivery

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/internal/domain/learning_history"
	"github.com/Kora-Dance/koradance-backend/pkg/middleware"
)

type LearningHistoryHandler struct {
	middlewareM    middleware.MiddlewareInterface
	learnHistoryUC learning_history.LearningHistoryUseCase
	handlerCfg     *helper.HandlerConfig
}

func NewLearningHistoryHandler(middlewareM middleware.MiddlewareInterface, handlerCfg *helper.HandlerConfig, learnHistoryUC learning_history.LearningHistoryUseCase) common.APIPathProvider {
	return &LearningHistoryHandler{
		middlewareM:    middlewareM,
		learnHistoryUC: learnHistoryUC,
		handlerCfg:     handlerCfg,
	}
}
