package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/service/stack"
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
	svc := stack.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read a deployment.
	out, err := svc.ReadDeployment(ctx, &stack.ReadDeploymentInput{DeploymentId: controlmonkey.String("dpl-123")})
	if err != nil {
		log.Fatalf("Control Monkey: failed to read deployment: %v", err)
	}

	// Output deployment, if exists.
	if out.Deployment != nil {
		log.Printf("Deployment %q: %s",
			controlmonkey.StringValue(out.Deployment.ID),
			stringutil.Stringify(out.Deployment))
	}
}
