package music

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

type MusicUseCase interface {
	GetMusicByID(ctx context.Context, musicID int64) (entity.MusicEntity, error)
	GetAllMusic(ctx context.Context) ([]entity.MusicEntity, error)
	UpsertMusic(ctx context.Context, music entity.MusicEntity) (entity.MusicEntity, error)
	DeleteMusicByID(ctx context.Context, musicID int64) error
}
