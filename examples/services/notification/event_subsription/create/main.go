package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"

	"github.com/control-monkey/controlmonkey-sdk-go/services/notification"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
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

	t := &notification.EventSubscription{
		NotificationEndpointId: controlmonkey.String("ne-p3p0vmuhde"),
		Scope:                  controlmonkey.String("organization"),
		EventType:              controlmonkey.String("aws::consoleOperation"),
	}

	out, err := svc.CreateEventSubscription(ctx, t)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create event subscription: %v", err)
	}

	if out != nil {
		log.Printf("Event subscription %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
