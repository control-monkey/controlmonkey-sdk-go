package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/namespace"
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
	svc := namespace.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Update namespace.
	namespaceId := "ns-123"
	out, err := svc.UpdateNamespace(ctx, namespaceId, &namespace.Namespace{
		Name: controlmonkey.String("namespace2"),
	})
	if err != nil {
		log.Fatalf("Control Monkey: failed to update namespace: %v", err)
	}

	// Output namespace, if was updated.
	if out.Namespace != nil {
		log.Printf("Namespace %q: %s",
			controlmonkey.StringValue(out.Namespace.ID),
			stringutil.Stringify(out.Namespace))
	}
}
