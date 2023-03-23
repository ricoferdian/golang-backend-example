package postgres

import (
	"github.com/jmoiron/sqlx"
	"kora-backend/internal/domain/auth"
)

type PostgresUserAuthRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresUserAuthRepository(dbCli *sqlx.DB) auth.UserAuthDatabaseRepo {
	return &PostgresUserAuthRepository{
		dbCli: dbCli,
	}
}
