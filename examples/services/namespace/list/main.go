package main

import (
	"context"
	"log"

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

	// List namespaces.
	namespaceId := "ns-7x4pqdfdt1"
	out, err := svc.ListNamespaces(ctx, &namespaceId, nil)
	if err != nil {
		log.Fatalf("Control Monkey: failed to read namespace: %v", err)
	}

	// Output namespace, if exists.
	if out != nil {
		log.Printf("Namespaces: %s",
			stringutil.Stringify(out))
	}

	name := "Dev"
	out, err = svc.ListNamespaces(ctx, nil, &name)
	if err != nil {
		log.Fatalf("Control Monkey: failed to read namespace: %v", err)
	}

	// Output namespace, if exists.
	if out != nil {
		log.Printf("Namespaces: %s",
			stringutil.Stringify(out))
	}
}
