package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/template"
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

	// Read template.
	templateId := "tmpl-6qu7mt9k8u"
	t, err := svc.ReadTemplate(ctx, templateId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to read template: %v", err)
	}

	// Output template, if exists.
	if t != nil {
		log.Printf("Template %q: %s",
			controlmonkey.StringValue(t.ID),
			stringutil.Stringify(t))
	}
}
