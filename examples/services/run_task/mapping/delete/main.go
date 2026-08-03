package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"github.com/control-monkey/controlmonkey-sdk-go/services/run_task"
)

func main() {
	sess := session.New()
	svc := run_task.New(sess)
	ctx := context.Background()

	// Only the identity triple may be sent on delete. The API rejects enforcementLevel and stage.
	m := &run_task.RunTaskMapping{
		RunTaskId:  controlmonkey.String("rtsk-123"),
		TargetId:   controlmonkey.String("ns-123"),
		TargetType: controlmonkey.String(commons.NamespaceTargetType),
	}

	if _, err := svc.DeleteRunTaskMapping(ctx, m); err != nil {
		log.Fatalf("Control Monkey: failed to delete run task mapping: %v", err)
	}

	log.Printf("Run task mapping deleted")
}
