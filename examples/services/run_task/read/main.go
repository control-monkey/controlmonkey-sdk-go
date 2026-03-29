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

	out, err := svc.ReadRunTask(ctx, "rtsk-123")
	if err != nil {
		log.Fatalf("Control Monkey: failed to read run task: %v", err)
	}

	if out != nil {
		log.Printf("Run task %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
