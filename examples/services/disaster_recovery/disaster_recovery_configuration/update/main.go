package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/disaster_recovery"

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

	// Update a new instance of the services's client with a Session.
	// Optional controlmonkey.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// services specific configuration.
	svc := disaster_recovery.New(sess)

	// Update a new context.
	ctx := context.Background()

	group1 := &map[string]interface{}{
		"vcsInfo": map[string]interface{}{
			"path": "bc/region/services/a/b/c",
		},
		"awsQuery": map[string]interface{}{
			"tags": []map[string]interface{}{
				{
					"key":   "Service",
					"value": "aaa",
				}},
			"region": "us-west-2",
		},
	}

	group2 := &map[string]interface{}{
		"vcsInfo": map[string]interface{}{
			"path": "bc/region/services",
		},
		"awsQuery": map[string]interface{}{
			"tags": []map[string]interface{}{
				{
					"key":   "Service",
					"value": "Room",
				},
			},
			"services":      []string{"AWS::EC2"},
			"resourceTypes": []string{"AWS::EC2::Instance"},
			"region":        "us-west-2",
		},
	}

	group3Str := `
 {
  "vcsInfo": {
    "path": "a/b/c"
  },
  "awsQuery": {
    "tags": [
      {
        "key": "Owner",
        "value": "Hi"
      }
    ],
    "region": "us-east-1",
    "services": [
      "AWS::S3"
    ],
    "resourceTypes": [
      "AWS::S3::Bucket"
    ]
  }
}
`

	var group3 *map[string]interface{}

	if err := json.Unmarshal([]byte(group3Str), &group3); err != nil {
		panic(err)
	}

	bs := &disaster_recovery.BackupStrategy{
		IncludeManagedResources: controlmonkey.Bool(false),
		Mode:                    controlmonkey.String("default"),
		VcsInfo: &disaster_recovery.VcsInfo{
			ProviderId: controlmonkey.String("vcsp-123"),
			RepoName:   controlmonkey.String("terraform"),
			Branch:     controlmonkey.String("main"),
		},
		Groups: []*map[string]interface{}{
			group1,
			group2,
			group3,
		},
	}

	s := &disaster_recovery.DisasterRecoveryConfiguration{
		Scope:          controlmonkey.String("aws"),
		CloudAccountId: controlmonkey.String("123123123"),
		BackupStrategy: bs,
	}
	// Update disaster_recovery.
	out, err := svc.UpdateDisasterRecoveryConfiguration(ctx, "drc-123", s)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create entity: %v", err)
	}

	// Output entity, if was created.
	if out != nil {
		log.Printf("Entity %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
