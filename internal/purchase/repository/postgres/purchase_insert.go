package postgres

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	sq "github.com/huandu/go-sqlbuilder"
)

func (c PostgresChoreoPurchaseRepository) InsertPurchasedChoreo(ctx context.Context, purchaseData entity.ChoreoPurchaseEntity) (*model.ChoreoPurchaseModel, error) {
	query, args := c.buildInsertLearningHistory(purchaseData)
	purchaseID, err := c.dbCli.QueryContext(ctx, c.dbCli.Rebind(query), args...)
	if err != nil {
		return nil, err
	}
	defer purchaseID.Close()
	for purchaseID.Next() {
		err = purchaseID.Scan(
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
