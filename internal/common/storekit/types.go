package storekit

import "kora-backend/app/helper"

type StoreKitModule struct {
	storeKitPassword string
	storeKitConfig   *helper.StoreKitVerifyConfig
}

type StoreKitVerifyReceiptRequest struct {
	ReceiptData           string `json:"receipt-data"`
	Password              string `json:"password"`
	ExcludeOldTransaction bool   `json:"exclude-old-transactions"`
}

type StoreKitVerifyReceiptResponse struct {
	Receipt     StoreKitVerifyReceiptDetail `json:"receipt"`
	Environment string                      `json:"environment"`
	Status      int                         `json:"status"`
}

type StoreKitVerifyReceiptDetail struct {
	InApp []InAppPurchaseReceiptData `json:"in_app"`
}

type InAppPurchaseReceiptData struct {
	ProductID     string `json:"product_id"`
	TransactionID string `json:"transaction_id"`
	PurchaseDate  string `json:"purchase_date"`
	OwnershipType string `json:"in_app_ownership_type"`
}

func NewStoreKitModule(storeKitConfig *helper.StoreKitVerifyConfig) (*StoreKitModule, error) {
	password, err := getStorekitPassEnv()
	if err != nil {
		return nil, err
	}
	return &StoreKitModule{
		storeKitPassword: password,
		storeKitConfig:   storeKitConfig,
	}, nil
}
