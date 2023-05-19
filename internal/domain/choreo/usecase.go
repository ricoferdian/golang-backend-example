package choreo

import (
	"context"
	"kora-backend/internal/entity"
)

type ChoreoUseCase interface {
	GetChoreoList(ctx context.Context) ([]entity.ChoreographyEntity, error)
	GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) ([]entity.ChoreographyDetailEntity, error)
	GetChoreoDetailByChoreoIDWithUserContent(ctx context.Context, userID int64, filter entity.ChoreoDetailFilterEntity) (choreoResult []entity.ChoreographyDetailEntity, err error)
	GetChoreoListWithUserContent(ctx context.Context, userID int64) ([]entity.ChoreographyEntity, error)
}
