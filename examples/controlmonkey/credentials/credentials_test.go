package main_test

import (
	"fmt"
	"os"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/credentials"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/featureflag"
)

func Example_static() {
	// Initialize a new static credentials provider.
	provider := credentials.NewStaticCredentials("secret")

	value, err := provider.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(value.Token)
	// Output: secret
}

func Example_env() {
	// Set token.
	//
	// Can be set using an environment variables as well, for example:
	// export CONTROL_MONKEY_TOKEN=secret
	os.Setenv(credentials.EnvCredentialsVarToken, "secret")

	// Unset.
	defer func() {
		os.Unsetenv(credentials.EnvCredentialsVarToken)
	}()

	// Initialize a new env credentials provider.
	provider := credentials.NewEnvCredentials()

	value, err := provider.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(value.Token)
	// Output: secret
}

func Example_chainAllowPartial() {
	// Set token.
	os.Setenv(credentials.EnvCredentialsVarToken, "secret")

	// Unset.
	defer func() {
		os.Unsetenv(credentials.EnvCredentialsVarToken)
	}()

	// Disable the usage of merging credentials in chain provider.
	//
	// Can be set using an environment variable as well, for example:
	// export CONTROL_MONKEY_FEATURE_FLAGS=MergeCredentialsChain=false
	featureflag.Set("MergeCredentialsChain=false")

	// Initialize a new chain credentials provider.
	provider := credentials.NewChainCredentials(
		&credentials.EnvProvider{},
		&credentials.StaticProvider{
			Value: credentials.Value{},
		},
	)

	value, err := provider.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(value.Token)
	// Output: secret
}

func Example_chainAllowMerge() {
	os.Setenv(credentials.EnvCredentialsVarToken, "secret")
	defer func() { os.Unsetenv(credentials.EnvCredentialsVarToken) }()

	// Enable the usage of merging credentials in chain provider.
	//
	// Can be set using an environment variable as well, for example:
	// export CONTROL_MONKEY_FEATURE_FLAGS=MergeCredentialsChain=true
	featureflag.Set("MergeCredentialsChain=true")

	// Initialize a new chain credentials provider.
	provider := credentials.NewChainCredentials(
		&credentials.EnvProvider{},
		&credentials.StaticProvider{
			Value: credentials.Value{},
		},
	)

	value, err := provider.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(value.Token)
	// Output: secret
}
