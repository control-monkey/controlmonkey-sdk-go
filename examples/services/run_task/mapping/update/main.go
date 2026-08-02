package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"github.com/control-monkey/controlmonkey-sdk-go/services/run_task"
)

func main() {
	sess := session.New()
	svc := run_task.New(sess)
	ctx := context.Background()

	// The mapping is identified by runTaskId + targetId + targetType. Those three are required on
	// update; enforcementLevel and stage are the mutable fields.
	m := &run_task.RunTaskMapping{
		RunTaskId:        controlmonkey.String("rtsk-123"),
		TargetId:         controlmonkey.String("ns-123"),
		TargetType:       controlmonkey.String(commons.NamespaceTargetType),
		EnforcementLevel: controlmonkey.String(commons.HardMandatory),
	}

	out, err := svc.UpdateRunTaskMapping(ctx, m)
	if err != nil {
		log.Fatalf("Control Monkey: failed to update run task mapping: %v", err)
	}

	if out != nil {
		log.Printf("Run task mapping: %s", stringutil.Stringify(out))
	}
}
