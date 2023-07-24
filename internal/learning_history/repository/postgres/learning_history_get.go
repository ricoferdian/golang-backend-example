package postgres

import (
	"context"
	"database/sql"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresLearningHistoryRepository) GetLearningHistoryList(ctx context.Context, filter model.LearningHistoryFilter) (result []model.LearningHistoryModel, err error) {
	query, args := c.buildGetLearningHistoryList(filter)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		choreoDetData, err := c.scanLearningHistoryData(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, choreoDetData)
	}
	return result, nil
}

func (c PostgresLearningHistoryRepository) buildGetLearningHistoryList(filter model.LearningHistoryFilter) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllLearningHistory)
	sb.From(tableLearningHistory)

	if filter.LearningHistoryID != 0 {
		sb.Where(sb.Equal("learning_history_id", filter.LearningHistoryID))
	}
	if filter.ChoreoDetailID != 0 {
		sb.Where(sb.Equal("det_choreo_id", filter.ChoreoDetailID))
	}
	if filter.UserID != 0 {
		sb.Where(sb.Equal("user_id", filter.UserID))
	}

	return sb.Build()
}

func (c PostgresLearningHistoryRepository) scanLearningHistoryData(row *sql.Rows) (result model.LearningHistoryModel, err error) {
	err = row.Scan(
		&result.LearningHistoryID,
		&result.ChoreoDetail,
		&result.Device,
		&result.Downloaded,
		&result.Expired,
		&result.ChoreoDetailID,
		&result.Progress,
		&result.RecordUrl,
		&result.ThumbnailUrl,
		&result.UserID,
	)

	return result, err
}
