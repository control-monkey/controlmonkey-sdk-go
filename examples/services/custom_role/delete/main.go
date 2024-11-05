package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/custom_role"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the service's client with a Session.
	// Optional controlmonkey.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := custom_role.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Delete CustomRole.
	_, err := svc.DeleteCustomRole(ctx, "cro-123")
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete CustomRole: %v", err)
	}

	log.Println("CustomRole was deleted successfully")
}
