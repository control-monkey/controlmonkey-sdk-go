package main

import (
	"context"
	"log"
	"os"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/notification"
)

func main() {
	sess := session.New()
	svc := notification.New(sess)

	ctx := context.Background()

	// Read app ID from env for convenience when running locally
	appId := os.Getenv("CONTROL_MONKEY_SLACK_APP_ID")
	if appId == "" {
		log.Fatalf("CONTROL_MONKEY_SLACK_APP_ID is not set")
	}
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
