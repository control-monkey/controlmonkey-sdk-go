package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"

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

	// Create variable.
	out, err := svc.CreateVariable(ctx, &variable.Variable{
		Scope:         controlmonkey.String("stack"),
		ScopeId:       controlmonkey.String("stk-123"),
		Key:           controlmonkey.String("key"),
		Type:          controlmonkey.String("tfVar"),
		Value:         controlmonkey.String("value"),
		IsSensitive:   controlmonkey.Bool(false),
		IsOverridable: controlmonkey.Bool(false),
		Description:   controlmonkey.String("variable description"),
		ValueConditions: []*cross_models.Condition{
			{
				Operator: controlmonkey.String("in"),
				Value:    controlmonkey.Any(controlmonkey.StringSlice("value", "value2")),
			},
		},
	})
	if err != nil {
		log.Fatalf("Control Monkey: failed to create variable: %v", err)
	}

	// Output variable, if was created.
	if out.Variable != nil {
		log.Printf("Variable %q: %s",
			controlmonkey.StringValue(out.Variable.ID),
			stringutil.Stringify(out.Variable))
	}
}
