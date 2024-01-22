package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/variable"
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
	svc := variable.New(sess)

	// Create a new context.
	ctx := context.Background()

	// List all variables.
	out, err := svc.ListVariables(ctx, &variable.ListVariablesInput{StackId: controlmonkey.String("stk-123")})
	if err != nil {
		log.Fatalf("Control Monkey: failed to read variables: %v", err)
	}

	// Output all variables, if any.
	if len(out.Variables) > 0 {
		log.Printf("Number of variables: %d", len(out.Variables))
		for _, v := range out.Variables {
			log.Printf("Variable %q: %s",
				controlmonkey.StringValue(v.ID),
				stringutil.Stringify(v))
		}
	}
}
