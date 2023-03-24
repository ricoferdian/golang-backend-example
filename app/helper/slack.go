package helper

import (
	slack "kora-backend/internal/common/slackwebhook"
	"os"
)

func SendServiceStartAlert(slackModule *slack.SlackWebhookModule) error {
	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Hostname", Value: getHostname()})
	payload := slack.Payload{
		Text:        "Service is starting",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slackModule.SendAlertWebhook(payload)
	if err != nil {
		return err
	}
	return nil
}

func SendServiceFailureAlert(slackModule *slack.SlackWebhookModule, errMsg error) error {
	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Hostname", Value: getHostname()})
	attachment1.AddField(slack.Field{Title: "Cause", Value: errMsg.Error()})
	payload := slack.Payload{
		Text:        "Service is has been stopped",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slackModule.SendAlertWebhook(payload)
	if err != nil {
		return err
	}
	return nil
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown-hostname"
	}
	return hostname
}
