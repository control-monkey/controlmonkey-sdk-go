package main

import (
	"context"
	"log"
	"os"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"
	"github.com/control-monkey/controlmonkey-sdk-go/services/namespace"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack"
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
		DeploymentBehavior: &cross_models.DeploymentBehavior{DeployOnPush: controlmonkey.Bool(false)},
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
		RunnerConfig: &cross_models.RunnerConfig{Mode: controlmonkey.String("managed")},
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

	created, err := stackSvc.CreateDependency(ctx, &stack.Dependency{
		StackId:          child.ID,
		DependsOnStackId: parent.ID,
		References: []*stack.DependencyRef{{
			OutputOfStackToDependOn: controlmonkey.String("output_key"),
			InputForStack:           controlmonkey.String("var_name"),
		}},
		TriggerOption: controlmonkey.String("always"),
	})
	if err != nil {
		log.Fatalf("ControlMonkey: failed to create dependency: %v", err)
	}

	depID := controlmonkey.StringValue(created.ID)
	if depID == "" {
		depID = os.Getenv("CONTROL_MONKEY_DEPENDENCY_ID")
		if depID == "" {
			log.Fatalf("ControlMonkey: dependency id is empty; set CONTROL_MONKEY_DEPENDENCY_ID to an existing dependency id")
		}
	}

	_, err = stackSvc.DeleteDependency(ctx, depID)
	if err != nil {
		log.Fatalf("ControlMonkey: failed to delete dependency %q: %v", depID, err)
	}

	log.Printf("Deleted dependency %q", depID)
}
