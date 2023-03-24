package choreo

import (
	"context"
	"kora-backend/internal/entity"
)

type ChoreoUseCase interface {
	GetChoreoList(ctx context.Context) ([]*entity.ChoreographyEntity, error)
	GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) ([]entity.ChoreographyDetailEntity, error)
}
