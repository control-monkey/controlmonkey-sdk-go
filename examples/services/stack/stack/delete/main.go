package main

import (
	"context"
	"log"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/stack"
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
	svc := stack.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Delete stack.
	stackId := "stk-x6ny8qj0n9"
	_, err := svc.DeleteStack(ctx, stackId)
	if err != nil {
		log.Fatalf("Control Monkey: failed to delete stack: %v", err)
	}

	log.Println("Stack was deleted successfully")
}
