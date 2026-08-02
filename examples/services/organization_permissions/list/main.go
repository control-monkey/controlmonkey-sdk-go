package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/organization_permissions"
)

func main() {
	sess := session.New()
	svc := organization_permissions.New(sess)
	ctx := context.Background()

	// teamId is required by the API.
	out, err := svc.ListOrganizationPermissions(ctx, "team-123")
	if err != nil {
		log.Fatalf("Control Monkey: failed to list organization permissions: %v", err)
	}

	if out != nil {
		for _, p := range out {
			log.Printf("Organization permission: %s", stringutil.Stringify(p))
		}
	}
}
