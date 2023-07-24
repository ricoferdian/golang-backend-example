package storekit

import (
	"bytes"
	"errors"
	"github.com/Kora-Dance/koradance-backend/pkg/entity"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	verifyReceiptEndpoint = "/verifyReceipt"
)

// getStorekitPassEnv used to get jwt secret key from environment variable.
func getStorekitPassEnv() (string, error) {
	env := os.Getenv("KORA_STOREKIT_PASSWORD")
	if env == "" {
		return "", errors.New("Unable to get kora storekit password environment variable")
	}
	return env, nil
}

func (s *StoreKitModule) VerifyPayment(paymentData entity.VerifyPaymentAppleIAPEntity) (StoreKitVerifyReceiptResponse, error) {
	// TODO move out the external http call client to a shared module
	client := &http.Client{
		Timeout: time.Duration(s.storeKitConfig.Timeout) * time.Millisecond,
	}

	var respBody StoreKitVerifyReceiptResponse

	req := StoreKitVerifyReceiptRequest{
		ReceiptData:           paymentData.ReceiptData,
		Password:              s.storeKitPassword,
		ExcludeOldTransaction: true,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return respBody, err
	}

	resp, err := client.Post(s.storeKitConfig.Hostname+verifyReceiptEndpoint, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return respBody, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return respBody, err
	}

	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}
