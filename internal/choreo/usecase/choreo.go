package usecase

import (
	"context"
	"errors"
	"github.com/Kora-Dance/koradance-backend/internal/choreo/helper"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"golang.org/x/sync/errgroup"
	"log"
	"strconv"
)

func (c ChoreoUseCaseImpl) GetChoreoByID(ctx context.Context, choreoID int64) (choreoResult entity.ChoreographyEntity, err error) {
	parentChoreo, err := c.baseRepo.ChoreoRepository().GetChoreoById(ctx, choreoID)
	if err != nil || parentChoreo == nil {
		return choreoResult, err
	}

	musicIds := []int64{parentChoreo.MusicID.Int64}
	cgpherIds := []int64{parentChoreo.ChoreographerID.Int64}
	// Get detail data from choreo list retrieved using goroutine
	musicMapCh := make(chan map[int64]model.MusicModel, 1)
	cgpherMapCh := make(chan map[int64]model.ChoreographerModel, 1)
	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(musicMapCh)
		musicMap, err := c.baseRepo.MusicRepository().GetMusicByIdsMap(errCtx, musicIds)
		if err != nil {
			return err
		}
		if len(musicMap) == 0 {
			return nil
		}
		select {
		case musicMapCh <- musicMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(cgpherMapCh)
		cgpherMap, err := c.baseRepo.ChoreographerRepository().GetChoreographerByIdsMap(errCtx, cgpherIds)
		if err != nil {
			return err
		}
		if len(cgpherMap) == 0 {
			return nil
		}
		select {
		case cgpherMapCh <- cgpherMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return choreoResult, err
	}

	cgpherMap := <-cgpherMapCh
	musicMap := <-musicMapCh

	// Convert to entity
	choreoEntity := helper.ChoreoModelToEntity(*parentChoreo)
	choreoEntity.Unlocked = c.checkContentUnlockStatus(*parentChoreo)
	choreoEntity.CurrentPrice = strconv.FormatInt(parentChoreo.TempPrice.Int64, 10)

	// Retrieve data from model
	if data, ok := cgpherMap[choreoEntity.ChoreographerID]; ok {
		cgpher := helper.ChoreographerModelToEntity(data)
		choreoEntity.ChoreographerData = &cgpher
	}
	if data, ok := musicMap[choreoEntity.MusicID]; ok {
		music := helper.MusicModelToEntity(data)
		choreoEntity.MusicData = &music
	}

	return choreoEntity, nil
}

// GetChoreoByIDWithUserContent is used to get choreo detail with user content
func (c ChoreoUseCaseImpl) GetChoreoByIDWithUserContent(ctx context.Context, userID int64, choreoID int64) (choreoResult entity.ChoreographyEntity, err error) {
	parentChoreo, err := c.baseRepo.ChoreoRepository().GetChoreoById(ctx, choreoID)
	if err != nil || parentChoreo == nil {
		return choreoResult, err
	}

	musicIds := []int64{parentChoreo.MusicID.Int64}
	cgpherIds := []int64{parentChoreo.ChoreographerID.Int64}
	choreoIds := []int64{parentChoreo.ChoreoID}

	// Get detail data from choreo list retrieved using goroutine
	musicMapCh := make(chan map[int64]model.MusicModel, 1)
	cgpherMapCh := make(chan map[int64]model.ChoreographerModel, 1)
	purchaseMapCh := make(chan map[int64]model.ChoreoPurchaseModel, 1)
	likeMapCh := make(chan map[int64]model.ChoreoLikeModel, 1)
	saveMapCh := make(chan map[int64]model.ChoreoSaveModel, 1)
	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(musicMapCh)
		musicMap, err := c.baseRepo.MusicRepository().GetMusicByIdsMap(errCtx, musicIds)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get music", err)
			return nil
		}
		if len(musicMap) == 0 {
			return nil
		}
		select {
		case musicMapCh <- musicMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(cgpherMapCh)
		cgpherMap, err := c.baseRepo.ChoreographerRepository().GetChoreographerByIdsMap(errCtx, cgpherIds)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get choreographer", err)
			return nil
		}
		if len(cgpherMap) == 0 {
			return nil
		}
		select {
		case cgpherMapCh <- cgpherMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(purchaseMapCh)
		purchaseMap, err := c.baseRepo.ChoreoPurchaseRepository().GetPurchasedChoreoByUserIDMap(errCtx, userID)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get purchased choreo", err)
			return nil
		}
		if len(purchaseMap) == 0 {
			return nil
		}
		select {
		case purchaseMapCh <- purchaseMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(saveMapCh)
		savedMap, err := c.baseRepo.ChoreoLikeSaveRepository().GetSavedChoreoByChoreoIDsMap(errCtx, userID, choreoIds)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get saved choreo", err)
			return nil
		}
		if len(savedMap) == 0 {
			return nil
		}
		select {
		case saveMapCh <- savedMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(likeMapCh)
		likeMap, err := c.baseRepo.ChoreoLikeSaveRepository().GetLikedChoreoByChoreoIDsMap(errCtx, userID, choreoIds)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get liked choreo", err)
			return nil
		}
		if len(likeMap) == 0 {
			return nil
		}
		select {
		case likeMapCh <- likeMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return choreoResult, err
	}

	cgpherMap := <-cgpherMapCh
	musicMap := <-musicMapCh
	purchaseMap := <-purchaseMapCh
	likeMap := <-likeMapCh
	saveMap := <-saveMapCh

	choreoEntity := helper.ChoreoModelToEntity(*parentChoreo)
	choreoEntity.Unlocked = c.checkContentUnlockStatus(*parentChoreo)
	choreoEntity.CurrentPrice = strconv.FormatInt(parentChoreo.TempPrice.Int64, 10)

	// Retrieve data from model
	if data, ok := cgpherMap[choreoEntity.ChoreographerID]; ok {
		cgpher := helper.ChoreographerModelToEntity(data)
		choreoEntity.ChoreographerData = &cgpher
	}
	if data, ok := musicMap[choreoEntity.MusicID]; ok {
		music := helper.MusicModelToEntity(data)
		choreoEntity.MusicData = &music
	}
	// If purchased, unlock content
	if _, ok := purchaseMap[choreoEntity.ChoreoID]; ok {
		choreoEntity.Unlocked = true
	}
	// like and save
	if _, ok := likeMap[choreoEntity.ChoreoID]; ok {
		choreoEntity.Like = true
	}
	if _, ok := saveMap[choreoEntity.ChoreoID]; ok {
		choreoEntity.Save = true
	}

	return choreoEntity, nil
}

func (c ChoreoUseCaseImpl) GetChoreoList(ctx context.Context, filter entity.ChoreoFilterEntity) ([]entity.ChoreographyEntity, error) {
	choreoList, _, musicIds, cgpherIds, err := c.baseRepo.ChoreoRepository().GetChoreoListWithMusicAndChoreographIds(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Get detail data from choreo list retrieved using goroutine
	musicMapCh := make(chan map[int64]model.MusicModel, 1)
	cgpherMapCh := make(chan map[int64]model.ChoreographerModel, 1)
	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(musicMapCh)
		musicMap, err := c.baseRepo.MusicRepository().GetMusicByIdsMap(errCtx, musicIds)
		if err != nil {
			return err
		}
		if len(musicMap) == 0 {
			return nil
		}
		select {
		case musicMapCh <- musicMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(cgpherMapCh)
		cgpherMap, err := c.baseRepo.ChoreographerRepository().GetChoreographerByIdsMap(errCtx, cgpherIds)
		if err != nil {
			return err
		}
		if len(cgpherMap) == 0 {
			return nil
		}
		select {
		case cgpherMapCh <- cgpherMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return nil, err
	}

	cgpherMap := <-cgpherMapCh
	musicMap := <-musicMapCh

	// Construct response
	var choreoResult []entity.ChoreographyEntity
	for _, choreoData := range choreoList {
		// Convert to entity
		choreoEntity := helper.ChoreoModelToEntity(choreoData)
		choreoEntity.Unlocked = c.checkContentUnlockStatus(choreoData)
		choreoEntity.CurrentPrice = strconv.FormatInt(choreoData.TempPrice.Int64, 10)

		// Retrieve data from model
		if data, ok := cgpherMap[choreoEntity.ChoreographerID]; ok {
			cgpher := helper.ChoreographerModelToEntity(data)
			choreoEntity.ChoreographerData = &cgpher
		}
		if data, ok := musicMap[choreoEntity.MusicID]; ok {
			music := helper.MusicModelToEntity(data)
			choreoEntity.MusicData = &music
		}

		// Append to result
		choreoResult = append(choreoResult, choreoEntity)
	}
	return choreoResult, nil
}

func (c ChoreoUseCaseImpl) GetChoreoListWithUserContent(ctx context.Context, userID int64, filter entity.ChoreoFilterEntity) ([]entity.ChoreographyEntity, error) {
	choreoList, choreoIds, musicIds, cgpherIds, err := c.baseRepo.ChoreoRepository().GetChoreoListWithMusicAndChoreographIds(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Get detail data from choreo list retrieved using goroutine
	musicMapCh := make(chan map[int64]model.MusicModel, 1)
	cgpherMapCh := make(chan map[int64]model.ChoreographerModel, 1)
	purchaseMapCh := make(chan map[int64]model.ChoreoPurchaseModel, 1)
	likeMapCh := make(chan map[int64]model.ChoreoLikeModel, 1)
	saveMapCh := make(chan map[int64]model.ChoreoSaveModel, 1)
	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(musicMapCh)
		musicMap, err := c.baseRepo.MusicRepository().GetMusicByIdsMap(errCtx, musicIds)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get music", err)
			return nil
		}
		if len(musicMap) == 0 {
			return nil
		}
		select {
		case musicMapCh <- musicMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(cgpherMapCh)
		cgpherMap, err := c.baseRepo.ChoreographerRepository().GetChoreographerByIdsMap(errCtx, cgpherIds)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get choreographer", err)
			return nil
		}
		if len(cgpherMap) == 0 {
			return nil
		}
		select {
		case cgpherMapCh <- cgpherMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(purchaseMapCh)
		purchaseMap, err := c.baseRepo.ChoreoPurchaseRepository().GetPurchasedChoreoByUserIDMap(errCtx, userID)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get purchased choreo", err)
			return nil
		}
		if len(purchaseMap) == 0 {
			return nil
		}
		select {
		case purchaseMapCh <- purchaseMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(saveMapCh)
		savedMap, err := c.baseRepo.ChoreoLikeSaveRepository().GetSavedChoreoByChoreoIDsMap(errCtx, userID, choreoIds)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get saved choreo", err)
			return nil
		}
		if len(savedMap) == 0 {
			return nil
		}
		select {
		case saveMapCh <- savedMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	g.Go(func() error {
		defer close(likeMapCh)
		likeMap, err := c.baseRepo.ChoreoLikeSaveRepository().GetLikedChoreoByChoreoIDsMap(errCtx, userID, choreoIds)
		if err != nil {
			log.Println("[ChoreoUseCaseImpl][Error] err get liked choreo", err)
			return nil
		}
		if len(likeMap) == 0 {
			return nil
		}
		select {
		case likeMapCh <- likeMap:
		case <-errCtx.Done():
			return errCtx.Err()
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return nil, err
	}

	cgpherMap := <-cgpherMapCh
	musicMap := <-musicMapCh
	purchaseMap := <-purchaseMapCh
	likeMap := <-likeMapCh
	saveMap := <-saveMapCh

	// Construct response
	var choreoResult []entity.ChoreographyEntity
	for _, choreoData := range choreoList {
		// Convert to entity
		choreoEntity := helper.ChoreoModelToEntity(choreoData)
		choreoEntity.Unlocked = c.checkContentUnlockStatus(choreoData)
		choreoEntity.CurrentPrice = strconv.FormatInt(choreoData.TempPrice.Int64, 10)

		// Retrieve data from model
		if data, ok := cgpherMap[choreoEntity.ChoreographerID]; ok {
			cgpher := helper.ChoreographerModelToEntity(data)
			choreoEntity.ChoreographerData = &cgpher
		}
		if data, ok := musicMap[choreoEntity.MusicID]; ok {
			music := helper.MusicModelToEntity(data)
			choreoEntity.MusicData = &music
		}
		// If purchased, unlock content
		if _, ok := purchaseMap[choreoEntity.ChoreoID]; ok {
			choreoEntity.Unlocked = true
		}
		// like and save
		if _, ok := likeMap[choreoEntity.ChoreoID]; ok {
			choreoEntity.Like = true
		}
		if _, ok := saveMap[choreoEntity.ChoreoID]; ok {
			choreoEntity.Save = true
		}

		// Append to result
		choreoResult = append(choreoResult, choreoEntity)
	}
	return choreoResult, nil
}

func (c ChoreoUseCaseImpl) InsertChoreo(ctx context.Context, choreo entity.ChoreographyEntity) (result entity.ChoreographyEntity, err error) {
	choreoData := helper.ChoreoEntityToModel(choreo)
	data, err := c.baseRepo.ChoreoRepository().InsertChoreo(ctx, choreoData)
	if err != nil {
		return result, err
	}
	return helper.ChoreoModelToEntity(data), nil
}

func (c ChoreoUseCaseImpl) UpdateChoreo(ctx context.Context, choreo entity.ChoreographyEntity) (result entity.ChoreographyEntity, err error) {
	if choreo.ChoreoID == 0 {
		return result, errors.New("choreo id cannot be empty")
	}
	choreoData := helper.ChoreoEntityToModel(choreo)
	_, err = c.baseRepo.ChoreoRepository().UpdateChoreo(ctx, choreoData)
	if err != nil {
		return result, err
	}
	data, err := c.baseRepo.ChoreoRepository().GetChoreoById(ctx, choreo.ChoreoID)
	if err != nil {
		return result, err
	}
	return helper.ChoreoModelToEntity(*data), nil
}
