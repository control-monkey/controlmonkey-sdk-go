package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/team"
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

	// Update team.
	teamId := "team-jcwfvxy84v"
	t, err := svc.UpdateTeam(ctx, teamId, &team.Team{
		Name:        controlmonkey.String("team2"),
		CustomIdpId: controlmonkey.String("abd"),
	})
	if err != nil {
		log.Fatalf("Control Monkey: failed to update team: %v", err)
	}

	// Output team, if was updated.
	if t != nil {
		log.Printf("Team %q: %s",
			controlmonkey.StringValue(t.ID),
			stringutil.Stringify(t))
	}
}
