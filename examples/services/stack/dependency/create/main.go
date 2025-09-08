package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"
	"github.com/control-monkey/controlmonkey-sdk-go/services/namespace"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack"
	"log"
)

func main() {
	sess := session.New()

	nsSvc := namespace.New(sess)
	stackSvc := stack.New(sess)
	ctx := context.Background()

	ns, err := nsSvc.CreateNamespace(ctx, &namespace.Namespace{
		Name:        controlmonkey.String("example-dep-ns"),
		Description: controlmonkey.String("namespace for dependency examples"),
	})
	if err != nil {
		log.Fatalf("ControlMonkey: failed to create namespace: %v", err)
	}

	data := &stack.Data{
		DeploymentBehavior: &stack.DeploymentBehavior{DeployOnPush: controlmonkey.Bool(false)},
		DeploymentApprovalPolicy: &cross_models.DeploymentApprovalPolicy{
			Rules: []*cross_models.DeploymentApprovalPolicyRule{{Type: controlmonkey.String("requireApproval")}},
		},
		VcsInfo: &stack.VcsInfo{
			ProviderId: controlmonkey.String("vcsp-123"),
			RepoName:   controlmonkey.String("repo"),
			Path:       controlmonkey.String("path"),
			Branch:     controlmonkey.String("main"),
		},
		RunTrigger:   &cross_models.RunTrigger{Patterns: controlmonkey.StringSlice("*")},
		IacConfig:    &cross_models.IacConfig{TerraformVersion: controlmonkey.String("1.4.5")},
		Policy:       &stack.Policy{TtlConfig: &stack.TtlConfig{Ttl: &stack.TtlDefinition{Type: controlmonkey.String("hours"), Value: controlmonkey.Int(3)}}},
		RunnerConfig: &stack.RunnerConfig{Mode: controlmonkey.String("managed")},
		AutoSync:     &cross_models.AutoSync{DeployWhenDriftDetected: controlmonkey.Bool(true)},
	}

	parent, err := stackSvc.CreateStack(ctx, &stack.Stack{
		IacType:     controlmonkey.String("terraform"),
		NamespaceId: ns.ID,
		Name:        controlmonkey.String("example-parent"),
		Description: controlmonkey.String("parent stack"),
		Data:        data,
	})
	if err != nil {
		log.Fatalf("ControlMonkey: failed to create parent stack: %v", err)
	}

	child, err := stackSvc.CreateStack(ctx, &stack.Stack{
		IacType:     controlmonkey.String("terraform"),
		NamespaceId: ns.ID,
		Name:        controlmonkey.String("example-child"),
		Description: controlmonkey.String("child stack"),
		Data:        data,
	})
	if err != nil {
		log.Fatalf("ControlMonkey: failed to create child stack: %v", err)
	}

	input := &stack.Dependency{
		StackId:          child.ID,
		DependsOnStackId: parent.ID,
		References: []*stack.DependencyRef{{
			OutputOfStackToDependOn: controlmonkey.String("output_key"),
			InputForStack:           controlmonkey.String("var_name"),
		}},
		TriggerOption: controlmonkey.String("always"),
	}

	out, err := stackSvc.CreateDependency(ctx, input)
	if err != nil {
		log.Fatalf("ControlMonkey: failed to create dependency: %v", err)
	}

	if out != nil {
		log.Printf("Dependency %q: %s", controlmonkey.StringValue(out.ID), stringutil.Stringify(out))
	}
}
