package choreo

import (
	"context"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

type ChoreoDatabaseRepo interface {
	GetChoreoList(ctx context.Context) ([]model.ChoreographyModel, error)
	GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) ([]model.ChoreographyDetailModel, error)
}

type ChoreoCacheRepo interface {
}

type ChoreoRepository interface {
	ChoreoDatabaseRepo
	ChoreoCacheRepo
}
