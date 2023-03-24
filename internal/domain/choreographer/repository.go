package choreographer

import (
	"context"
	"kora-backend/internal/model"
)

type ChoreographerDatabaseRepo interface {
	GetChoreographerById(ctx context.Context, choreographerID int64) (*model.ChoreographerModel, error)
	GetChoreographerByIdsMap(ctx context.Context, choreographerIDs []int64) (map[int64]model.ChoreographerModel, error)
}

type ChoreographerCacheRepo interface {
}

type ChoreographerRepository interface {
	ChoreographerDatabaseRepo
	ChoreographerCacheRepo
}
