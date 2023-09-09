package helper

import (
	"github.com/Kora-Dance/koradance-backend/internal/model"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
)

func MusicModelToEntity(musicModel model.MusicModel) entity2.MusicEntity {
	return entity2.MusicEntity{
		MusicID:    musicModel.MusicID,
		ArtistName: musicModel.ArtistName,
		Title:      musicModel.Title,
	}
}

func MusicEntityToModel(musicModel entity2.MusicEntity) model.MusicModel {
	return model.MusicModel{
		MusicID:    musicModel.MusicID,
		ArtistName: musicModel.ArtistName,
		Title:      musicModel.Title,
	}
}
