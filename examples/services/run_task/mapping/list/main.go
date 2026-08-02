package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/run_task"
)

func main() {
	sess := session.New()
	svc := run_task.New(sess)
	ctx := context.Background()

	out, err := svc.ListRunTaskMappings(ctx, "rtsk-123")
	if err != nil {
		log.Fatalf("Control Monkey: failed to list run task mappings: %v", err)
	}

	if out != nil {
		for _, m := range out {
			log.Printf("Run task mapping: %s", stringutil.Stringify(m))
		}
	}
}
