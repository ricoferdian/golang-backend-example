package repository

import "kora-backend/internal/domain/music"

type MusicRepositoryImpl struct {
	music.MusicDatabaseRepo
	music.MusicCacheRepo
}

func NewMusicRepository(
	dbRepo music.MusicDatabaseRepo,
	redisRepo music.MusicCacheRepo,
) music.MusicRepository {
	return &MusicRepositoryImpl{
		dbRepo,
		redisRepo,
	}
}
