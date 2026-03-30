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

	t := new(run_task.RunTask)
	t.SetName(controlmonkey.String("My Run Task Updated"))
	t.SetIsEnabled(controlmonkey.Bool(false))

	out, err := svc.UpdateRunTask(ctx, "rtsk-123", t)
	if err != nil {
		log.Fatalf("Control Monkey: failed to update run task: %v", err)
	}

	if out != nil {
		log.Printf("Run task %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
