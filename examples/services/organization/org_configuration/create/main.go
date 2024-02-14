package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"github.com/control-monkey/controlmonkey-sdk-go/services/organization"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the services's client with a Session.
	// Optional controlmonkey.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// services specific configuration.
	svc := organization.New(sess)

	// Create a new context.
	ctx := context.Background()

	o := organization.OrgConfiguration{
		IacConfig: &organization.IacConfig{
			TerraformVersion: controlmonkey.String("1.5.0"),
		},
		S3StateFilesLocations: []*organization.S3StateFilesLocation{
			{
				BucketName:   controlmonkey.String("bucket1"),
				BucketRegion: controlmonkey.String("us-east-1"),
				AwsAccountId: controlmonkey.String("123456789"),
			},
		},
		SuppressedResources: &organization.SuppressedResources{
			ManagedByTags: []*organization.TagProperties{
				{
					Key: controlmonkey.String("Owner"),
				},
			},
		},
		ReportConfigurations: []*organization.ReportConfiguration{
			{
				Type: controlmonkey.String(commons.WeeklyReportType),
				Recipients: &organization.ReportRecipients{
					AllAdmins: controlmonkey.Bool(false),
					EmailAddresses: []*string{
						controlmonkey.String("example1@exmplae.com"),
						controlmonkey.String("example2@exmplae.com"),
					},
					EmailAddressesToExclude: []*string{
						controlmonkey.String("example3@exmplae.com"),
						controlmonkey.String("example4@exmplae.com"),
					},
				},
				Enabled: controlmonkey.Bool(true),
			},
		},
	}

	// Create namespace.
	out, err := svc.UpsertOrgConfiguration(ctx, &o)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create org configuration: %v", err)
	}

	// Output namespace, if was created.
	if out != nil {
		log.Printf("Org configuration: %s", stringutil.Stringify(out))
	}
}
