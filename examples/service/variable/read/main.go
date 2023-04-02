package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/service/variable"
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

	// Read variable.
	out, err := svc.ReadVariable(ctx, &variable.ReadVariableInput{VariableId: controlmonkey.String("var-123")})
	if err != nil {
		log.Fatalf("Control Monkey: failed to read variable: %v", err)
	}

	// Output variable, if exists.
	if out.Variable != nil {
		log.Printf("Variable %q: %s",
			controlmonkey.StringValue(out.Variable.ID),
			stringutil.Stringify(out.Variable))
	}
}
