package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack_discovery_configuration"
)

func main() {
	sess := session.New()
	sdc := stack_discovery_configuration.New(sess)
	ctx := context.Background()

	configID := "sdc-123"
	update := &stack_discovery_configuration.StackDiscoveryConfiguration{
		Description: controlmonkey.String("Updated configuration for discovering Terraform stacks"),
		StackConfig: &stack_discovery_configuration.StackConfig{
			DeploymentBehavior: &cross_models.DeploymentBehavior{
				DeployOnPush: controlmonkey.Bool(true), // Changed from false to true
			},
		},
	}

	out, err := sdc.UpdateStackDiscoveryConfiguration(ctx, configID, update)
	if err != nil {
		log.Fatalf("ControlMonkey: failed to update stack discovery configuration %q: %v", configID, err)
	}

	if out != nil {
		log.Printf("Stack Discovery Configuration %q: %s", controlmonkey.StringValue(out.ID), stringutil.Stringify(out))
	}
}
