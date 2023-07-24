package whatsapp

import (
	"github.com/Kora-Dance/koradance-backend/app/helper"
	slack "github.com/Kora-Dance/koradance-backend/pkg/slackwebhook"
	"go.mau.fi/whatsmeow"
)

type WhatsappModule struct {
	client      *whatsmeow.Client
	slackModule *slack.SlackWebhookModule
	dbCfg       *helper.DatabaseConfig
}

func NewWhatsappModule(dbCfg *helper.DatabaseConfig, slackModule *slack.SlackWebhookModule) *WhatsappModule {
	client := getClient(dbCfg)
	return &WhatsappModule{
		client:      client,
		slackModule: slackModule,
		dbCfg:       dbCfg,
	}
}

func (m *WhatsappModule) Stop() {
	if m.client != nil {
		client.Disconnect()
	}
}
