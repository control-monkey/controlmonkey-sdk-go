package main

import (
	"context"
	"log"
	"os"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/notification"
)

func main() {
	sess := session.New()
	svc := notification.New(sess)

	ctx := context.Background()

	appId := os.Getenv("CONTROL_MONKEY_SLACK_APP_ID")
	if appId == "" {
		log.Fatalf("CONTROL_MONKEY_SLACK_APP_ID is not set")
	}

	_, err := svc.DeleteNotificationSlackApp(ctx, appId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete notification slack app: %v", err)
	}
}
