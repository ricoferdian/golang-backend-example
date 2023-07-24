package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	sq "github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
)

func (p PostgresLikeSaveRepository) GetSavedChoreoByUserID(ctx context.Context, userID int64) ([]model.ChoreoSaveModel, error) {
	if userID == 0 {
		return nil, errors.New("userID must be supplied")
	}
	query, args := p.buildGetAllSaved(userID)
	rows, err := p.dbCli.QueryContext(ctx, p.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []model.ChoreoSaveModel
	for rows.Next() {
		sdata, err := p.scanData(rows)
		if err != nil {
			return nil, err
		}
		data = append(data, sdata)
	}

	return data, nil
}

func (p PostgresLikeSaveRepository) scanData(row *sql.Rows) (result model.ChoreoSaveModel, err error) {
	err = row.Scan(
		&result.ChoreoID,
		&result.UserID,
	)

	return result, err
}

func (p PostgresLikeSaveRepository) GetSavedChoreoByChoreoIDsMap(ctx context.Context, userID int64, choreoIDs []int64) (map[int64]model.ChoreoSaveModel, error) {
	if userID == 0 {
		return nil, errors.New("userID must be supplied")
	}
	query, args := p.buildGetSavedByChoreoIDs(userID, choreoIDs)
	rows, err := p.dbCli.QueryContext(ctx, p.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[int64]model.ChoreoSaveModel)
	for rows.Next() {
		sdata, err := p.scanData(rows)
		if err != nil {
			return nil, err
		}
		data[sdata.ChoreoID] = sdata
	}

	return data, nil
}

func (c PostgresLikeSaveRepository) buildGetAllSaved(userID int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectSave)
	sb.From(tableSelectSave)
	sb.Where(sb.Equal("user_id", userID))

	return sb.Build()
}

func (c PostgresLikeSaveRepository) buildGetSavedByChoreoIDs(userID int64, choreoIDs []int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectSave)
	sb.From(tableSelectSave)

	anySb := sq.Buildf("%v WHERE choreo_id = any(%v) AND user_id = %v", sb, pq.Array(choreoIDs), userID)

	return anySb.Build()
}
