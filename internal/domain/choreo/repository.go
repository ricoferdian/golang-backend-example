package choreo

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"io"
)

type ChoreoDatabaseRepo interface {
	GetChoreoById(ctx context.Context, choreoID int64) (*model.ChoreographyModel, error)
	GetChoreoList(ctx context.Context, filter entity.ChoreoFilterEntity) ([]model.ChoreographyModel, error)
	GetChoreoByIdsMap(ctx context.Context, choreoIDs []int64) (map[int64]model.ChoreographyModel, error)
	// GetChoreoListWithMusicAndChoreographIds used to retrieve choreo list with list of music and choreographer id
	//
	// the list are unique
	GetChoreoDetailById(ctx context.Context, choreoDetailID int64) (*model.ChoreographyDetailModel, error)
	GetChoreoListWithMusicAndChoreographIds(ctx context.Context, filter entity.ChoreoFilterEntity) ([]model.ChoreographyModel, []int64, []int64, []int64, error)
	GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) ([]model.ChoreographyDetailModel, error)
	// Insert
	InsertChoreo(ctx context.Context, choreo model.ChoreographyModel) (result model.ChoreographyModel, err error)
	InsertChoreoDetail(ctx context.Context, detail model.ChoreographyDetailModel) (result model.ChoreographyDetailModel, err error)
	// Update link
	UpdateChoreoLink(ctx context.Context, choreo model.ChoreographyModel) (result model.ChoreographyModel, err error)
	UpdateChoreoDetailLink(ctx context.Context, detail model.ChoreographyDetailModel) (result model.ChoreographyDetailModel, err error)

	UpdateChoreo(ctx context.Context, choreo model.ChoreographyModel) (result model.ChoreographyModel, err error)
	UpdateChoreoDetail(ctx context.Context, choreo model.ChoreographyDetailModel) (result model.ChoreographyDetailModel, err error)
}

type S3ChoreoContentRepo interface {
	UploadChoreoContent(choreoID int64, fileName string, fileReader io.Reader) (string, error)
	UploadChoreoDetailContent(choreoDetailID int64, fileName string, fileReader io.Reader) (string, error)
}

type ChoreoCacheRepo interface {
}

type ChoreoRepository interface {
	ChoreoDatabaseRepo
	ChoreoCacheRepo
	S3ChoreoContentRepo
}
