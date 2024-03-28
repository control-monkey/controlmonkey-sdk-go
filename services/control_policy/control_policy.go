package control_policy

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region ControlPolicy

// region Structure

type ControlPolicy struct {
	ID          *string                 `json:"id,omitempty"` // read-only
	Name        *string                 `json:"name,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Type        *string                 `json:"type,omitempty"`
	Parameters  *map[string]interface{} `json:"parameters,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateControlPolicy(ctx context.Context, input *ControlPolicy) (*ControlPolicy, error) {
	r := client.NewRequest(http.MethodPost, "/controlPolicy")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	controlPolicy, err := controlPoliciesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicy)
	if len(controlPolicy) > 0 {
		output = controlPolicy[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadControlPolicy(ctx context.Context, controlPolicyId string) (*ControlPolicy, error) {
	path, err := uritemplates.Expand("/controlPolicy/{controlPolicyId}", uritemplates.Values{"controlPolicyId": controlPolicyId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	controlPolicy, err := controlPoliciesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicy)
	if len(controlPolicy) > 0 {
		output = controlPolicy[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateControlPolicy(ctx context.Context, id string, input *ControlPolicy) (*ControlPolicy, error) {
	path, err := uritemplates.Expand("/controlPolicy/{controlPolicyId}", uritemplates.Values{"controlPolicyId": id})
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

	controlPolicy, err := controlPoliciesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicy)

	if len(controlPolicy) > 0 {
		output = controlPolicy[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteControlPolicy(ctx context.Context, id string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/controlPolicy/{controlPolicyId}", uritemplates.Values{"controlPolicyId": id})
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

func controlPolicyFromJSON(in []byte) (*ControlPolicy, error) {
	b := new(ControlPolicy)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func controlPolicysFromJSON(in []byte) ([]*ControlPolicy, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ControlPolicy, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := controlPolicyFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func controlPoliciesFromHttpResponse(resp *http.Response) ([]*ControlPolicy, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return controlPolicysFromJSON(body)
}

//endregion

//region Setters

func (o ControlPolicy) MarshalJSON() ([]byte, error) {
	type noMethod ControlPolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ControlPolicy) SetID(v *string) *ControlPolicy {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *ControlPolicy) SetName(v *string) *ControlPolicy {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *ControlPolicy) SetDescription(v *string) *ControlPolicy {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *ControlPolicy) SetType(v *string) *ControlPolicy {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *ControlPolicy) SetParameters(v *map[string]interface{}) *ControlPolicy {
	if o.Parameters = v; o.Parameters == nil {
		o.nullFields = append(o.nullFields, "Parameters")
	}
	return o
}

//endregion

//endregion
