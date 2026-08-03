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

	out, err := svc.ReadCustomFlow(ctx, "cflw-123")
	if err != nil {
		log.Fatalf("Control Monkey: failed to read custom flow: %v", err)
	}

	if out != nil {
		log.Printf("Custom flow %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
