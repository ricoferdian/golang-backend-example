package usecase

import (
	"context"
	choreoHelper "github.com/Kora-Dance/koradance-backend/internal/choreo/helper"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/internal/purchase/helper"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"golang.org/x/sync/errgroup"
)

func (c ChoreoPurchaseUseCaseImpl) GetPurchasedChoreo(ctx context.Context, userID int64) (resultEntity []entity.ChoreoPurchaseEntity, err error) {
	purchasedList, choreoIDs, err := c.baseRepo.ChoreoPurchaseRepository().GetPurchasedChoreoByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Get detail data from choreo list retrieved using goroutine
	choreoMapCh := make(chan map[int64]model.ChoreographyModel, 1)
	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(choreoMapCh)
		choreoMap, err := c.baseRepo.ChoreoRepository().GetChoreoByIdsMap(errCtx, choreoIDs)
		if err != nil {
			return err
		}
		if len(choreoMap) == 0 {
			return nil
		}
		select {
		case choreoMapCh <- choreoMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return nil, err
	}

	choreoMap := <-choreoMapCh
	for _, purchaseData := range purchasedList {
		choreoPurchaseEntity := helper.ChoreoPurchaseModelToEntity(purchaseData)

		// Retrieve data from model
		if data, ok := choreoMap[choreoPurchaseEntity.ChoreoID]; ok {
			choreo := choreoHelper.ChoreoModelToEntity(data)
			choreoPurchaseEntity.Choreography = &choreo
		}
		resultEntity = append(resultEntity, choreoPurchaseEntity)

	}
	return resultEntity, nil
}
