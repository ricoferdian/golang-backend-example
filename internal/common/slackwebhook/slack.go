package slackwebhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// getSecretEnv used to get jwt secret key from environment variable.
func getWebhookURL(key string) (string, error) {
	env := os.Getenv(key)
	if env == "" {
		return "", errors.New("Unable to get JWT secret key environment variable")
	}
	return env, nil
}

func (s *SlackWebhookModule) SendAlertWebhook(payload Payload) error {
	return s.sendWebhook(s.alertWebhookUrl, payload)
}

func (s *SlackWebhookModule) sendWebhook(webhookUrl string, payload Payload) error {
	postBody, _ := json.Marshal(payload)
	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(webhookUrl, "application/json", requestBody)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("error sending msg. Status: %v", resp.Status)
	}

	return nil
}
