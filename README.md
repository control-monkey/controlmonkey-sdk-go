# ControlMonkey SDK Go

The official ControlMonkey SDK for the Go programming language.

## Table of Contents

- [Installation](#installation)
- [Authentication](#authentication)
- [Complete SDK Example](#complete-sdk-example)
- [Documentation](#documentation)
- [Examples](#examples)
- [Getting Help](#getting-help)
- [Community](#community)
- [Contributing](#contributing)
- [License](#license)

## Installation

The best way to get started working with the SDK is to use go get to add the SDK to your Go application using Go modules.

```
go get -u github.com/controlmonkey/controlmonkey-sdk-go
```

Without Go Modules, or in a GOPATH with Go 1.11 or 1.12 use the /... suffix on the go get to retrieve all of the SDK's dependencies.

```
go get -u github.com/controlmonkey/controlmonkey-sdk-go/...
```

## Authentication

Set a `ChainProvider` that will search for a provider which returns credentials.

The `ChainProvider` provides a way of chaining multiple providers together
which will pick the first available using priority order of the Providers
in the list. If none of the Providers retrieve valid credentials, `ChainProvider`'s
`Retrieve()` will return the error `ErrNoValidProvidersFoundInChain`. If a Provider
is found which returns valid credentials `ChainProvider` will cache that Provider
for all calls until `Retrieve` is called again.

Example of `ChainProvider` to be used with an `EnvCredentialsProvider` and
`FileCredentialsProvider`. In this example `EnvProvider` will first check if
any credentials are available via the CONTROL_MONKEY_TOKEN environment variable. If there are
none `ChainProvider` will check the next `Provider` in the list, `FileProvider`
in this case. If `FileCredentialsProvider` does not return any credentials
`ChainProvider` will return the error `ErrNoValidProvidersFoundInChain`.

```go
// Initial credentials loaded from SDK's default credential chain. Such as
// the environment, shared credentials (~/.controlmonkey/credentials), etc.
sess := session.New()

// Create the chain credentials.
creds := credentials.NewChainCredentials(
    new(credentials.FileProvider),
    new(credentials.EnvProvider),
)

// Create service client value configured for credentials
// from the chain.
svc := variable.New(sess, &controlmonkey.Config{Credentials: creds})
```

## Complete SDK Example

```go
package main

import (
	"context"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/stringutil"
	"github.com/control-monkey/controlmonkey-sdk-go/service/variable"
	"log"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the service's client with a Session.
	// Optional controlmonkey.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := variable.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read variable.
	out, err := svc.ReadVariable(ctx, &variable.ReadVariableInput{VariableId: controlmonkey.String("var-123")})
	if err != nil {
		log.Fatalf("Control Monkey: failed to read variable: %v", err)
	}

	// Output variable, if exists.
	if out.Variable != nil {
		log.Printf("Variable %q: %s",
			controlmonkey.StringValue(out.Variable.ID),
			stringutil.Stringify(out.Variable))
	}
}

```

## Documentation

For a comprehensive documentation, check out the [Control Monkey Documentation](https://docs.controlmonkey.io/) website.

## Examples

For a list of examples, check out the [examples](/examples) directory.

## Getting Help

We use GitHub issues for tracking bugs and feature requests. Please use these community resources for getting help:

- Ask a question on [Stack Overflow](https://stackoverflow.com/) and tag it with [controlmonkey-sdk-go](https://stackoverflow.com/questions/tagged/controlmonkey-sdk-go/).
- Open an [issue](https://github.com/controlmonkey/controlmonkey-sdk-go/issues/new/choose/).

## Contributing

Please see the [contribution guidelines](.github/CONTRIBUTING.md).

## License

Code is licensed under the [Apache License 2.0](LICENSE). See [NOTICE.md](NOTICE.md) for complete details, including software and third-party licenses and permissions.
