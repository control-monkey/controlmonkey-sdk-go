package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/notification"
)

func main() {
	sess := session.New()
	svc := notification.New(sess)

	ctx := context.Background()

	appId := "nsa-123"

	_, err := svc.DeleteNotificationSlackApp(ctx, appId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete notification slack app: %v", err)
	}
}
