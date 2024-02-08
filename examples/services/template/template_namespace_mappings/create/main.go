package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/services/template"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
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

	t := &template.TemplateNamespaceMapping{
		TemplateId:  controlmonkey.String("tmpl-26gne1x95w"),
		NamespaceId: controlmonkey.String("ns-7x4pqdfdt1"),
	}

	_, err := svc.CreateTemplateNamespaceMapping(ctx, t)
	if err != nil {
		log.Fatalf("Control Monkey: failed to create template namespace mapping: %v", err)
	}

	log.Println("Template namespace mapping was created successfully")
}
