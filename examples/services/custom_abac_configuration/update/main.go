package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/custom_abac_configuration"

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
	svc := custom_abac_configuration.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Update CustomAbacConfiguration.
	out, err := svc.UpdateCustomAbacConfiguration(ctx, "cac-123", &custom_abac_configuration.CustomAbacConfiguration{
		Name: controlmonkey.String("Updated Custom Abac Configuration test"),
		Roles: []*custom_abac_configuration.Role{
			{
				OrgId:   controlmonkey.String("o-123"),
				OrgRole: controlmonkey.String("admin"),
				TeamIds: controlmonkey.StringSlice("team-123"),
			},
		},
	})
	if err != nil {
		log.Fatalf("Control Monkey: failed to update CustomAbacConfiguration: %v", err)
	}

	// Output CustomAbacConfiguration
	if out != nil {
		log.Printf("CustomAbacConfiguration %q", stringutil.Stringify(out))
	}
}
