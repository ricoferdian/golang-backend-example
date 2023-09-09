package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/common/general"
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"io"
)

func wrapLinkChoreoDetail(detail *model.ChoreographyDetailModel, fileCategory int, url string) model.ChoreographyDetailModel {
	if fileCategory == constants.FileCategoryTestVideo {
		detail.CDNTestVideoURL.String = url
		detail.CDNTestVideoURL.Valid = true
		return *detail
	}
	if fileCategory == constants.FileCategoryVideo {
		detail.CDNVideoURL.String = url
		detail.CDNVideoURL.Valid = true
		return *detail
	}
	detail.CDNVideoThumbnailURL.String = url
	detail.CDNVideoThumbnailURL.Valid = true
	return *detail
}

func (c ChoreoUseCaseImpl) UploadChoreoDetailContent(ctx context.Context, choreoDetailID int64, fileCategory int, originalFileName string, fileReader io.Reader) (result interface{}, err error) {
	data, err := c.baseRepo.ChoreoRepository().GetChoreoDetailById(ctx, choreoDetailID)
	if err != nil {
		return result, err
	}
	if data == nil {
		return result, errors.New("choreo detail id not found")
	}
	format, err := general.GetFileFormat(originalFileName)
	if err != nil {
		return result, err
	}
	fName := fmt.Sprintf("%s.%s", general.GenerateUniqueFileName(), format)
	url, err := c.baseRepo.ChoreoRepository().UploadChoreoDetailContent(choreoDetailID, fName, fileReader)
	if err != nil {
		return result, err
	}
	_, err = c.baseRepo.ChoreoRepository().UpdateChoreoDetailLink(ctx, wrapLinkChoreoDetail(data, fileCategory, url))
	if err != nil {
		return result, err
	}

	return helper.ChoreoDetailToEntity(*data), nil
}

func wrapLinkChoreo(data *model.ChoreographyModel, fileCategory int, url string) model.ChoreographyModel {
	if fileCategory == constants.FileCategoryVideo {
		data.CDNVideoPreviewURL.String = url
		data.CDNVideoPreviewURL.Valid = true
		return *data
	}
	data.CDNVideoThumbnailURL.String = url
	data.CDNVideoThumbnailURL.Valid = true
	return *data
}

func (c ChoreoUseCaseImpl) UploadChoreoContent(ctx context.Context, choreoID int64, fileCategory int, originalFileName string, fileReader io.Reader) (result interface{}, err error) {
	data, err := c.baseRepo.ChoreoRepository().GetChoreoById(ctx, choreoID)
	if err != nil {
		return result, err
	}
	if data == nil {
		return result, errors.New("choreo id not found")
	}
	format, err := general.GetFileFormat(originalFileName)
	if err != nil {
		return result, err
	}
	fName := fmt.Sprintf("%s.%s", general.GenerateUniqueFileName(), format)
	url, err := c.baseRepo.ChoreoRepository().UploadChoreoContent(choreoID, fName, fileReader)
	if err != nil {
		return result, err
	}
	_, err = c.baseRepo.ChoreoRepository().UpdateChoreoLink(ctx, wrapLinkChoreo(data, fileCategory, url))
	if err != nil {
		return result, err
	}
	return helper.ChoreoModelToEntity(*data), nil
}
