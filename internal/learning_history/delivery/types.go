package delivery

import (
	"kora-backend/app/helper"
	"kora-backend/internal/domain/common"
	"kora-backend/internal/domain/learning_history"
)

type LearningHistoryHandler struct {
	middlewareM    common.MiddlewareInterface
	learnHistoryUC learning_history.LearningHistoryUseCase
	handlerCfg     *helper.HandlerConfig
}

func NewLearningHistoryHandler(middlewareM common.MiddlewareInterface, handlerCfg *helper.HandlerConfig, learnHistoryUC learning_history.LearningHistoryUseCase) common.APIPathProvider {
	return &LearningHistoryHandler{
		middlewareM:    middlewareM,
		learnHistoryUC: learnHistoryUC,
		handlerCfg:  handlerCfg,
	}
}
