package main

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func sendMessage(client *whatsmeow.Client, number string, message string) error {
	jid, err := types.ParseJID(number + "@s.whatsapp.net")
	if err != nil {
		return fmt.Errorf("invalid JID: %w", err)
	}
	_, err = client.SendMessage(context.Background(), jid, &waProto.Message{
		Conversation: proto.String(message),
	})
	return err
}
