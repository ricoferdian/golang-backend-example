package usecase

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/internal/domain/music"
)

type MusicUseCaseImpl struct {
	baseRepo common.BaseRepository
}

func NewMusicUseCase(baseRepo common.BaseRepository) music.MusicUseCase {
	return &MusicUseCaseImpl{
		baseRepo: baseRepo,
	}
}
