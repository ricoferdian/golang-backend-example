package choreo

import (
	"context"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

type ChoreoDatabaseRepo interface {
	GetChoreoList(ctx context.Context) ([]model.ChoreographyModel, error)
	// GetChoreoListWithMusicAndChoreographIds used to retrieve choreo list with list of music and choreographer id
	//
	// the list are unique
	GetChoreoListWithMusicAndChoreographIds(ctx context.Context) ([]model.ChoreographyModel, []int64, []int64, error)
	GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) ([]model.ChoreographyDetailModel, error)
}

type ChoreoCacheRepo interface {
}

type ChoreoRepository interface {
	ChoreoDatabaseRepo
	ChoreoCacheRepo
}
