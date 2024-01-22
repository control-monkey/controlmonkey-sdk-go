package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack"
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
	svc := stack.New(sess)

	// Create a new context.
	ctx := context.Background()

	// List all stacks.
	namespaceId := "ns-x82yjdyahc"
	out, err := svc.ListStacks(ctx, &stack.ListStacksParams{NamespaceId: controlmonkey.String(namespaceId)})
	if err != nil {
		log.Fatalf("Control Monkey: failed to read stacks: %v", err)
	}

	// Output all stacks, if any.
	if len(out.Stacks) > 0 {
		log.Printf("Number of stacks: %d", len(out.Stacks))
		for _, v := range out.Stacks {
			log.Printf("Stack %q: %s",
				controlmonkey.StringValue(v.ID),
				stringutil.Stringify(v))
		}
	}
}
