package postgres

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/purchase"
	"github.com/jmoiron/sqlx"
)

type PostgresChoreoPurchaseRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresChoreoPurchaseRepository(dbCli *sqlx.DB) purchase.ChoreoPurchaseDatabaseRepo {
	return &PostgresChoreoPurchaseRepository{
		dbCli: dbCli,
	}
}
