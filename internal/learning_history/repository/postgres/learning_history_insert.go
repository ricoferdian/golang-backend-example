package postgres

import (
	"context"
	"database/sql"
	sq "github.com/huandu/go-sqlbuilder"
	"kora-backend/internal/entity"
	"kora-backend/internal/learning_history/helper"
	"kora-backend/internal/model"
)

func (c PostgresLearningHistoryRepository) InsertLearningHistory(ctx context.Context, history entity.SubmitLearningHistoryEntity) (*model.SubmitLearningHistoryModel, error) {
	query, args := c.buildInsertLearningHistory(history)
	historyId, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer historyId.Close()
	for historyId.Next() {
		err = historyId.Scan(
			&history.LearningHistoryID,
		)
		break
	}
	result := helper.SubmitLearningHistoryEntityToModel(history)

	return &result, nil
}

func (c PostgresLearningHistoryRepository) buildInsertLearningHistory(history entity.SubmitLearningHistoryEntity) (string, []interface{}) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto(tableLearningHistory)
	ib.Cols(columnInsertLearningHistory)
	ib.Values(
		history.ChoreoDetailID,
		history.UserID,
		history.ChoreoDetail,
		history.Device,
		history.Progress,
		history.RecordUrl,
		history.ThumbnailUrl,
	)
	addQ := sq.Buildf("%v RETURNING learning_history_id", ib)

	return addQ.Build()
}
