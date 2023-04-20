package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/namespace"
	"log"
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
	svc := namespace.New(sess)

	// Create a new context.
	ctx := context.Background()

	policy := &namespace.Policy{
		TtlConfig: &namespace.TtlConfig{
			MaxTtl: &namespace.TtlDefinition{
				Type:  controlmonkey.String("hours"),
				Value: controlmonkey.Int(3),
			},
			DefaultTtl: &namespace.TtlDefinition{
				Type:  controlmonkey.String("hours"),
				Value: controlmonkey.Int(2),
			}}}

	externalCredentials1 := namespace.ExternalCredentials{
		Type:                  controlmonkey.String("awsAssumeRole"),
		ExternalCredentialsId: controlmonkey.String("ext-123"),
	}
	externalCredentials := []*namespace.ExternalCredentials{&externalCredentials1}

	n := namespace.Namespace{
		Name:                controlmonkey.String("namespace1"),
		Description:         controlmonkey.String("description"),
		ExternalCredentials: &externalCredentials,
		Policy:              policy,
	}
	// Create namespace.
	out, err := svc.CreateNamespace(ctx, &n)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create namespace: %v", err)
	}

	// Output namespace, if was created.
	if out.Namespace != nil {
		log.Printf("Namespace %q: %s",
			controlmonkey.StringValue(out.Namespace.ID),
			stringutil.Stringify(out.Namespace))
	}
}
