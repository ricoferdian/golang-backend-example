package choreo

import (
	"context"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
	"io"
)

type ChoreoUseCase interface {
	GetChoreoByID(ctx context.Context, choreoID int64) (choreoResult entity2.ChoreographyEntity, err error)
	GetChoreoByIDWithUserContent(ctx context.Context, userID int64, choreoID int64) (choreoResult entity2.ChoreographyEntity, err error)
	GetChoreoList(ctx context.Context, filter entity2.ChoreoFilterEntity) ([]entity2.ChoreographyEntity, error)
	GetChoreoDetailByChoreoID(ctx context.Context, filter entity2.ChoreoDetailFilterEntity) ([]entity2.ChoreographyDetailEntity, error)
	GetChoreoDetailByChoreoIDWithUserContent(ctx context.Context, userID int64, filter entity2.ChoreoDetailFilterEntity) (choreoResult []entity2.ChoreographyDetailEntity, err error)
	GetChoreoListWithUserContent(ctx context.Context, userID int64, filter entity2.ChoreoFilterEntity) ([]entity2.ChoreographyEntity, error)
	UploadChoreoDetailContent(ctx context.Context, choreoDetailID int64, fileCategory int, fileName string, fileReader io.Reader) (interface{}, error)
	UploadChoreoContent(ctx context.Context, choreoID int64, fileCategory int, fileName string, fileReader io.Reader) (interface{}, error)
	InsertChoreo(ctx context.Context, choreo entity2.ChoreographyEntity) (result entity2.ChoreographyEntity, err error)
	InsertChoreoDetail(ctx context.Context, detail entity2.ChoreographyDetailEntity) (result entity2.ChoreographyDetailEntity, err error)
	UpdateChoreo(ctx context.Context, choreo entity2.ChoreographyEntity) (result entity2.ChoreographyEntity, err error)
	UpdateChoreoDetail(ctx context.Context, detail entity2.ChoreographyDetailEntity) (result entity2.ChoreographyDetailEntity, err error)
}
