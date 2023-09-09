package usecase

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func (m MusicUseCaseImpl) GetMusicByID(ctx context.Context, musicID int64) (entity.MusicEntity, error) {
	musicData, err := m.baseRepo.MusicRepository().GetMusicById(ctx, musicID)
	if musicData == nil || err != nil {
		return entity.MusicEntity{}, err
	}
	return helper.MusicModelToEntity(*musicData), nil
}

func (m MusicUseCaseImpl) GetAllMusic(ctx context.Context) ([]entity.MusicEntity, error) {
	musicList, err := m.baseRepo.MusicRepository().GetAllMusic(ctx)
	if err != nil {
		return nil, err
	}
	var result []entity.MusicEntity
	for _, musicData := range musicList {
		data := helper.MusicModelToEntity(musicData)
		result = append(result, data)
	}
	return result, nil
}

func (m MusicUseCaseImpl) UpsertMusic(ctx context.Context, music entity.MusicEntity) (entity.MusicEntity, error) {
	musicModel := helper.MusicEntityToModel(music)
	musicResult, err := m.baseRepo.MusicRepository().UpsertMusic(ctx, musicModel)
	if err != nil || musicResult == nil {
		return music, err
	}
	music = helper.MusicModelToEntity(*musicResult)
	return music, nil
}

func (m MusicUseCaseImpl) DeleteMusicByID(ctx context.Context, musicID int64) error {
	err := m.baseRepo.MusicRepository().DeleteMusicByID(ctx, musicID)
	if err != nil {
		return err
	}
	return nil
}
