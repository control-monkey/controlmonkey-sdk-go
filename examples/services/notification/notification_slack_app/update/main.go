package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/notification"
)

func main() {
	sess := session.New()
	svc := notification.New(sess)

	ctx := context.Background()

	appId := "nsa-123"
	app := &notification.NotificationSlackApp{
		Name:         controlmonkey.String("my-renamed-slack-app"),
		BotAuthToken: controlmonkey.String("xoxb-updated"),
	}

	out, err := svc.UpdateNotificationSlackApp(ctx, appId, app)
	if err != nil {
		log.Fatalf("Control Monkey: failed to update notification slack app: %v", err)
	}

	if out != nil {
		log.Printf("Notification Slack App: %s", stringutil.Stringify(out))
	}
}
