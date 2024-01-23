package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
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

	// Delete template.
	templateId := "tmpl-6qu7mt9k8u"
	_, err := svc.DeleteTemplate(ctx, templateId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete template: %v", err)
	}

	log.Println("Template was deleted successfully")
}
