package purchase

import (
	"context"
	"kora-backend/internal/entity"
)

type ChoreoPurchaseUseCase interface {
	GetPurchasedChoreo(ctx context.Context, userID int64) ([]entity.ChoreoPurchaseEntity, error)
	VerifyPurchaseChoreo(ctx context.Context, userID int64, paymentData entity.VerifyPaymentAppleIAPEntity) (*entity.ChoreoPurchaseEntity, error)
}
