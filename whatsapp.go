package main

import (
	"context"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func sendNewsletterMessage(client *whatsmeow.Client, number string, msg string) error {
	jid, _ := types.ParseJID("120363379766529413@newsletter")

	msgWhatsApp := &waE2E.Message{
		Conversation: proto.String(msg),
	}

	_, err := client.SendMessage(context.Background(), jid, msgWhatsApp)
	return err
}
