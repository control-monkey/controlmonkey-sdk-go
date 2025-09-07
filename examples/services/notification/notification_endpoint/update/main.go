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
	// All clients require a Session. The Session provides the client with
	// shared configuration such as credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the services's client with a Session.
	// Optional controlmonkey.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// services specific configuration.
	svc := notification.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Update notification endpoint.
	notificationEndpointId := os.Getenv("CONTROL_MONKEY_NOTIFICATION_ENDPOINT_ID")
	if notificationEndpointId == "" {
		log.Fatalf("CONTROL_MONKEY_NOTIFICATION_ENDPOINT_ID is not set")
	}
	t, err := svc.UpdateNotificationEndpoint(ctx, notificationEndpointId, &notification.Endpoint{
		Name: controlmonkey.String("notification endpoint2"),
		Url:  controlmonkey.String("x2"),
		NotificationEndpointSlackAppConfig: &notification.NotificationEndpointSlackAppConfig{
			NotificationSlackAppId: controlmonkey.String("nsa-REPLACE"),
			ChannelId:              controlmonkey.String("C0987654321"),
		},
		EmailAddresses: []*string{
			controlmonkey.String("example3@example.com"),
		},
	})
	if err != nil {
		log.Fatalf("Control Monkey: failed to update notification endpoint: %v", err)
	}

	// Output notification endpoint, if was updated.
	if t != nil {
		log.Printf("Notification endpoint %q: %s",
			controlmonkey.StringValue(t.ID),
			stringutil.Stringify(t))
	}
}
