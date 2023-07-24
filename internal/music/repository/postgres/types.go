package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/Kora-Dance/koradance-backend/internal/domain/music"
)

type PostgresMusicRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresMusicRepository(dbCli *sqlx.DB) music.MusicDatabaseRepo {
	return &PostgresMusicRepository{
		dbCli: dbCli,
	}
}
