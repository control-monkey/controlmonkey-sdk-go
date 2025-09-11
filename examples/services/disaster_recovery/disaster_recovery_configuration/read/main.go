package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/disaster_recovery"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
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
	svc := disaster_recovery.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read disaster_recovery.
	drcId := "drc-123"
	out, err := svc.ReadDisasterRecoveryConfiguration(ctx, drcId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to read entity: %v", err)
	}

	// Output entity, if exists.
	if out != nil {
		log.Printf("Entity %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
