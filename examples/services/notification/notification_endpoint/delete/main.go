package main

import (
	"context"
	"log"
	"os"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
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

	// Delete notification endpoint.
	notificationEndpointId := os.Getenv("CONTROL_MONKEY_NOTIFICATION_ENDPOINT_ID")
	if notificationEndpointId == "" {
		log.Fatalf("CONTROL_MONKEY_NOTIFICATION_ENDPOINT_ID is not set")
	}
	_, err := svc.DeleteNotificationEndpoint(ctx, notificationEndpointId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete notification endpoint: %v", err)
	}

	log.Println("Notification endpoint was deleted successfully")
}
