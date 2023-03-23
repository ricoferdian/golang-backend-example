package postgres

import (
	"context"
	"database/sql"
	sq "github.com/huandu/go-sqlbuilder"
	"kora-backend/internal/model"
)

func (c PostgresChoreoRepository) GetChoreoList(ctx context.Context) (result []model.ChoreographyModel, err error) {
	query, args := c.buildGetChoreoList()
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return result, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		choreoData, err := c.scanChoreoData(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, choreoData)
	}
	return result, nil
}

func (c PostgresChoreoRepository) buildGetChoreoList() (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllChoreo)
	sb.From(tableMasterChoreo)

	return sb.Build()
}

func (c PostgresChoreoRepository) scanChoreoData(row *sql.Rows) (result model.ChoreographyModel, err error) {
	err = row.Scan(
		&result.ChoreoID,
		&result.Title,
		&result.Description,
		&result.Difficulty,
		&result.Duration,
		&result.IsActive,
		&result.Position,
		&result.VideoPreviewURL,
		&result.VideoThumbnailURL,
		&result.ChoreographerID,
		&result.MusicID,
	)

	return result, err
}
