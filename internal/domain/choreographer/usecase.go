package choreographer

import (
	"context"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
)

type ChoreographerUseCase interface {
	GetChoreographerList(ctx context.Context) ([]entity2.ChoreographerEntity, error)
	GetChoreographerByID(ctx context.Context, filter entity2.ChoreographerFilter) (*entity2.ChoreographerEntity, error)
	UpsertChoreographer(ctx context.Context, choreographerData entity2.ChoreographerEntity) (entity2.ChoreographerEntity, error)
	DeleteChoreographerByID(ctx context.Context, musicID int64) error
}
