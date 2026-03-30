package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/run_task"
)

func main() {
	sess := session.New()
	svc := run_task.New(sess)
	ctx := context.Background()

	out, err := svc.ListRunTasks(ctx, nil, nil)
	if err != nil {
		log.Fatalf("Control Monkey: failed to list run tasks: %v", err)
	}

	if out != nil {
		for _, rt := range out {
			log.Printf("Run task %q: %s",
				controlmonkey.StringValue(rt.ID),
				stringutil.Stringify(rt))
		}
	}
}
