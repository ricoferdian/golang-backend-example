package usecase

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"kora-backend/internal/choreo/helper"
	"kora-backend/internal/entity"
)

func (c ChoreoUseCaseImpl) GetChoreoDetailByChoreoID(ctx context.Context, filter entity.ChoreoDetailFilterEntity) (choreoResult []entity.ChoreographyDetailEntity, err error) {
	parentChoreo, err := c.baseRepo.ChoreoRepository().GetChoreoById(ctx, filter.ChoreoID)
	if err != nil || parentChoreo == nil {
		return choreoResult, err
	}
	isUnlocked := c.checkContentUnlockStatus(*parentChoreo)
	if !isUnlocked {
		return choreoResult, errors.New("choreo not unlocked yet")
	}

	choreoList, err := c.baseRepo.ChoreoRepository().GetChoreoDetailByChoreoID(ctx, filter)
	if err != nil {
		return choreoResult, err
	}
	for _, choreoData := range choreoList {
		choreoResult = append(choreoResult, helper.ChoreoDetailToEntity(choreoData))
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

	if !isFree && !isPurchased {
		return choreoResult, errors.New("choreo not unlocked yet")
	}

	choreoList, err := c.baseRepo.ChoreoRepository().GetChoreoDetailByChoreoID(ctx, filter)
	if err != nil {
		return choreoResult, err
	}
	for _, choreoData := range choreoList {
		choreoResult = append(choreoResult, helper.ChoreoDetailToEntity(choreoData))
	}
	return choreoResult, nil
}
