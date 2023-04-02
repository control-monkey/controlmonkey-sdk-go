package credentials

import (
	"fmt"
	"os"
)

const (
	// EnvCredentialsProviderName specifies the name of the Env provider.
	EnvCredentialsProviderName = "EnvCredentialsProvider"

	// EnvCredentialsVarToken specifies the name of the environment variable
	// points to the ControlMonkey Token.
	EnvCredentialsVarToken = "CONTROL_MONKEY_TOKEN"
)

// ErrEnvCredentialsNotFound is returned when no credentials can be found in the
// process's environment.
var ErrEnvCredentialsNotFound = fmt.Errorf("controlmonkey: %s not found in environment", EnvCredentialsVarToken)

// A EnvProvider retrieves credentials from the environment variables of the
// running process.
//
// Environment variables used:
// * Token   : CONTROL_MONKEY_TOKEN
type EnvProvider struct{}

// NewEnvCredentials returns a pointer to a new Credentials object wrapping the
// environment variable provider.
func NewEnvCredentials() *Credentials {
	return NewCredentials(&EnvProvider{})
}

// Retrieve retrieves the keys from the environment.
func (e *EnvProvider) Retrieve() (Value, error) {
	value := Value{
		Token:        os.Getenv(EnvCredentialsVarToken),
		ProviderName: EnvCredentialsProviderName,
	}

	if value.IsEmpty() {
		return value, ErrEnvCredentialsNotFound
	}

	return value, nil
}

// String returns the string representation of the provider.
func (e *EnvProvider) String() string { return EnvCredentialsProviderName }
