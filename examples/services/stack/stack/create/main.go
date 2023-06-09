package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack"
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
	svc := stack.New(sess)

	// Create a new context.
	ctx := context.Background()

	data := &stack.Data{
		DeploymentBehavior: &stack.DeploymentBehavior{
			DeployOnPush:    controlmonkey.Bool(true),
			WaitForApproval: controlmonkey.Bool(true),
		},
		VcsInfo: &stack.VcsInfo{
			ProviderId: controlmonkey.String("vcsp-jgkig4q04e"),
			RepoName:   controlmonkey.String("workspace/repo"),
			Path:       controlmonkey.String("path/to/code"),
			Branch:     controlmonkey.String("main"),
		},
		RunTrigger: &stack.RunTrigger{Patterns: controlmonkey.StringSlicePointer("path/**/*")},
		IacConfig: &stack.IacConfig{
			TerraformVersion: controlmonkey.String("1.4.5"),
		},
		Policy: &stack.Policy{
			TtlConfig: &stack.TtlConfig{
				Ttl: &stack.TtlDefinition{
					Type:  controlmonkey.String("hours"),
					Value: controlmonkey.Int(3),
				}}},
	}

	s := &stack.Stack{
		IacType:     controlmonkey.String("terraform"),
		NamespaceId: controlmonkey.String("ns-x82yjdyahc"),
		Name:        controlmonkey.String("stack1"),
		Description: controlmonkey.String("description"),
		Data:        data,
	}
	// Create stack.
	out, err := svc.CreateStack(ctx, &stack.CreateStackInput{Stack: s})
	if err != nil {
		log.Fatalf("Control Monkey: failed to create stack: %v", err)
	}

	// Output stack, if was created.
	if out.Stack != nil {
		log.Printf("Stack %q: %s",
			controlmonkey.StringValue(out.Stack.ID),
			stringutil.Stringify(out.Stack))
	}
}
