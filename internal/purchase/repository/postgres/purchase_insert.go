package postgres

import (
	"context"
	"database/sql"
	sq "github.com/huandu/go-sqlbuilder"
	"kora-backend/internal/entity"
	"kora-backend/internal/model"
	"kora-backend/internal/purchase/helper"
)

func (c PostgresChoreoPurchaseRepository) InsertPurchasedChoreo(ctx context.Context, purchaseData entity.ChoreoPurchaseEntity) (*model.ChoreoPurchaseModel, error) {
	query, args := c.buildInsertLearningHistory(purchaseData)
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
			&purchaseData.ChoreoPurchaseID,
		)
		break
	}
	result := helper.ChoreoPurchaseEntityToModel(purchaseData)

	return &result, nil
}

func (c PostgresChoreoPurchaseRepository) buildInsertLearningHistory(purchaseData entity.ChoreoPurchaseEntity) (string, []interface{}) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto(tablePurchasedChoreo)
	ib.Cols(columnInsertPurchasedChoreo)
	ib.Values(
		purchaseData.ChoreoID,
		purchaseData.UserID,
		purchaseData.Receipt,
		purchaseData.Status,
	)
	addQ := sq.Buildf("%v RETURNING choreo_purchase_id", ib)

	return addQ.Build()
}
