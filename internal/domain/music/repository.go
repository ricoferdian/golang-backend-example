package music

import (
	"context"
	"kora-backend/internal/model"
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
