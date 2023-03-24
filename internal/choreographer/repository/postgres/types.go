package postgres

import (
	"github.com/jmoiron/sqlx"
	"kora-backend/internal/domain/choreographer"
)

type PostgresChoreographerRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresChoreographerRepository(dbCli *sqlx.DB) choreographer.ChoreographerDatabaseRepo {
	return &PostgresChoreographerRepository{
		dbCli: dbCli,
	}
}
