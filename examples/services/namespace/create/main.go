package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"
	"github.com/control-monkey/controlmonkey-sdk-go/services/namespace"
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
	svc := namespace.New(sess)

	// Create a new context.
	ctx := context.Background()

	externalCredentials1 := namespace.ExternalCredentials{
		Type:                  controlmonkey.String("awsAssumeRole"),
		ExternalCredentialsId: controlmonkey.String("ext-stage"),
		AwsProfileName:        controlmonkey.String("stage"),
	}
	externalCredentials2 := namespace.ExternalCredentials{
		Type:                  controlmonkey.String("awsAssumeRole"),
		ExternalCredentialsId: controlmonkey.String("ext-dev"),
		AwsProfileName:        controlmonkey.String("dev"),
	}
	externalCredentials3 := namespace.ExternalCredentials{
		Type:                  controlmonkey.String("gcpServiceAccount"),
		ExternalCredentialsId: controlmonkey.String("ext-gcp"),
	}
	externalCredentials := []*namespace.ExternalCredentials{&externalCredentials1, &externalCredentials2, &externalCredentials3}

	n := namespace.Namespace{
		Name:                controlmonkey.String("go namespace"),
		Description:         controlmonkey.String("description"),
		ExternalCredentials: externalCredentials,
		IacConfig: &namespace.IacConfig{
			TerraformVersion:  controlmonkey.String("1.5.0"),
			TerragruntVersion: controlmonkey.String("0.39.0"),
		},
		RunnerConfig: &namespace.RunnerConfig{
			Mode:          controlmonkey.String("selfHosted"),
			Groups:        []*string{controlmonkey.String("default")},
			IsOverridable: controlmonkey.Bool(true),
		},
		DeploymentApprovalPolicy: &namespace.DeploymentApprovalPolicy{
			Rules: []*cross_models.DeploymentApprovalPolicyRule{
				{
					Type: controlmonkey.String("requireTwoApprovals"),
				},
				{
					Type: controlmonkey.String("requireTeamsApproval"),
					Parameters: &map[string]interface{}{
						"teams": controlmonkey.StringSlice("team-123"),
					},
				},
			},
			OverrideBehavior: controlmonkey.String("allow"),
		},
	}
	// Create namespace.
	out, err := svc.CreateNamespace(ctx, &n)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create namespace: %v", err)
	}

	// Output namespace, if was created.
	if out != nil {
		log.Printf("Namespace %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
