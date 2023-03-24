package postgres

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
	"kora-backend/internal/model"
)

func (c PostgresChoreographerRepository) GetChoreographerByIdsMap(ctx context.Context, choreographerIDs []int64) (map[int64]model.ChoreographerModel, error) {
	if len(choreographerIDs) == 0 {
		return nil, errors.New("choreographerIDs must be supplied")
	}
	query, args := c.buildGetChoreographerByIDs(choreographerIDs)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int64]model.ChoreographerModel)
	for rows.Next() {
		crgpherData, err := c.scanChoreographerData(rows)
		if err != nil {
			return nil, err
		}
		result[crgpherData.ChoreographerID] = crgpherData
	}
	return result, nil
}

func (c PostgresChoreographerRepository) buildGetChoreographerByIDs(choreographerIDs []int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllChoreographer)
	sb.From(tableMasterChoreographer)

	anySb := sq.Buildf("%v WHERE choreographer_id = any(%v)", sb, pq.Array(choreographerIDs))

	return anySb.Build()
}

func (c PostgresChoreographerRepository) GetChoreographerById(ctx context.Context, choreographerID int64) (*model.ChoreographerModel, error) {
	if choreographerID == 0 {
		return nil, errors.New("choreographerID must be supplied")
	}
	query, args := c.buildGetChoreographerByID(choreographerID)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var choreographerData model.ChoreographerModel
	for rows.Next() {
		choreographerData, err = c.scanChoreographerData(rows)
		if err != nil {
			return nil, err
		}
		break
	}
	return &choreographerData, nil
}

func (c PostgresChoreographerRepository) buildGetChoreographerByID(choreographerID int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllChoreographer)
	sb.From(tableMasterChoreographer)

	sb.Where(sb.Equal("choreographer_id", choreographerID))

	return sb.Build()
}

func (c PostgresChoreographerRepository) scanChoreographerData(row *sql.Rows) (result model.ChoreographerModel, err error) {
	err = row.Scan(
		&result.ChoreographerID,
		&result.ChoreographerName,
	)

	return result, err
}
