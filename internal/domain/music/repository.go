package music

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
)

type MusicDatabaseRepo interface {
	GetAllMusic(ctx context.Context) ([]model.MusicModel, error)
	GetMusicById(ctx context.Context, musicID int64) (*model.MusicModel, error)
	GetMusicByIdsMap(ctx context.Context, musicIDs []int64) (map[int64]model.MusicModel, error)
	UpsertMusic(ctx context.Context, music model.MusicModel) (*model.MusicModel, error)
	DeleteMusicByID(ctx context.Context, musicID int64) error
}

type MusicCacheRepo interface {
}

type MusicRepository interface {
	MusicDatabaseRepo
	MusicCacheRepo
}
