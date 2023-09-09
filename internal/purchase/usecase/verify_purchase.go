package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kora-Dance/koradance-backend/internal/common/constants"
	"github.com/Kora-Dance/koradance-backend/internal/helper"
	entity2 "github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/Kora-Dance/koradance-backend/pkg/storekit"
)

const (
	PrefixProductID = "com.nhara.kora.product"
)

func (c ChoreoPurchaseUseCaseImpl) VerifyPurchaseChoreo(ctx context.Context, userID int64, paymentData entity2.VerifyPaymentAppleIAPEntity) (*entity2.ChoreoPurchaseEntity, error) {
	// First, check if choreo already purchased
	purchasedData, err := c.baseRepo.ChoreoPurchaseRepository().GetPurchasedChoreoByID(ctx, userID, paymentData.ChoreoID)
	if err != nil {
		return nil, err
	}
	if purchasedData != nil {
		purchased := helper.ChoreoPurchaseModelToEntity(*purchasedData)
		return &purchased, nil
	}

	storekitData, err := c.storeKitM.VerifyPayment(paymentData)
	if err != nil {
		return nil, err
	}

	if storekitData.Status != constants.StoreKitVerifyStatusOk {
		return nil, errors.New(fmt.Sprintf("app store payment verification status not ok, got : %d", storekitData.Status))
	}

	isValid := c.checkPurchasedChoreo(paymentData.ChoreoID, storekitData.Receipt.InApp)
	if !isValid {
		return nil, errors.New("payment verification failed")
	}

	purchasedChoreo := entity2.ChoreoPurchaseEntity{
		UserID:   userID,
		ChoreoID: paymentData.ChoreoID,
		Receipt:  paymentData.ReceiptData,
		Status:   constants.ChoreoPurchaseStatusVerified,
	}

	verifyResult, err := c.baseRepo.ChoreoPurchaseRepository().InsertPurchasedChoreo(ctx, purchasedChoreo)
	if err != nil {
		return nil, err
	}
	submitEntityResult := helper.ChoreoPurchaseModelToEntity(*verifyResult)
	return &submitEntityResult, nil
}

func (c ChoreoPurchaseUseCaseImpl) checkPurchasedChoreo(choreoID int64, iapReceiptList []storekit.InAppPurchaseReceiptData) bool {
	for _, iapReceipt := range iapReceiptList {
		if iapReceipt.ProductID == fmt.Sprintf("%s.%d", PrefixProductID, choreoID) {
			return true
		}
	}
	return false
}
