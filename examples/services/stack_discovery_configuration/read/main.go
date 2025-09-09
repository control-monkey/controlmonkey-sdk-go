package main

import (
	"context"
	"log"
	"os"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
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

	out, err := sdc.ReadStackDiscoveryConfiguration(ctx, configID)
	if err != nil {
		log.Fatalf("ControlMonkey: failed to read stack discovery configuration %q: %v", configID, err)
	}

	if out != nil {
		log.Printf("Stack Discovery Configuration %q: %s", controlmonkey.StringValue(out.ID), stringutil.Stringify(out))
	}
}
