package delivery

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/internal/domain/music"
	"github.com/Kora-Dance/koradance-backend/pkg/middleware"
)

type MusicHandler struct {
	middlewareM middleware.MiddlewareInterface
	musicUC     music.MusicUseCase
	handlerCfg  *helper.HandlerConfig
}

func NewMusicHandler(middlewareM middleware.MiddlewareInterface, handlerCfg *helper.HandlerConfig, musicUC music.MusicUseCase) common.APIPathProvider {
	return &MusicHandler{
		middlewareM: middlewareM,
		musicUC:     musicUC,
		handlerCfg:  handlerCfg,
	}
}
