package postgres

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/like_save"
	"github.com/jmoiron/sqlx"
)

type PostgresLikeSaveRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresLikeSaveRepository(dbCli *sqlx.DB) like_save.LikeSaveDatabaseRepo {
	return &PostgresLikeSaveRepository{
		dbCli: dbCli,
	}
}
