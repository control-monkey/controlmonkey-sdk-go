package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/services/disaster_recovery"
	"log"

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
	svc := disaster_recovery.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Delete disaster_recovery.
	entityId := "drc-123"
	_, err := svc.DeleteDisasterRecoveryConfiguration(ctx, entityId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete entity: %v", err)
	}

	log.Println("Entity was deleted successfully")
}
