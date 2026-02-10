package external_credential

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
)

//region ExternalCredential

//region Structure

type ExternalCredential struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) ListExternalCredentials(ctx context.Context, credentialsVendor string, credentialId *string, credentialName *string) ([]*ExternalCredential, error) {
	r := client.NewRequest(http.MethodGet, "/org/externalCredentials")

	r.Params.Set("credentialsVendor", credentialsVendor)

	if credentialId != nil {
		r.Params.Set("credentialId", *credentialId)
	}
	if credentialName != nil {
		r.Params.Set("credentialName", *credentialName)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := externalCredentialsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

//endregion

//region Private Methods

func externalCredentialFromJSON(in []byte) (*ExternalCredential, error) {
	b := new(ExternalCredential)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func externalCredentialsFromJSON(in []byte) ([]*ExternalCredential, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ExternalCredential, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := externalCredentialFromJSON(rb)
		if err != nil {
			return nil, err
		}

		out[i] = b
	}
	return out, nil
}

func externalCredentialsFromHttpResponse(resp *http.Response) ([]*ExternalCredential, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return externalCredentialsFromJSON(body)
}

//endregion

//region Setters

func (o ExternalCredential) MarshalJSON() ([]byte, error) {
	type noMethod ExternalCredential
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ExternalCredential) SetID(v *string) *ExternalCredential {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *ExternalCredential) SetName(v *string) *ExternalCredential {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

//endregion

//endregion
