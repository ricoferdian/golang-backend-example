package music

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
)

type MusicDatabaseRepo interface {
	GetMusicById(ctx context.Context, musicID int64) (*model.MusicModel, error)
	GetMusicByIdsMap(ctx context.Context, musicIDs []int64) (map[int64]model.MusicModel, error)
}

type MusicCacheRepo interface {
}

type MusicRepository interface {
	MusicDatabaseRepo
	MusicCacheRepo
}
