package main

import (
	"context"
	"log"
	"os"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack_discovery_configuration"
)

func main() {
	sess := session.New()
	sdc := stack_discovery_configuration.New(sess)
	ctx := context.Background()

	configID := os.Getenv("CONTROL_MONKEY_STACK_DISCOVERY_CONFIG_ID")
	if configID == "" {
		log.Fatalf("ControlMonkey: set CONTROL_MONKEY_STACK_DISCOVERY_CONFIG_ID to an existing configuration id")
	}

	_, err := sdc.DeleteStackDiscoveryConfiguration(ctx, configID)
	if err != nil {
		log.Fatalf("ControlMonkey: failed to delete stack discovery configuration %q: %v", configID, err)
	}

	log.Printf("Deleted stack discovery configuration %q", configID)
}
