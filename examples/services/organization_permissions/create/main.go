package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/organization_permissions"
)

func main() {
	sess := session.New()
	svc := organization_permissions.New(sess)
	ctx := context.Background()

	// Exactly one principal, and always a custom role of type organizationRole.
	p := &organization_permissions.OrganizationPermission{
		TeamId:       controlmonkey.String("team-123"),
		CustomRoleId: controlmonkey.String("cro-123"),
	}

	if _, err := svc.CreateOrganizationPermission(ctx, p); err != nil {
		log.Fatalf("Control Monkey: failed to create organization permission: %v", err)
	}

	log.Printf("Organization permission created")
}
