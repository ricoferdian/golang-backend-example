package postgres

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreographer"
	"github.com/jmoiron/sqlx"
)

type PostgresChoreographerRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresChoreographerRepository(dbCli *sqlx.DB) choreographer.ChoreographerDatabaseRepo {
	return &PostgresChoreographerRepository{
		dbCli: dbCli,
	}
}
