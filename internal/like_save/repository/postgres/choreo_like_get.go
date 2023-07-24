package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	sq "github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
)

// GetLikedChoreoByChoreoIDsMap gets all liked choreos by choreoIDs in a map
func (p PostgresLikeSaveRepository) GetLikedChoreoByChoreoIDsMap(ctx context.Context, userID int64, choreoIDs []int64) (map[int64]model.ChoreoLikeModel, error) {
	if userID == 0 {
		return nil, errors.New("userID must be supplied")
	}
	query, args := p.buildGetLikedByChoreoIDs(userID, choreoIDs)
	rows, err := p.dbCli.QueryContext(ctx, p.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[int64]model.ChoreoLikeModel)
	for rows.Next() {
		sdata, err := p.scanLikeData(rows)
		if err != nil {
			return nil, err
		}
		data[sdata.ChoreoID] = sdata
	}

	return data, nil
}

// scanLikeData scans a single row from the sql query
func (p PostgresLikeSaveRepository) scanLikeData(row *sql.Rows) (result model.ChoreoLikeModel, err error) {
	err = row.Scan(
		&result.ChoreoID,
		&result.UserID,
	)

	return result, err
}

// buildGetLikedByChoreoIDs builds the query to get all liked choreos by choreoIDs
func (c PostgresLikeSaveRepository) buildGetLikedByChoreoIDs(userID int64, choreoIDs []int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectLike)
	sb.From(tableSelectLike)

	anySb := sq.Buildf("%v WHERE choreo_id = any(%v) AND user_id = %v", sb, pq.Array(choreoIDs), userID)

	return anySb.Build()
}
