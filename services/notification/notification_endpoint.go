package notification

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region Endpoint

//region Structure

type Endpoint struct {
	ID       *string `json:"id,omitempty"` // read-only
	Name     *string `json:"name,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Url      *string `json:"url,omitempty"`

	// forceSendFields is a read of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a read of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this read has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateNotificationEndpoint(ctx context.Context, input *Endpoint) (*Endpoint, error) {
	r := client.NewRequest(http.MethodPost, baseUrl+endpointUrl)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	endpoint, err := endpointFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Endpoint)
	if len(endpoint) > 0 {
		output = endpoint[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadNotificationEndpoint(ctx context.Context, endpointId string) (*Endpoint, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{endpointId}", uritemplates.Values{"endpointId": endpointId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	endpoint, err := endpointFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Endpoint)
	if len(endpoint) > 0 {
		output = endpoint[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateNotificationEndpoint(ctx context.Context, endpointId string, input *Endpoint) (*Endpoint, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{endpointId}", uritemplates.Values{"endpointId": endpointId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	endpoint, err := endpointFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Endpoint)
	if len(endpoint) > 0 {
		output = endpoint[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteNotificationEndpoint(ctx context.Context, endpointId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{endpointId}", uritemplates.Values{"endpointId": endpointId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output := new(commons.EmptyResponse)
	return output, nil
}

//endregion

//region Private Methods

func endpointFromJSON(in []byte) (*Endpoint, error) {
	b := new(Endpoint)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func endpointsFromJSON(in []byte) ([]*Endpoint, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Endpoint, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := endpointFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func endpointFromHttpResponse(resp *http.Response) ([]*Endpoint, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return endpointsFromJSON(body)
}

//endregion

//region Setters

func (o Endpoint) MarshalJSON() ([]byte, error) {
	type noMethod Endpoint
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Endpoint) SetName(v *string) *Endpoint {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Endpoint) SetProtocol(v *string) *Endpoint {
	if o.Protocol = v; o.Protocol == nil {
		o.nullFields = append(o.nullFields, "Protocol")
	}
	return o
}

func (o *Endpoint) SetUrl(v *string) *Endpoint {
	if o.Url = v; o.Url == nil {
		o.nullFields = append(o.nullFields, "Url")
	}
	return o
}

//endregion

//endregion
