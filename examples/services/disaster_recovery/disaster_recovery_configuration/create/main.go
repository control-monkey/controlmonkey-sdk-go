package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/services/disaster_recovery"
	"log"

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
	svc := disaster_recovery.New(sess)

	// Create a new context.
	ctx := context.Background()

	bs := &disaster_recovery.BackupStrategy{
		IncludeManagedResources: controlmonkey.Bool(false),
		Mode:                    controlmonkey.String("default"),
		VcsInfo: &disaster_recovery.VcsInfo{
			ProviderId: controlmonkey.String("vcsp-123"),
			RepoName:   controlmonkey.String("terraform"),
			Branch:     controlmonkey.String("main"),
		},
		Groups: nil,
	}

	s := &disaster_recovery.DisasterRecoveryConfiguration{
		Scope:          controlmonkey.String("aws"),
		CloudAccountId: controlmonkey.String("123456789"),
		BackupStrategy: bs,
	}
	// Create disaster_recovery.
	out, err := svc.CreateDisasterRecoveryConfiguration(ctx, s)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create entity: %v", err)
	}

	// Output entity, if was created.
	if out != nil {
		log.Printf("Entity %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
