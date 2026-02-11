package external_credentials

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
)

//region ExternalCredentials

//region Structure

type ExternalCredentials struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) ListExternalCredentials(ctx context.Context, credentialsVendor string, credentialsId *string, credentialsName *string) ([]*ExternalCredentials, error) {
	r := client.NewRequest(http.MethodGet, "/org/externalCredentials")

	r.Params.Set("credentialsVendor", credentialsVendor)

	if credentialsId != nil {
		r.Params.Set("credentialsId", *credentialsId)
	}
	if credentialsName != nil {
		r.Params.Set("credentialsName", *credentialsName)
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

func externalCredentialsFromJSON(in []byte) (*ExternalCredentials, error) {
	b := new(ExternalCredentials)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func externalCredentialsListFromJSON(in []byte) ([]*ExternalCredentials, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ExternalCredentials, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := externalCredentialsFromJSON(rb)
		if err != nil {
			return nil, err
		}

		out[i] = b
	}
	return out, nil
}

func externalCredentialsFromHttpResponse(resp *http.Response) ([]*ExternalCredentials, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return externalCredentialsListFromJSON(body)
}

//endregion

//region Setters

func (o ExternalCredentials) MarshalJSON() ([]byte, error) {
	type noMethod ExternalCredentials
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ExternalCredentials) SetID(v *string) *ExternalCredentials {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *ExternalCredentials) SetName(v *string) *ExternalCredentials {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

//endregion

//endregion
