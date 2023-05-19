package usecase

import (
	"context"
	"errors"
	"fmt"
	"kora-backend/internal/common/constants"
	"kora-backend/internal/common/storekit"
	"kora-backend/internal/entity"
	"kora-backend/internal/purchase/helper"
)

const (
	PrefixProductID = "com.nhara.kora"
)

func (c ChoreoPurchaseUseCaseImpl) VerifyPurchaseChoreo(ctx context.Context, userID int64, paymentData entity.VerifyPaymentAppleIAPEntity) (*entity.ChoreoPurchaseEntity, error) {
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

	purchasedChoreo := entity.ChoreoPurchaseEntity{
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
