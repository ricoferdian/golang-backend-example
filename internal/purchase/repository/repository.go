package repository

import (
	"github.com/Kora-Dance/koradance-backend/internal/domain/purchase"
)

type ChoreoPurchaseRepositoryImpl struct {
	purchase.ChoreoPurchaseDatabaseRepo
	purchase.ChoreoPurchaseCacheRepo
}

func NewChoreoPurchaseRepository(
	dbRepo purchase.ChoreoPurchaseDatabaseRepo,
	redisRepo purchase.ChoreoPurchaseCacheRepo,
) purchase.ChoreoPurchaseRepository {
	return &ChoreoPurchaseRepositoryImpl{
		dbRepo,
		redisRepo,
	}
}
