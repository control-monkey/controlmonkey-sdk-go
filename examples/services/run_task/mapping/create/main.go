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

	m := &run_task.RunTaskMapping{
		RunTaskId:        controlmonkey.String("rtsk-123"),
		TargetId:         controlmonkey.String("ns-123"),
		TargetType:       controlmonkey.String(commons.NamespaceTargetType),
		EnforcementLevel: controlmonkey.String(commons.Warning),
		Stage:            controlmonkey.String(commons.PostPlanStage),
	}

	out, err := svc.CreateRunTaskMapping(ctx, m)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create run task mapping: %v", err)
	}

	if out != nil {
		log.Printf("Run task mapping: %s", stringutil.Stringify(out))
	}
}
