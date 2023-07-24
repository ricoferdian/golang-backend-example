package postgres

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/choreo"
	"github.com/jmoiron/sqlx"
)

type PostgresChoreoRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresChoreoRepository(dbCli *sqlx.DB) choreo.ChoreoDatabaseRepo {
	return &PostgresChoreoRepository{
		dbCli: dbCli,
	}
}
