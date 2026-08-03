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

	// Partial update: flowYaml is left untouched when it is not sent.
	f := &custom_flow.CustomFlow{
		Name: controlmonkey.String("My Custom Flow - renamed"),
	}

	out, err := svc.UpdateCustomFlow(ctx, "cflw-123", f)
	if err != nil {
		log.Fatalf("Control Monkey: failed to update custom flow: %v", err)
	}

	if out != nil {
		log.Printf("Custom flow %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
