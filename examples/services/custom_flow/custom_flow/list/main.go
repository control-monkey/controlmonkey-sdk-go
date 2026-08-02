package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/custom_flow"
)

func main() {
	sess := session.New()
	svc := custom_flow.New(sess)
	ctx := context.Background()

	out, err := svc.ListCustomFlows(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatalf("Control Monkey: failed to list custom flows: %v", err)
	}

	if out != nil {
		for _, f := range out {
			log.Printf("Custom flow %q: %s",
				controlmonkey.StringValue(f.ID),
				stringutil.Stringify(f))
		}
	}
}
