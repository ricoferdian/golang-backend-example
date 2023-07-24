package repository

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/like_save"
)

type LikeSaveRepositoryImpl struct {
	like_save.LikeSaveDatabaseRepo
	like_save.LikeSaveCacheRepo
}

func NewLikeSaveRepository(
	dbRepo like_save.LikeSaveDatabaseRepo,
	redisRepo like_save.LikeSaveCacheRepo,
) like_save.LikeSaveRepository {
	return &LikeSaveRepositoryImpl{
		dbRepo,
		redisRepo,
	}
}
