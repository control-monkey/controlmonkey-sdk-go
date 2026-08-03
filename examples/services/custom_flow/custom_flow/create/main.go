package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/custom_flow"
)

const flowYaml = `version: 1
customFlows:
  run:
    steps:
      terraformInit:
        before:
          - name: say hello
            cmd: echo "hello from custom flow"
            failureBehavior: continue
`

func main() {
	sess := session.New()
	svc := custom_flow.New(sess)
	ctx := context.Background()

	f := &custom_flow.CustomFlow{
		Name:      controlmonkey.String("My Custom Flow"),
		FlowYaml:  controlmonkey.String(flowYaml),
		IsEnabled: controlmonkey.Bool(true),
	}

	out, err := svc.CreateCustomFlow(ctx, f)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create custom flow: %v", err)
	}

	if out != nil {
		log.Printf("Custom flow %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
