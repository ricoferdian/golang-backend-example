package usecase

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"golang.org/x/sync/errgroup"
	"log"
)

func (c ChoreoUseCaseImpl) GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) (choreoResult []entity.ChoreographyDetailEntity, err error) {
	parentChoreo, err := c.baseRepo.ChoreoRepository().GetChoreoById(ctx, filter.ChoreoID)
	if err != nil || parentChoreo == nil {
		return choreoResult, err
	}
	isUnlocked := c.checkContentUnlockStatus(*parentChoreo)

	choreoList, err := c.baseRepo.ChoreoRepository().GetChoreoDetailByChoreoID(ctx, filter)
	if err != nil {
		return choreoResult, err
	}
	for _, choreoData := range choreoList {
		data := helper.ChoreoDetailToEntity(choreoData)
		if !isUnlocked {
			data.VideoURL = ""
		}
		choreoResult = append(choreoResult, data)
	}
	return choreoResult, nil
}

func (c ChoreoUseCaseImpl) GetChoreoDetailByChoreoIDWithUserContent(ctx context.Context, userID int64, filter entity.ChoreoDetailFilterEntity) (choreoResult []entity.ChoreographyDetailEntity, err error) {
	// Payment verification
	isFreeCh := make(chan bool, 1)
	isPurchasedCh := make(chan bool, 1)
	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(isFreeCh)
		parentChoreo, err := c.baseRepo.ChoreoRepository().GetChoreoById(errCtx, filter.ChoreoID)
		if err != nil {
			return err
		}
		if parentChoreo == nil {
			return nil
		}
		select {
		case isFreeCh <- c.checkContentUnlockStatus(*parentChoreo):
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(isPurchasedCh)
		purchasedChoreo, err := c.baseRepo.ChoreoPurchaseRepository().GetPurchasedChoreoByID(errCtx, userID, filter.ChoreoID)
		if err != nil {
			return err
		}
		if purchasedChoreo == nil {
			return nil
		}
		select {
		case isPurchasedCh <- purchasedChoreo.ChoreoID != 0:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return nil, err
	}

	isFree := <-isFreeCh
	isPurchased := <-isPurchasedCh

	choreoList, err := c.baseRepo.ChoreoRepository().GetChoreoDetailByChoreoID(ctx, filter)
	if err != nil {
		return choreoResult, err
	}
	for _, choreoData := range choreoList {
		data := helper.ChoreoDetailToEntity(choreoData)
		if !isFree && !isPurchased {
			data.VideoURL = ""
		}
		choreoResult = append(choreoResult, data)
	}
	return choreoResult, nil
}

func (c ChoreoUseCaseImpl) InsertChoreoDetail(ctx context.Context, detail entity.ChoreographyDetailEntity) (result entity.ChoreographyDetailEntity, err error) {
	choreoDetail := helper.ChoreoDetailEntityToModel(detail)
	data, err := c.baseRepo.ChoreoRepository().InsertChoreoDetail(ctx, choreoDetail)
	if err != nil {
		return result, err
	}
	return helper.ChoreoDetailToEntity(data), nil
}

func (c ChoreoUseCaseImpl) UpdateChoreoDetail(ctx context.Context, detail entity.ChoreographyDetailEntity) (result entity.ChoreographyDetailEntity, err error) {
	choreoDetail := helper.ChoreoDetailEntityToModel(detail)
	_, err = c.baseRepo.ChoreoRepository().UpdateChoreoDetail(ctx, choreoDetail)
	if err != nil {
		return result, err
	}
	data, err := c.baseRepo.ChoreoRepository().GetChoreoDetailById(ctx, detail.ChoreoDetailID)
	if err != nil {
		return result, err
	}
	if data == nil {
		log.Print("[ChoreoUseCaseImpl] nil choreo detail id")
		return result, nil
	}
	return helper.ChoreoDetailToEntity(*data), nil
}

func (c ChoreoUseCaseImpl) DeleteChoreoDetailByID(ctx context.Context, choreoDetailID int64) error {
	return c.baseRepo.ChoreoRepository().DeleteChoreoDetailByID(ctx, choreoDetailID)
}
