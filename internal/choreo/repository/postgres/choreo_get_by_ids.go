package postgres

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
	"github.com/Kora-Dance/koradance-backend/internal/model"
)

func (c PostgresChoreoRepository) GetChoreoByIdsMap(ctx context.Context, choreoIDs []int64) (map[int64]model.ChoreographyModel, error) {
	if len(choreoIDs) == 0 {
		return nil, errors.New("choreoIDs must be supplied")
	}
	query, args := c.buildGetChoreoByIDs(choreoIDs)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int64]model.ChoreographyModel)
	for rows.Next() {
		choreoData, err := c.scanChoreoData(rows)
		if err != nil {
			return nil, err
		}
		result[choreoData.ChoreoID] = choreoData
	}
	return result, nil
}

func (c PostgresChoreoRepository) buildGetChoreoByIDs(choreoIDs []int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectAllChoreo)
	sb.From(tableMasterChoreo)

	anySb := sq.Buildf("%v WHERE choreo_id = any(%v)", sb, pq.Array(choreoIDs))

	return anySb.Build()
}
