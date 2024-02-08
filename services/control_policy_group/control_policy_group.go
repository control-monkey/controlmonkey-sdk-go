package control_policy_group

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region ControlPolicyGroupMapping

// region Structure

type ControlPolicyGroupMapping struct {
	ControlPolicyGroupId *string                `json:"controlPolicyGroupId,omitempty"`
	TargetId             *string                `json:"targetId,omitempty"`
	TargetType           *string                `json:"targetType,omitempty"`       //commons.PolicyMappingTargetTypes
	EnforcementLevel     *string                `json:"enforcementLevel,omitempty"` //commons.GroupEnforcementLevelTypes
	OverrideEnforcements []*OverrideEnforcement `json:"overrideEnforcements,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type OverrideEnforcement struct {
	ControlPolicyId  *string   `json:"controlPolicyId,omitempty"`
	EnforcementLevel *string   `json:"enforcementLevel,omitempty"`
	StackIds         []*string `json:"stackIds,omitempty"` //commons.EnforcementLevelTypes

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateControlPolicyGroupMapping(ctx context.Context, input *ControlPolicyGroupMapping) (*ControlPolicyGroupMapping, error) {
	r := client.NewRequest(http.MethodPost, "/controlPolicyGroup/controlPolicyGroupMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	controlPolicyGroupMapping, err := controlPolicyGroupMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicyGroupMapping)
	if len(controlPolicyGroupMapping) > 0 {
		output = controlPolicyGroupMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) ListControlPolicyGroupMappings(ctx context.Context, controlPolicyGroupId string) ([]*ControlPolicyGroupMapping, error) {
	r := client.NewRequest(http.MethodGet, "/controlPolicyGroup/controlPolicyGroupMapping")
	r.Params.Set("controlPolicyGroupId", controlPolicyGroupId)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := controlPolicyGroupMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) UpdateControlPolicyGroupMapping(ctx context.Context, input *ControlPolicyGroupMapping) (*ControlPolicyGroupMapping, error) {
	r := client.NewRequest(http.MethodPut, "/controlPolicyGroup/controlPolicyGroupMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	controlPolicyGroupMapping, err := controlPolicyGroupMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicyGroupMapping)

	if len(controlPolicyGroupMapping) > 0 {
		output = controlPolicyGroupMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteControlPolicyGroupMapping(ctx context.Context, input *ControlPolicyGroupMapping) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/controlPolicyGroup/controlPolicyGroupMapping")
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

func controlPolicyGroupMappingFromJSON(in []byte) (*ControlPolicyGroupMapping, error) {
	b := new(ControlPolicyGroupMapping)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func controlPolicyGroupMappingsFromJSON(in []byte) ([]*ControlPolicyGroupMapping, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ControlPolicyGroupMapping, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := controlPolicyGroupMappingFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func controlPolicyGroupMappingsFromHttpResponse(resp *http.Response) ([]*ControlPolicyGroupMapping, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return controlPolicyGroupMappingsFromJSON(body)
}

//endregion

//region Setters

func (o ControlPolicyGroupMapping) MarshalJSON() ([]byte, error) {
	type noMethod ControlPolicyGroupMapping
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ControlPolicyGroupMapping) SetControlPolicyGroupId(v *string) *ControlPolicyGroupMapping {
	if o.ControlPolicyGroupId = v; o.ControlPolicyGroupId == nil {
		o.nullFields = append(o.nullFields, "ControlPolicyGroupId")
	}
	return o
}

func (o *ControlPolicyGroupMapping) SetTargetId(v *string) *ControlPolicyGroupMapping {
	if o.TargetId = v; o.TargetId == nil {
		o.nullFields = append(o.nullFields, "TargetId")
	}
	return o
}

func (o *ControlPolicyGroupMapping) SetTargetType(v *string) *ControlPolicyGroupMapping {
	if o.TargetType = v; o.TargetType == nil {
		o.nullFields = append(o.nullFields, "TargetType")
	}
	return o
}

func (o *ControlPolicyGroupMapping) SetEnforcementLevel(v *string) *ControlPolicyGroupMapping {
	if o.EnforcementLevel = v; o.EnforcementLevel == nil {
		o.nullFields = append(o.nullFields, "EnforcementLevel")
	}
	return o
}

func (o *ControlPolicyGroupMapping) SetOverrideEnforcements(v []*OverrideEnforcement) *ControlPolicyGroupMapping {
	if o.OverrideEnforcements = v; o.OverrideEnforcements == nil {
		o.nullFields = append(o.nullFields, "OverrideEnforcements")
	}
	return o
}

//region Override Enforcement

func (o OverrideEnforcement) MarshalJSON() ([]byte, error) {
	type noMethod OverrideEnforcement
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *OverrideEnforcement) SetControlPolicyId(v *string) *OverrideEnforcement {
	if o.ControlPolicyId = v; o.ControlPolicyId == nil {
		o.nullFields = append(o.nullFields, "ControlPolicyId")
	}
	return o
}

func (o *OverrideEnforcement) SetEnforcementLevel(v *string) *OverrideEnforcement {
	if o.EnforcementLevel = v; o.EnforcementLevel == nil {
		o.nullFields = append(o.nullFields, "EnforcementLevel")
	}
	return o
}

func (o *OverrideEnforcement) SetStackIds(v []*string) *OverrideEnforcement {
	if o.StackIds = v; o.StackIds == nil {
		o.nullFields = append(o.nullFields, "StackIds")
	}
	return o
}

//endregion

//endregion

//endregion
