package postgres

import (
	"context"
	"database/sql"
	sq "github.com/huandu/go-sqlbuilder"
	"kora-backend/internal/model"
)

func (c PostgresChoreoRepository) getChoreoListRows(ctx context.Context) (rows *sql.Rows, err error) {
	query, args := c.buildGetChoreoList()
	rows, err = c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (c PostgresChoreoRepository) GetChoreoById(ctx context.Context, choreoID int64) (result *model.ChoreographyModel, err error) {
	query, args := c.buildGetChoreoByID(choreoID)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
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
		return &choreoData, nil
	}
	return nil, nil
}

func (c PostgresChoreoRepository) GetChoreoList(ctx context.Context) (result []model.ChoreographyModel, err error) {
	rows, err := c.getChoreoListRows(ctx)
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

func (c PostgresChoreoRepository) GetChoreoListWithMusicAndChoreographIds(ctx context.Context) (result []model.ChoreographyModel, choreoIds []int64, musicIds []int64, choreographerIds []int64, err error) {
	rows, err := c.getChoreoListRows(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	defer rows.Close()

	var musicIdExist = make(map[int64]bool)
	var cgpherIdExist = make(map[int64]bool)
	for rows.Next() {
		choreoData, musicId, cgpherId, err := c.scanChoreoDataWithRelatedIds(rows)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		if musicId != 0 && !musicIdExist[musicId] {
			musicIds = append(musicIds, musicId)
			musicIdExist[musicId] = true
		}
		if cgpherId != 0 && !cgpherIdExist[cgpherId] {
			choreographerIds = append(choreographerIds, cgpherId)
			cgpherIdExist[cgpherId] = true
		}
		choreoIds = append(choreoIds, choreoData.ChoreoID)
		result = append(result, choreoData)
	}
	return result, choreoIds, musicIds, choreographerIds, nil
}

func (c PostgresChoreoRepository) buildGetChoreoList() (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllChoreo)
	sb.From(tableMasterChoreo)

	return sb.Build()
}

func (c PostgresChoreoRepository) buildGetChoreoByID(choreoID int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllChoreo)
	sb.From(tableMasterChoreo)
	sb.Where(sb.Equal("choreo_id", choreoID))

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
		&result.TempPrice,
	)

	return result, err
}

func (c PostgresChoreoRepository) scanChoreoDataWithRelatedIds(row *sql.Rows) (result model.ChoreographyModel, musicId int64, cgpherId int64, err error) {
	result, err = c.scanChoreoData(row)

	return result, result.MusicID.Int64, result.ChoreographerID.Int64, err
}
