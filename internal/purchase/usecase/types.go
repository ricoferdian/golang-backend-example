package usecase

import (
	"kora-backend/internal/common/storekit"
	"kora-backend/internal/domain/common"
	"kora-backend/internal/domain/purchase"
)

type ChoreoPurchaseUseCaseImpl struct {
	baseRepo  common.BaseRepository
	storeKitM *storekit.StoreKitModule
}

func NewChoreoPurchaseUseCase(baseRepo common.BaseRepository, storeKitM *storekit.StoreKitModule) purchase.ChoreoPurchaseUseCase {
	return &ChoreoPurchaseUseCaseImpl{
		baseRepo:  baseRepo,
		storeKitM: storeKitM,
	}
}
