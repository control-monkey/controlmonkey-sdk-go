package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/blueprint"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.Î
	sess := session.New()

	// Create a new instance of the services's client with a Session.
	// Optional controlmonkey.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// services specific configuration.
	svc := blueprint.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read blueprint.
	blueprintId := "blp-lzwmxv8q8f"
	t, err := svc.ListBlueprintNamespaceMappings(ctx, blueprintId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to read blueprint: %v", err)
	}

	// Output blueprint, if exists.
	if t != nil {
		log.Printf("Blueprint namespace mappings: %s",
			stringutil.Stringify(t))
	}
}
