package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/run_task"
)

func main() {
	sess := session.New()
	svc := run_task.New(sess)
	ctx := context.Background()

	_, err := svc.DeleteRunTask(ctx, "rtsk-123")
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete run task: %v", err)
	}

	log.Printf("Run task deleted successfully")
}
