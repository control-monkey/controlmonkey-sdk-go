package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/control_policy"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
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
	svc := control_policy.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create ControlPolicy.
	out, err := svc.UpdateControlPolicy(ctx, "pol-zefk9g04sf", &control_policy.ControlPolicy{
		Name: controlmonkey.String("Updated Policy test"),
	})
	if err != nil {
		log.Fatalf("Control Monkey: failed to update ControlPolicy: %v", err)
	}

	// Output ControlPolicy, if was updated.
	if out != nil {
		log.Printf("ControlPolicy %q", stringutil.Stringify(out))
	}
}
