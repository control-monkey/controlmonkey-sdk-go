package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/custom_flow"
)

func main() {
	sess := session.New()
	svc := custom_flow.New(sess)
	ctx := context.Background()

	// The API refuses while the flow still has target mappings.
	if _, err := svc.DeleteCustomFlow(ctx, "cflw-123"); err != nil {
		log.Fatalf("Control Monkey: failed to delete custom flow: %v", err)
	}

	log.Printf("Custom flow deleted")
}
