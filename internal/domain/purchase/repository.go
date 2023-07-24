package purchase

import (
	"context"
	"github.com/Kora-Dance/koradance-backend/internal/model"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
)

type ChoreoPurchaseDatabaseRepo interface {
	GetPurchasedChoreoByID(ctx context.Context, userID int64, choreoID int64) (*model.ChoreoPurchaseModel, error)
	GetPurchasedChoreoByUserIDMap(ctx context.Context, userID int64) (map[int64]model.ChoreoPurchaseModel, error)
	GetPurchasedChoreoByUserID(ctx context.Context, userID int64) ([]model.ChoreoPurchaseModel, []int64, error)
	InsertPurchasedChoreo(ctx context.Context, purchaseModel entity.ChoreoPurchaseEntity) (*model.ChoreoPurchaseModel, error)
}

type ChoreoPurchaseCacheRepo interface {
}

type ChoreoPurchaseRepository interface {
	ChoreoPurchaseDatabaseRepo
	ChoreoPurchaseCacheRepo
}
