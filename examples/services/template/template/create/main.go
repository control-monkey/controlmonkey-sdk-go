package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/template"

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
	svc := template.New(sess)

	// Create a new context.
	ctx := context.Background()

	t := &template.Template{
		Name:        controlmonkey.String("template1"),
		IacType:     controlmonkey.String("terraform"),
		Description: controlmonkey.String("description"),
		VcsInfo: &template.VcsInfo{
			ProviderId: controlmonkey.String("vcsp-123"),
			RepoName:   controlmonkey.String("repo"),
			Path:       controlmonkey.String("path"),
			Branch:     controlmonkey.String("main"),
		},
		Policy: &template.Policy{
			TtlConfig: &template.TtlConfig{
				MaxTtl: &template.TtlDefinition{
					Type:  controlmonkey.String("days"),
					Value: controlmonkey.Int(3),
				},
				DefaultTtl: &template.TtlDefinition{
					Type:  controlmonkey.String("hours"),
					Value: controlmonkey.Int(3),
				},
			},
		},
	}

	// Create template.
	out, err := svc.CreateTemplate(ctx, t)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create template: %v", err)
	}

	// Output template, if was created.
	if out != nil {
		log.Printf("Template %q: %s",
			controlmonkey.StringValue(out.ID),
			stringutil.Stringify(out))
	}
}
