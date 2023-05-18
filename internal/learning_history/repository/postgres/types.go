package postgres

import (
	"github.com/jmoiron/sqlx"
	"kora-backend/internal/domain/learning_history"
)

type PostgresLearningHistoryRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresLearningHistoryRepository(dbCli *sqlx.DB) learning_history.LearningHistoryDatabaseRepo {
	return &PostgresLearningHistoryRepository{
		dbCli: dbCli,
	}
}
