package choreographer

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
)

type ChoreographerDatabaseRepo interface {
	GetChoreographerList(ctx context.Context) ([]model.ChoreographerModel, error)
	GetChoreographerById(ctx context.Context, choreographerID int64) (*model.ChoreographerModel, error)
	GetChoreographerByIdsMap(ctx context.Context, choreographerIDs []int64) (map[int64]model.ChoreographerModel, error)
}

type ChoreographerCacheRepo interface {
}

type ChoreographerRepository interface {
	ChoreographerDatabaseRepo
	ChoreographerCacheRepo
}
