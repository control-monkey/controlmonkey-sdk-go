package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack"
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
			DeployOnPush: controlmonkey.Bool(false),
		},
		DeploymentApprovalPolicy: &stack.DeploymentApprovalPolicy{
			Rules: []*cross_models.DeploymentApprovalPolicyRule{
				{
					Type: controlmonkey.String("requireApproval"),
				},
			},
		},
		VcsInfo: &stack.VcsInfo{
			ProviderId: controlmonkey.String("vcsp-123"),
			RepoName:   controlmonkey.String("repo"),
			Path:       controlmonkey.String("path"),
			Branch:     controlmonkey.String("main"),
		},
		RunTrigger: &stack.RunTrigger{Patterns: controlmonkey.StringSlice()},
		IacConfig: &stack.IacConfig{
			TerraformVersion: controlmonkey.String("1.4.5"),
		},
		Policy: &stack.Policy{
			TtlConfig: &stack.TtlConfig{
				Ttl: &stack.TtlDefinition{
					Type:  controlmonkey.String("hours"),
					Value: controlmonkey.Int(3),
				},
			},
		},
		RunnerConfig: &stack.RunnerConfig{
			Mode: controlmonkey.String("managed"),
		},
		AutoSync: &stack.AutoSync{
			DeployWhenDriftDetected: controlmonkey.Bool(true),
		},
	}

	s := &stack.Stack{
		IacType:     controlmonkey.String("terraform"),
		NamespaceId: controlmonkey.String("ns-hhpdqtybv3"),
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
