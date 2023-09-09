package choreographer

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"io"
)

type ChoreographerDatabaseRepo interface {
	GetChoreographerList(ctx context.Context) ([]model.ChoreographerModel, error)
	GetChoreographerById(ctx context.Context, choreographerID int64) (*model.ChoreographerModel, error)
	GetChoreographerByIdsMap(ctx context.Context, choreographerIDs []int64) (map[int64]model.ChoreographerModel, error)
	UpsertChoreographerByIds(ctx context.Context, choreographerData model.ChoreographerModel) (*model.ChoreographerModel, error)
	DeleteChoreographerByID(ctx context.Context, choreographID int64) error
}

type ChoreographerCacheRepo interface {
}

type S3ChoreographerContentRepo interface {
	UploadChoreographerContent(choreographerID int64, fileName string, fileReader io.Reader) (string, error)
}

type ChoreographerRepository interface {
	ChoreographerDatabaseRepo
	ChoreographerCacheRepo
	S3ChoreographerContentRepo
}
