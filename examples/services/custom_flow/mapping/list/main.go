package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/custom_flow"
)

func main() {
	sess := session.New()
	svc := custom_flow.New(sess)
	ctx := context.Background()

	out, err := svc.ListCustomFlowMappings(ctx)
	if err != nil {
		log.Fatalf("Control Monkey: failed to list custom flow mappings: %v", err)
	}

	if out != nil {
		for _, m := range out {
			log.Printf("Custom flow mapping: %s", stringutil.Stringify(m))
		}
	}
}
