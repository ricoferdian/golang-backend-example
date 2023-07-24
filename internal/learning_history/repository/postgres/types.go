package postgres

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/learning_history"
	"github.com/jmoiron/sqlx"
)

type PostgresLearningHistoryRepository struct {
	dbCli *sqlx.DB
}

func NewPostgresLearningHistoryRepository(dbCli *sqlx.DB) learning_history.LearningHistoryDatabaseRepo {
	return &PostgresLearningHistoryRepository{
		dbCli: dbCli,
	}
}
