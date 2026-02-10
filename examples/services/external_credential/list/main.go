package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/external_credential"
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
	svc := external_credential.New(sess)

	// Create a new context.
	ctx := context.Background()

	// List external_credential.
	credential_vendor := "aws"
	credential_name := "awsProdCreds"
	out, err := svc.ListExternalCredentials(ctx, credential_vendor, nil, &credential_name)
	if err != nil {
		log.Fatalf("Control Monkey: failed to read external_credentials: %v", err)
	}

	// Output external_credentials, if exists.
	if out != nil {
		log.Printf("External Credentials: %s",
			stringutil.Stringify(out))
	}
}
