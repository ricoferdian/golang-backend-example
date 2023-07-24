package like_save

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
)

type LikeSaveDatabaseRepo interface {
	// Saves
	GetSavedChoreoByUserID(ctx context.Context, userID int64) ([]model.ChoreoSaveModel, error)
	GetSavedChoreoByChoreoIDsMap(ctx context.Context, userID int64, choreoIDs []int64) (map[int64]model.ChoreoSaveModel, error)
	// Likes
	GetLikedChoreoByChoreoIDsMap(ctx context.Context, userID int64, choreoIDs []int64) (map[int64]model.ChoreoLikeModel, error)
}

type LikeSaveCacheRepo interface {
}

type LikeSaveRepository interface {
	LikeSaveDatabaseRepo
	LikeSaveCacheRepo
}
