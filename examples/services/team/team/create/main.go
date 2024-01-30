package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/team"

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

	// Create a new instance of the services's client with a Session.
	// Optional controlmonkey.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// services specific configuration.
	svc := team.New(sess)

	// Create a new context.
	ctx := context.Background()

	t := &team.Team{
		Name:        controlmonkey.String("team1"),
		CustomIdpId: controlmonkey.String("abc"),
	}

	// Create team.
	out, err := svc.CreateTeam(ctx, t)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create team: %v", err)
	}

	// Output team, if was created.
	if out != nil {
		log.Printf("Team %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
