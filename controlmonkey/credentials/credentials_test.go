package credentials

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/featureflag"
)

type mockProvider struct {
	creds Value
}

func (m *mockProvider) Retrieve() (Value, error) {
	if m.creds.IsEmpty() {
		return m.creds, errors.New("controlmonkey: invalid credentials")
	}

	return m.creds, nil
}

func (m *mockProvider) String() string { return "mock" }

func TestChainCredentials(t *testing.T) {
	tests := map[string]struct {
		providers []Provider
		features  string
		want      Value
		err       error
	}{
		"no_providers": {
			providers: []Provider{},
			err:       ErrNoValidProvidersFoundInChain,
		},
		"single_provider_valid": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token: "token",
					},
				},
			},
			want: Value{
				Token: "token",
			},
		},
		"single_provider_invalid": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token: "",
					},
				},
			},
			err: errorList{
				errors.New("controlmonkey: invalid credentials"),
			},
		},
		"multiple_providers_valid": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token: "token1",
					},
				},
				&mockProvider{
					creds: Value{
						Token: "token2",
					},
				},
			},
			want: Value{
				Token: "token1",
			},
		},
		"multiple_providers_invalid": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token: "",
					},
				},
				&mockProvider{
					creds: Value{
						Token: "",
					},
				},
			},
			err: errorList{
				errors.New("controlmonkey: invalid credentials"),
				errors.New("controlmonkey: invalid credentials"),
			},
		},
		"providers_first_no_token": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token: "",
					},
				},
				&mockProvider{
					creds: Value{
						Token: "token2",
					},
				},
			},
			features: "MergeCredentialsChain=false",
			want: Value{
				Token: "token2",
			},
		},
		"providers_first_token": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token: "token1",
					},
				},
				&mockProvider{
					creds: Value{
						Token: "token2",
					},
				},
			},
			features: "MergeCredentialsChain=false",
			want: Value{
				Token: "token1",
			},
		},
		"providers_first_no_token_with_merge": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token: "",
					},
				},
				&mockProvider{
					creds: Value{
						Token: "token2",
					},
				},
			},
			features: "MergeCredentialsChain=true",
			want: Value{
				Token: "token2",
			},
		},
		"providers_first_token_with_merge": {
			providers: []Provider{
				&mockProvider{
					creds: Value{
						Token: "token1",
					},
				},
				&mockProvider{
					creds: Value{
						Token: "token2",
					},
				},
			},
			features: "MergeCredentialsChain=true",
			want: Value{
				Token: "token1",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.features != "" {
				origFlags := featureflag.All()
				defer func() { featureflag.Set(origFlags.String()) }() // restore
				featureflag.Set(test.features)
			}
			creds, err := NewChainCredentials(test.providers...).Get()
			if err != nil {
				if test.err != nil {
					if !strings.Contains(err.Error(), test.err.Error()) {
						t.Fatalf("want: %v, got: %v", test.err, err)
					}
					return // want failure
				} else {
					t.Fatalf("want: nil, got: %v", err)
				}
			}
			if e, a := test.want, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("want: %v, got: %v", e, a)
			}
		})
	}
}

func TestFileCredentials(t *testing.T) {
	var (
		filenameINI         = filepath.Join("testdata", "credentials_ini")
		filenameJSON        = filepath.Join("testdata", "credentials_json")
		filenameINIInvalid  = filepath.Join("testdata", "credentials_ini_invalid")
		filenameJSONInvalid = filepath.Join("testdata", "credentials_json_invalid")
	)

	tests := map[string]struct {
		filename string
		profile  string
		want     Value
		err      error
	}{
		"file_not_exist": {
			filename: "file_not_exist",
			profile:  "default",
			err:      errors.New("controlmonkey: failed to load credentials file: open file_not_exist: no such file or directory"),
		},
		"invalid_ini": {
			filename: filenameINIInvalid,
			err:      errors.New("controlmonkey: failed to load credentials file: unclosed section: [profile_nam"),
		},
		"invalid_json": {
			filename: filenameJSONInvalid,
			err:      errors.New("controlmonkey: failed to load credentials file: key-value delimiter not found: {\"token"),
		},
		"profile_not_exist": {
			filename: filenameINI,
			profile:  "profile_not_exist",
			err:      errors.New("controlmonkey: failed to load credentials file: section \"profile_not_exist\" does not exist"),
		},
		"valid_ini_profile_default": {
			filename: filenameINI,
			want: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "default_token",
			},
		},
		"valid_ini_profile_complete_credentials": {
			filename: filenameINI,
			profile:  "complete_credentials",
			want: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "complete_credentials_token",
			},
		},
		"valid_json": {
			filename: filenameJSON,
			want: Value{
				ProviderName: FileCredentialsProviderName,
				Token:        "token",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			creds, err := NewFileCredentials(test.profile, test.filename).Get()
			if err != nil {
				if test.err != nil {
					if !strings.Contains(err.Error(), test.err.Error()) {
						t.Fatalf("want: %v, got: %v", test.err, err)
					}
					return // want failure
				} else {
					t.Fatalf("want: nil, got: %v", err)
				}
			}
			if e, a := test.want, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("want: %v, got: %v", e, a)
			}
		})
	}
}

func TestEnvCredentials(t *testing.T) {
	origEnv := os.Environ()
	defer func() { // restore env
		os.Clearenv()

		for _, kv := range origEnv {
			p := strings.SplitN(kv, "=", 2)
			k, v := p[0], ""
			if len(p) > 1 {
				v = p[1]
			}
			os.Setenv(k, v)
		}
	}()

	tests := map[string]struct {
		env  map[string]string
		want Value
		err  error
	}{
		"no_token": {
			env: map[string]string{},
			err: errors.New("controlmonkey: CONTROL_MONKEY_TOKEN not found in environment"),
		},
		"with_token": {
			env: map[string]string{
				"CONTROL_MONKEY_TOKEN": "token",
			},
			want: Value{
				ProviderName: EnvCredentialsProviderName,
				Token:        "token",
			},
		},
		"all_variables": {
			env: map[string]string{
				"CONTROL_MONKEY_TOKEN": "token",
			},
			want: Value{
				ProviderName: EnvCredentialsProviderName,
				Token:        "token",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			os.Clearenv()
			for k, v := range test.env {
				os.Setenv(k, v)
			}

			creds, err := NewEnvCredentials().Get()
			if err != nil {
				if test.err != nil {
					if !strings.Contains(err.Error(), test.err.Error()) {
						t.Fatalf("want: %v, got: %v", test.err, err)
					}
					return // want failure
				} else {
					t.Fatalf("want: nil, got: %v", err)
				}
			}
			if e, a := test.want, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("want: %v, got: %v", e, a)
			}
		})
	}
}

func TestStaticCredentials(t *testing.T) {
	tests := map[string]struct {
		token   string
		account string
		want    Value
		err     error
	}{
		"empty_credentials": {
			token: "",
			err:   errors.New("controlmonkey: static credentials are empty"),
		},
		"full_credentials": {
			token: "token",
			want: Value{
				ProviderName: StaticCredentialsProviderName,
				Token:        "token",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			creds, err := NewStaticCredentials(test.token).Get()
			if err != nil {
				if test.err != nil {
					if !strings.Contains(err.Error(), test.err.Error()) {
						t.Fatalf("want: %v, got: %v", test.err, err)
					}
					return // want failure
				} else {
					t.Fatalf("want: nil, got: %v", err)
				}
			}
			if e, a := test.want, creds; !reflect.DeepEqual(e, a) {
				t.Errorf("want: %v, got: %v", e, a)
			}
		})
	}
}
