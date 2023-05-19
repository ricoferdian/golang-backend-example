package postgres

import (
	"github.com/jmoiron/sqlx"
	"kora-backend/internal/domain/purchase"
)

type PostgresChoreoPurchaseRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresChoreoPurchaseRepository(dbCli *sqlx.DB) purchase.ChoreoPurchaseDatabaseRepo {
	return &PostgresChoreoPurchaseRepository{
		dbCli: dbCli,
	}
}
