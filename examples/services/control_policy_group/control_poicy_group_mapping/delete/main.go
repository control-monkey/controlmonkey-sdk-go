package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/services/control_policy_group"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
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
	svc := control_policy_group.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create ControlPolicyGroupMapping.
	_, err := svc.DeleteControlPolicyGroupMapping(ctx, &control_policy_group.ControlPolicyGroupMapping{
		ControlPolicyGroupId: controlmonkey.String("cmpg-2"),
		TargetId:             controlmonkey.String("ns-80lf0h8t6f"),
		TargetType:           controlmonkey.String("namespace"),
	})
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete ControlPolicyGroupMapping: %v", err)
	}

	log.Println("ControlPolicyGroupMapping was deleted successfully")
}
