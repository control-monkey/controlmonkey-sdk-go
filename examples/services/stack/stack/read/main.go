package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack"
	"log"
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

	// Read stack.
	stackId := "stk-kpajfm381w"
	out, err := svc.ReadStack(ctx, stackId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to read stack: %v", err)
	}

	// Output stack, if exists.
	if out.Stack != nil {
		log.Printf("Stack %q: %s",
			controlmonkey.StringValue(out.Stack.ID),
			stringutil.Stringify(out.Stack))
	}
}
