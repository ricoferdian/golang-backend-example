package postgres

import (
	"context"
	"database/sql"
	sq "github.com/huandu/go-sqlbuilder"
	"github.com/Kora-Dance/koradance-backend/internal/model"
)

func (c PostgresChoreoPurchaseRepository) GetPurchasedChoreoByID(ctx context.Context, userID int64, choreoID int64) (result *model.ChoreoPurchaseModel, err error) {
	query, args := c.buildGetPurchasedChoreoByUserID(userID, choreoID)
	rows, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		choreoData, err := c.scanPurchasedChoreoData(rows)
		if err != nil {
			return nil, err
		}
		return &choreoData, nil
	}
	return result, nil
}

func (c PostgresChoreoPurchaseRepository) getPurchasedListRows(ctx context.Context, userID int64) (rows *sql.Rows, err error) {
	query, args := c.buildGetPurchasedChoreoByUserID(userID, 0)
	rows, err = c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (c PostgresChoreoPurchaseRepository) GetPurchasedChoreoByUserIDMap(ctx context.Context, userID int64) (result map[int64]model.ChoreoPurchaseModel, err error) {
	rows, err := c.getPurchasedListRows(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mapResult := make(map[int64]model.ChoreoPurchaseModel)
	for rows.Next() {
		choreoData, err := c.scanPurchasedChoreoData(rows)
		if err != nil {
			return nil, err
		}
		mapResult[choreoData.ChoreoID] = choreoData
	}
	return mapResult, nil
}

func (c PostgresChoreoPurchaseRepository) GetPurchasedChoreoByUserID(ctx context.Context, userID int64) (result []model.ChoreoPurchaseModel, choreoIds []int64, err error) {
	rows, err := c.getPurchasedListRows(ctx, userID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var choreoIdExist = make(map[int64]bool)
	for rows.Next() {
		choreoDetData, choreoId, err := c.scanPurchasedChoreoDataWithRelatedIds(rows)
		if err != nil {
			return nil, nil, err
		}
		if choreoId != 0 && !choreoIdExist[choreoId] {
			choreoIds = append(choreoIds, choreoId)
			choreoIdExist[choreoId] = true
		}
		result = append(result, choreoDetData)
	}
	return result, choreoIds, nil
}

func (c PostgresChoreoPurchaseRepository) buildGetPurchasedChoreoByUserID(userID int64, choreoID int64) (string, []interface{}) {
	sb := sq.NewSelectBuilder()
	sb.Select(columnSelectPurchasedChoreoNoReceipt)
	sb.From(tablePurchasedChoreo)

	sb.Where(sb.Equal("user_id", userID))
	if choreoID != 0 {
		sb.Where(sb.Equal("choreo_id", choreoID))
	}

	return sb.Build()
}

func (c PostgresChoreoPurchaseRepository) scanPurchasedChoreoData(row *sql.Rows) (result model.ChoreoPurchaseModel, err error) {
	err = row.Scan(
		&result.ChoreoPurchaseID,
		&result.ChoreoID,
		&result.UserID,
		&result.Status,
	)

	return result, err
}

func (c PostgresChoreoPurchaseRepository) scanPurchasedChoreoDataWithRelatedIds(row *sql.Rows) (result model.ChoreoPurchaseModel, choreoId int64, err error) {
	result, err = c.scanPurchasedChoreoData(row)

	return result, result.ChoreoID, err
}
