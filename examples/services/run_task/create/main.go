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

	t := &run_task.RunTask{
		Name:      controlmonkey.String("My Run Task"),
		Url:       controlmonkey.String("https://my-run-task-endpoint.example.com/callback"),
		IsEnabled: controlmonkey.Bool(true),
		HmacKey:   controlmonkey.String("my-secret-hmac-key"),
	}

	out, err := svc.CreateRunTask(ctx, t)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create run task: %v", err)
	}

	if out != nil {
		log.Printf("Run task %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
