package postgres

import (
	"github.com/jmoiron/sqlx"
	"kora-backend/internal/domain/authdomain"
)

type PostgresUserAuthRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresUserAuthRepository(dbCli *sqlx.DB) authdomain.UserAuthDatabaseRepo {
	return PostgresUserAuthRepository{
		dbCli: dbCli,
	}
}
