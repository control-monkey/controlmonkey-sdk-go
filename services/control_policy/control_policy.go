package control_policy

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region ControlPolicyMapping

// region Structure

type ControlPolicyMapping struct {
	ControlPolicyId  *string `json:"controlPolicyId,omitempty"`
	TargetId         *string `json:"targetId,omitempty"`
	TargetType       *string `json:"targetType,omitempty"`       //commons.PolicyMappingTargetTypes
	EnforcementLevel *string `json:"enforcementLevel,omitempty"` //commons.EnforcementLevelTypes

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateControlPolicyMapping(ctx context.Context, input *ControlPolicyMapping) (*ControlPolicyMapping, error) {
	r := client.NewRequest(http.MethodPost, "/controlPolicy/controlPolicyMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	controlPolicyMapping, err := controlPolicyMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicyMapping)
	if len(controlPolicyMapping) > 0 {
		output = controlPolicyMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) ListControlPolicyMappings(ctx context.Context, controlPolicyId string) ([]*ControlPolicyMapping, error) {
	r := client.NewRequest(http.MethodGet, "/controlPolicy/controlPolicyMapping")
	r.Params.Set("controlPolicyId", controlPolicyId)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := controlPolicyMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) UpdateControlPolicyMapping(ctx context.Context, input *ControlPolicyMapping) (*ControlPolicyMapping, error) {
	r := client.NewRequest(http.MethodPut, "/controlPolicy/controlPolicyMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	controlPolicyMapping, err := controlPolicyMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicyMapping)

	if len(controlPolicyMapping) > 0 {
		output = controlPolicyMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteControlPolicyMapping(ctx context.Context, input *ControlPolicyMapping) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/controlPolicy/controlPolicyMapping")
	r.Obj = input

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

func controlPolicyMappingFromJSON(in []byte) (*ControlPolicyMapping, error) {
	b := new(ControlPolicyMapping)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func controlPolicyMappingsFromJSON(in []byte) ([]*ControlPolicyMapping, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ControlPolicyMapping, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := controlPolicyMappingFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func controlPolicyMappingsFromHttpResponse(resp *http.Response) ([]*ControlPolicyMapping, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return controlPolicyMappingsFromJSON(body)
}

//endregion

//region Setters

func (o ControlPolicyMapping) MarshalJSON() ([]byte, error) {
	type noMethod ControlPolicyMapping
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ControlPolicyMapping) SetControlPolicyId(v *string) *ControlPolicyMapping {
	if o.ControlPolicyId = v; o.ControlPolicyId == nil {
		o.nullFields = append(o.nullFields, "ControlPolicyId")
	}
	return o
}

func (o *ControlPolicyMapping) SetTargetId(v *string) *ControlPolicyMapping {
	if o.TargetId = v; o.TargetId == nil {
		o.nullFields = append(o.nullFields, "TargetId")
	}
	return o
}

func (o *ControlPolicyMapping) SetTargetType(v *string) *ControlPolicyMapping {
	if o.TargetType = v; o.TargetType == nil {
		o.nullFields = append(o.nullFields, "TargetType")
	}
	return o
}

func (o *ControlPolicyMapping) SetEnforcementLevel(v *string) *ControlPolicyMapping {
	if o.EnforcementLevel = v; o.EnforcementLevel == nil {
		o.nullFields = append(o.nullFields, "EnforcementLevel")
	}
	return o
}

//endregion

//endregion
