package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"
	"github.com/control-monkey/controlmonkey-sdk-go/services/namespace"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack_discovery_configuration"
)

func main() {
	sess := session.New()
	nsSvc := namespace.New(sess)
	sdc := stack_discovery_configuration.New(sess)
	ctx := context.Background()

	ns, err := nsSvc.CreateNamespace(ctx, &namespace.Namespace{
		Name:        controlmonkey.String("example-discovery-ns"),
		Description: controlmonkey.String("namespace for stack discovery examples"),
	})
	if err != nil {
		log.Fatalf("ControlMonkey: failed to create namespace: %v", err)
	}

	input := &stack_discovery_configuration.StackDiscoveryConfiguration{
		Name:        controlmonkey.String("My Stack Discovery Config"),
		NamespaceId: ns.ID,
		Description: controlmonkey.String("Configuration for discovering Terraform stacks in repositories"),
		VcsPatterns: []*stack_discovery_configuration.VcsPattern{
			{
				ProviderId: controlmonkey.String("github-provider-123"),
				RepoName:   controlmonkey.String("my-terraform-repo"),
				PathPatterns: []*string{
					controlmonkey.String("environments/*/terraform/*.tf"),
					controlmonkey.String("modules/**/*.tf"),
				},
				ExcludePathPatterns: []*string{
					controlmonkey.String("**/.terraform/**"),
					controlmonkey.String("**/terraform.tfstate"),
				},
				Branch: controlmonkey.String("main"),
			},
		},
		StackConfig: &stack_discovery_configuration.StackConfig{
			IacType: controlmonkey.String("terraform"),
			DeploymentBehavior: &cross_models.DeploymentBehavior{
				DeployOnPush:    controlmonkey.Bool(false),
				WaitForApproval: controlmonkey.Bool(true),
			},
			DeploymentApprovalPolicy: &cross_models.DeploymentApprovalPolicy{
				Rules: []*cross_models.DeploymentApprovalPolicyRule{
					{Type: controlmonkey.String("requireApproval")},
				},
			},
			RunTrigger: &cross_models.RunTrigger{
				Patterns: controlmonkey.StringSlice("*"),
			},
			IacConfig: &cross_models.IacConfig{
				TerraformVersion: controlmonkey.String("1.4.5"),
			},
			RunnerConfig: &cross_models.RunnerConfig{
				Mode: controlmonkey.String("managed"),
			},
			AutoSync: &cross_models.AutoSync{
				DeployWhenDriftDetected: controlmonkey.Bool(true),
			},
		},
	}

	out, err := sdc.CreateStackDiscoveryConfiguration(ctx, input)
	if err != nil {
		log.Fatalf("ControlMonkey: failed to create stack discovery configuration: %v", err)
	}

	if out != nil {
		log.Printf("Stack Discovery Configuration %q: %s", controlmonkey.StringValue(out.ID), stringutil.Stringify(out))
	}
}
