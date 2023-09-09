package postgres

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresLearningHistoryRepository) InsertLearningHistory(ctx context.Context, history entity.SubmitLearningHistoryEntity) (*model.SubmitLearningHistoryModel, error) {
	query, args := c.buildInsertLearningHistory(history)
	historyId, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
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
