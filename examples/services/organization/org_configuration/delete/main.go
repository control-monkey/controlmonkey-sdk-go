package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/organization"
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
	svc := organization.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Delete namespace.
	_, err := svc.DeleteOrgConfiguration(ctx)
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete org configuration: %v", err)
	}

	log.Println("org configuration was deleted successfully")
}
