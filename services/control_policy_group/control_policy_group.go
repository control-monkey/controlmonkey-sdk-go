package control_policy_group

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region ControlPolicyGroup

// region Structure

type ControlPolicyGroup struct {
	ID              *string          `json:"id,omitempty"` // read-only
	Name            *string          `json:"name,omitempty"`
	Description     *string          `json:"description,omitempty"`
	ControlPolicies []*ControlPolicy `json:"controlPolicies,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ControlPolicy struct {
	ControlPolicyId *string `json:"controlPolicyId,omitempty"` // read-only
	Severity        *string `json:"severity,omitempty"`        //commons.SeverityTypes

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateControlPolicyGroup(ctx context.Context, input *ControlPolicyGroup) (*ControlPolicyGroup, error) {
	r := client.NewRequest(http.MethodPost, "/controlPolicyGroup")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	controlPolicyGroup, err := controlPolicyGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicyGroup)
	if len(controlPolicyGroup) > 0 {
		output = controlPolicyGroup[0]
	}

	return output, nil
}

func (s *ServiceOp) ListControlPolicyGroups(ctx context.Context, controlPolicyGroupId *string, controlPolicyGroupName *string, includeManaged *bool) ([]*ControlPolicyGroup, error) {
	r := client.NewRequest(http.MethodGet, "/controlPolicyGroup")

	if controlPolicyGroupId != nil {
		r.Params.Set("controlPolicyGroupId", *controlPolicyGroupId)
	}
	if controlPolicyGroupName != nil {
		r.Params.Set("controlPolicyGroupName", *controlPolicyGroupName)
	}
	if includeManaged != nil {
		r.Params.Set("includeManaged", strconv.FormatBool(*includeManaged))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := controlPolicyGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (s *ServiceOp) ReadControlPolicyGroup(ctx context.Context, controlPolicyGroupId string) (*ControlPolicyGroup, error) {
	path, err := uritemplates.Expand("/controlPolicyGroup/{controlPolicyGroupId}", uritemplates.Values{"controlPolicyGroupId": controlPolicyGroupId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	controlPolicyGroup, err := controlPolicyGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicyGroup)
	if len(controlPolicyGroup) > 0 {
		output = controlPolicyGroup[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateControlPolicyGroup(ctx context.Context, id string, input *ControlPolicyGroup) (*ControlPolicyGroup, error) {
	path, err := uritemplates.Expand("/controlPolicyGroup/{controlPolicyGroupId}", uritemplates.Values{"controlPolicyGroupId": id})
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

	controlPolicyGroup, err := controlPolicyGroupsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ControlPolicyGroup)

	if len(controlPolicyGroup) > 0 {
		output = controlPolicyGroup[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteControlPolicyGroup(ctx context.Context, id string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/controlPolicyGroup/{controlPolicyGroupId}", uritemplates.Values{"controlPolicyGroupId": id})
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

func controlPolicyGroupFromJSON(in []byte) (*ControlPolicyGroup, error) {
	b := new(ControlPolicyGroup)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func controlPolicyGroupsFromJSON(in []byte) ([]*ControlPolicyGroup, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*ControlPolicyGroup, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := controlPolicyGroupFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func controlPolicyGroupsFromHttpResponse(resp *http.Response) ([]*ControlPolicyGroup, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return controlPolicyGroupsFromJSON(body)
}

//endregion

//region Setters

func (o ControlPolicyGroup) MarshalJSON() ([]byte, error) {
	type noMethod ControlPolicyGroup
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ControlPolicyGroup) SetID(v *string) *ControlPolicyGroup {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *ControlPolicyGroup) SetName(v *string) *ControlPolicyGroup {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *ControlPolicyGroup) SetDescription(v *string) *ControlPolicyGroup {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *ControlPolicyGroup) SetControlPolicies(v []*ControlPolicy) *ControlPolicyGroup {
	if o.ControlPolicies = v; o.ControlPolicies == nil {
		o.nullFields = append(o.nullFields, "ControlPolicies")
	}
	return o
}

//region ControlPolicies

func (o ControlPolicy) MarshalJSON() ([]byte, error) {
	type noMethod ControlPolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ControlPolicy) SetControlPolicyId(v *string) *ControlPolicy {
	if o.ControlPolicyId = v; o.ControlPolicyId == nil {
		o.nullFields = append(o.nullFields, "ControlPolicyId")
	}
	return o
}

func (o *ControlPolicy) SetSeverity(v *string) *ControlPolicy {
	if o.Severity = v; o.Severity == nil {
		o.nullFields = append(o.nullFields, "Severity")
	}
	return o
}

//endregion

//endregion

//endregion
