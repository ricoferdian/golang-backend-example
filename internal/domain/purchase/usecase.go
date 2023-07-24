package purchase

import (
	"context"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
)

type ChoreoPurchaseUseCase interface {
	GetPurchasedChoreo(ctx context.Context, userID int64) ([]entity2.ChoreoPurchaseEntity, error)
	VerifyPurchaseChoreo(ctx context.Context, userID int64, paymentData entity2.VerifyPaymentAppleIAPEntity) (*entity2.ChoreoPurchaseEntity, error)
}
