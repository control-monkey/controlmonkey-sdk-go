package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/control_policy"

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

	// List all ControlPolicyMapping.
	out, err := svc.ListControlPolicyMappings(ctx, "pol-zxwvyemmbs2")
	if err != nil {
		log.Fatalf("Control Monkey: failed to read ControlPolicyMapping: %v", err)
	}

	// Output all ControlPolicyMapping, if any.
	if len(out) > 0 {
		log.Printf("Number of ControlPolicyMapping: %d", len(out))
		for _, v := range out {
			log.Printf("ControlPolicyMapping %q", stringutil.Stringify(v))
		}
	}
}
