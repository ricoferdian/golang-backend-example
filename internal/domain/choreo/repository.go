package choreo

import (
	"context"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

type ChoreoDatabaseRepo interface {
	GetChoreoById(ctx context.Context, choreoID int64) (*model.ChoreographyModel, error)
	GetChoreoList(ctx context.Context) ([]model.ChoreographyModel, error)
	GetChoreoByIdsMap(ctx context.Context, choreoIDs []int64) (map[int64]model.ChoreographyModel, error)
	// GetChoreoListWithMusicAndChoreographIds used to retrieve choreo list with list of music and choreographer id
	//
	// the list are unique
	GetChoreoListWithMusicAndChoreographIds(ctx context.Context) ([]model.ChoreographyModel, []int64, []int64, []int64, error)
	GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) ([]model.ChoreographyDetailModel, error)
}

type ChoreoCacheRepo interface {
}

type ChoreoRepository interface {
	ChoreoDatabaseRepo
	ChoreoCacheRepo
}
