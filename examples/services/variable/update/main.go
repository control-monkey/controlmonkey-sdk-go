package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/variable"
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
	svc := variable.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Update variable.
	out, err := svc.UpdateVariable(ctx, controlmonkey.String("var-123"), &variable.Variable{
		Value: controlmonkey.String("value"),
	})
	if err != nil {
		log.Fatalf("Control Monkey: failed to update variable: %v", err)
	}

	// Output variable, if was updated.
	if out.Variable != nil {
		log.Printf("Variable %q: %s",
			controlmonkey.StringValue(out.Variable.ID),
			stringutil.Stringify(out.Variable))
	}
}
