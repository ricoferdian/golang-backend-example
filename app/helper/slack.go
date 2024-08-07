package helper

import (
	"errors"
	slack "github.com/Kora-Dance/koradance-backend/pkg/slackwebhook"
	"os"
)

func SendServiceStartAlert(slackModule *slack.SlackWebhookModule) error {
	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Hostname", Value: GetHostname()})
	payload := slack.Payload{
		Text:        "Service is starting",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slackModule.SendSlackWebhook(payload)
	if err != nil {
		return err
	}
	return nil
}

func SendWhatsappMessageSlack(slackModule *slack.SlackWebhookModule, phoneNumber, message string) error {
	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: phoneNumber, Value: message})
	payload := slack.Payload{
		Text:        "Whatsapp Message",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slackModule.SendSlackWebhook(payload)
	if err != nil {
		return err
	}
	return nil
}

func SendServiceFailureAlert(slackModule *slack.SlackWebhookModule, errMsg error) error {
	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Hostname", Value: GetHostname()})
	attachment1.AddField(slack.Field{Title: "Cause", Value: errMsg.Error()})
	payload := slack.Payload{
		Text:        "Service is has been stopped",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slackModule.SendSlackWebhook(payload)
	if err != nil {
		return err
	}
	return nil
}

// getSecretEnv used to get jwt secret key from environment variable.
func GetSlackWebhookAlertUrl() (string, error) {
	env := os.Getenv("ALERT_SLACK_WEBHOOK_URL")
	if env == "" {
		return "", errors.New("unable to get slack webhook URL")
	}
	return env, nil
}
