//go:generate easytags $GOFILE json
package entity

type SecureOtpRequest struct {
	CountryCode       string `json:"country_code"`
	TransferMediaType int    `json:"transfer_media_type"`
	ReceiverIdentity  string `json:"receiver_identity"`
	Password          string `json:"password"`
}
