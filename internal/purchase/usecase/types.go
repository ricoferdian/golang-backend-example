package usecase

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/common"
	"github.com/Kora-Dance/koradance-backend/internal/domain/purchase"
	"github.com/Kora-Dance/koradance-backend/pkg/storekit"
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
