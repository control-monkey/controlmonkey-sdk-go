package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack_discovery_configuration"
)

func main() {
	sess := session.New()
	sdc := stack_discovery_configuration.New(sess)
	ctx := context.Background()

	configID := "sdc-123"
	_, err := sdc.DeleteStackDiscoveryConfiguration(ctx, configID)
	if err != nil {
		log.Fatalf("ControlMonkey: failed to delete stack discovery configuration %q: %v", configID, err)
	}

	log.Printf("Deleted stack discovery configuration %q", configID)
}
