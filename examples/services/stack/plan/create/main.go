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

	// Create a new instance of the service's client with a Session.
	// Optional controlmonkey.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := stack.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create plan.
	out, err := svc.CreatePlan(ctx, &stack.CreatePlanInput{StackId: controlmonkey.String("stk-123")})
	if err != nil {
		log.Fatalf("Control Monkey: failed to create plan: %v", err)
	}

	// Output plan, if was created.
	if out.Plan != nil {
		log.Printf("Plan %q: %s",
			controlmonkey.StringValue(out.Plan.ID),
			stringutil.Stringify(out.Plan))
	}
}
