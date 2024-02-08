package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/blueprint"

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
	svc := blueprint.New(sess)

	// Create a new context.
	ctx := context.Background()

	t := &blueprint.BlueprintNamespaceMapping{
		BlueprintId: controlmonkey.String("blp-lzwmxv8q8f"),
		NamespaceId: controlmonkey.String("ns-7x4pqdfdt1"),
	}

	_, err := svc.CreateBlueprintNamespaceMapping(ctx, t)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create blueprint namespace mapping: %v", err)
	}

	log.Println("Blueprint namespace mapping was created successfully")
}
