package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"github.com/control-monkey/controlmonkey-sdk-go/services/custom_flow"
)

func main() {
	sess := session.New()
	svc := custom_flow.New(sess)
	ctx := context.Background()

	m := &custom_flow.CustomFlowMapping{
		CustomFlowId: controlmonkey.String("cflw-123"),
		TargetId:     controlmonkey.String("ns-123"),
		TargetType:   controlmonkey.String(commons.NamespaceTargetType),
	}

	if _, err := svc.DeleteCustomFlowMapping(ctx, m); err != nil {
		log.Fatalf("Control Monkey: failed to delete custom flow mapping: %v", err)
	}

	log.Printf("Custom flow mapping deleted")
}
