package postgres

import (
	"github.com/jmoiron/sqlx"
	"kora-backend/internal/domain/choreo"
)

type PostgresChoreoRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresChoreoRepository(dbCli *sqlx.DB) choreo.ChoreoDatabaseRepo {
	return &PostgresChoreoRepository{
		dbCli: dbCli,
	}
}
