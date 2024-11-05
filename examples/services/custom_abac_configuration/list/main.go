package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/custom_abac_configuration"

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
	svc := custom_abac_configuration.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create CustomAbacConfiguration.
	out, err := svc.ListCustomAbacConfigurations(ctx, nil, nil)
	if err != nil {
		log.Fatalf("Control Monkey: failed to read CustomAbacConfiguration: %v", err)
	}

	// Output CustomAbacConfiguration, if was updated.
	if out != nil {
		log.Printf("CustomAbacConfiguration %q", stringutil.Stringify(out))
	}
}
