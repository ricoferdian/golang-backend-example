//go:generate easytags $GOFILE json
package entity

type VerifyPaymentAppleIAPEntity struct {
	ChoreoID    int64  `json:"choreo_id" validate:"required" form:"choreo_id"`
	ReceiptData string `json:"receipt_data" validate:"required" form:"receipt_data"`
}
