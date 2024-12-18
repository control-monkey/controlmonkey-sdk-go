package custom_abac_configuration

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

const (
	baseUrl     = "/iam/org"
	endpointUrl = "/customAbacConfiguration"
)

//region CustomAbacConfiguration

// region Structure

type CustomAbacConfiguration struct {
	ID           *string `json:"id,omitempty"` // read-only
	CustomAbacId *string `json:"customAbacId,omitempty"`
	Name         *string `json:"name,omitempty"`
	Roles        []*Role `json:"roles,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Role struct {
	OrgId   *string   `json:"orgId,omitempty"`
	OrgRole *string   `json:"orgRole,omitempty"`
	TeamIds []*string `json:"teamIds,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateCustomAbacConfiguration(ctx context.Context, input *CustomAbacConfiguration) (*CustomAbacConfiguration, error) {
	r := client.NewRequest(http.MethodPost, baseUrl+endpointUrl)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	customAbacConfiguration, err := customAbacConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomAbacConfiguration)
	if len(customAbacConfiguration) > 0 {
		output = customAbacConfiguration[0]
	}

	return output, nil
}

func (s *ServiceOp) ListCustomAbacConfigurations(ctx context.Context, customAbacConfigurationId *string, customAbacConfigurationName *string) ([]*CustomAbacConfiguration, error) {
	r := client.NewRequest(http.MethodGet, baseUrl+endpointUrl)

	if customAbacConfigurationId != nil {
		r.Params.Set("customAbacConfigurationId", *customAbacConfigurationId)
	}
	if customAbacConfigurationName != nil {
		r.Params.Set("customAbacConfigurationName", *customAbacConfigurationName)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	outputs, err := customAbacConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}

func (s *ServiceOp) ReadCustomAbacConfiguration(ctx context.Context, customAbacConfigurationId string) (*CustomAbacConfiguration, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{customAbacConfigurationId}", uritemplates.Values{"customAbacConfigurationId": customAbacConfigurationId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	customAbacConfiguration, err := customAbacConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomAbacConfiguration)
	if len(customAbacConfiguration) > 0 {
		output = customAbacConfiguration[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateCustomAbacConfiguration(ctx context.Context, id string, input *CustomAbacConfiguration) (*CustomAbacConfiguration, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{customAbacConfigurationId}", uritemplates.Values{"customAbacConfigurationId": id})
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

	customAbacConfiguration, err := customAbacConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomAbacConfiguration)

	if len(customAbacConfiguration) > 0 {
		output = customAbacConfiguration[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteCustomAbacConfiguration(ctx context.Context, id string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{customAbacConfigurationId}", uritemplates.Values{"customAbacConfigurationId": id})
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

func customAbacConfigurationFromJSON(in []byte) (*CustomAbacConfiguration, error) {
	b := new(CustomAbacConfiguration)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func customAbacConfigurationsFromJSON(in []byte) ([]*CustomAbacConfiguration, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*CustomAbacConfiguration, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := customAbacConfigurationFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func customAbacConfigurationsFromHttpResponse(resp *http.Response) ([]*CustomAbacConfiguration, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return customAbacConfigurationsFromJSON(body)
}

//endregion

//region Setters

func (o CustomAbacConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod CustomAbacConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CustomAbacConfiguration) SetID(v *string) *CustomAbacConfiguration {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *CustomAbacConfiguration) SetCustomAbacId(v *string) *CustomAbacConfiguration {
	if o.CustomAbacId = v; o.CustomAbacId == nil {
		o.nullFields = append(o.nullFields, "CustomAbacId")
	}
	return o
}

func (o *CustomAbacConfiguration) SetName(v *string) *CustomAbacConfiguration {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *CustomAbacConfiguration) SetRoles(v []*Role) *CustomAbacConfiguration {
	if o.Roles = v; o.Roles == nil {
		o.nullFields = append(o.nullFields, "Roles")
	}
	return o
}

//region Roles

func (o Role) MarshalJSON() ([]byte, error) {
	type noMethod Role
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Role) SetOrgId(v *string) *Role {
	if o.OrgId = v; o.OrgId == nil {
		o.nullFields = append(o.nullFields, "OrgId")
	}
	return o
}

func (o *Role) SetOrgRole(v *string) *Role {
	if o.OrgRole = v; o.OrgRole == nil {
		o.nullFields = append(o.nullFields, "OrgRole")
	}
	return o
}

func (o *Role) SetTeamIds(v []*string) *Role {
	if o.TeamIds = v; o.TeamIds == nil {
		o.nullFields = append(o.nullFields, "TeamIds")
	}
	return o
}

//endregion

//endregion

//endregion
