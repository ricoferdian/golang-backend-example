package whatsapp

import (
	"context"
	"fmt"
	"github.com/Kora-Dance/koradance-backend/app/helper"
	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"log"
)

const (
	jidServer = "s.whatsapp.net"
)

var (
	client *whatsmeow.Client
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		//fmt.Println("Received a message!", v.Message.GetConversation())
		textMsg := "Hello, if you need more information about Kora, you can visit https://www.koradance.com"
		msg := &proto.Message{
			Conversation: &textMsg,
		}
		jid := v.Info.Sender
		_, err := client.SendMessage(context.Background(), jid, msg)
		if err != nil {
			log.Println("[WhatsappClient][Error] err", err)
		}
	}
}

func (m *WhatsappModule) SendMessage(message, phoneNumber string) error {
	jid := types.JID{
		User:   phoneNumber,
		Server: jidServer,
	}

	msg := &proto.Message{
		Conversation: &message,
	}
	if client != nil {
		_, err := client.SendMessage(context.Background(), jid, msg)
		if err != nil {
			client = getClient(m.dbCfg)
			return err
		}
		return nil
	}

	err := helper.SendWhatsappMessageSlack(m.slackModule, phoneNumber, message)
	if err != nil {
		return err
	}
	return nil
}

func getClient(cfg *helper.DatabaseConfig) *whatsmeow.Client {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", cfg.DriverName, cfg.Username, cfg.Password, cfg.Hostname, cfg.Port, cfg.DBName)
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("postgres", dsn, dbLog)
	if err != nil {
		log.Println("[WhatsappClient][Error]", err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		log.Println("[WhatsappClient][Error]", err)
	}
	//clientLog := waLog.Stdout("Client", "DEBUG", true)
	client = whatsmeow.NewClient(deviceStore, nil)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			log.Println("[WhatsappClient][Error]", err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				log.Println("[WhatsappClient][Event] QR code:", evt.Code)
			} else {
				log.Println("[WhatsappClient][Event] Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			log.Println("[WhatsappClient][Error]", err)
		}
	}
	return client
}
