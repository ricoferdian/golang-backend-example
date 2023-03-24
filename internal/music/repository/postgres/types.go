package postgres

import (
	"github.com/jmoiron/sqlx"
	"kora-backend/internal/domain/music"
)

type PostgresMusicRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresMusicRepository(dbCli *sqlx.DB) music.MusicDatabaseRepo {
	return &PostgresMusicRepository{
		dbCli: dbCli,
	}
}
