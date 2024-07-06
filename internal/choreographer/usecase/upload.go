package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Kora-Dance/koradance-backend/internal/common/general"
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	"io"
)

func (c ChoreographerUseCaseImpl) UploadChoreographerContent(ctx context.Context, choreographerID int64, fileCategory int, originalFileName string, fileReader io.Reader) (result interface{}, err error) {
	choreographerData, err := c.baseRepo.ChoreographerRepository().GetChoreographerById(ctx, choreographerID)
	if err != nil {
		return result, err
	}
	if choreographerData == nil {
		return result, errors.New("choreographer id not found")
	}
	format, err := general.GetFileFormat(originalFileName)
	if err != nil {
		return result, err
	}
	fName := fmt.Sprintf("%s.%s", general.GenerateUniqueFileName(), format)
	url, err := c.baseRepo.ChoreographerRepository().UploadChoreographerContent(choreographerID, fName, fileReader)
	if err != nil {
		return result, err
	}
	choreographerData.ProfileImageURL = sql.NullString{
		String: url,
		Valid:  true,
	}
	_, err = c.baseRepo.ChoreographerRepository().UpsertChoreographerByIds(ctx, *choreographerData)
	if err != nil {
		return result, err
	}
	return helper.ChoreographerModelToEntity(*choreographerData), nil
}
