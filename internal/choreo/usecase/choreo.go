package usecase

import (
	"context"
	"golang.org/x/sync/errgroup"
	"kora-backend/internal/choreo/helper"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
)

func (c ChoreoUseCaseImpl) GetChoreoList(ctx context.Context) ([]entity.ChoreographyEntity, error) {
	choreoList, musicIds, cgpherIds, err := c.baseRepo.ChoreoRepository().GetChoreoListWithMusicAndChoreographIds(ctx)
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
