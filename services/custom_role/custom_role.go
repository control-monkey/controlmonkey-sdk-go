package custom_role

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
	endpointUrl = "/customRole"
)

//region CustomRole

// region Structure

type CustomRole struct {
	ID               *string       `json:"id,omitempty"` // read-only
	Name             *string       `json:"name,omitempty"`
	Description      *string       `json:"description,omitempty"`
	Permissions      []*Permission `json:"permissions,omitempty"`
	StackRestriction *string       `json:"stackRestriction,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Permission struct {
	Name *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateCustomRole(ctx context.Context, input *CustomRole) (*CustomRole, error) {
	r := client.NewRequest(http.MethodPost, baseUrl+endpointUrl)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	customRole, err := customRolesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomRole)
	if len(customRole) > 0 {
		output = customRole[0]
	}

	return output, nil
}

func (s *ServiceOp) ListCustomRoles(ctx context.Context, customRoleId *string, customRoleName *string) ([]*CustomRole, error) {
	r := client.NewRequest(http.MethodGet, baseUrl+endpointUrl)

	if customRoleId != nil {
		r.Params.Set("customRoleId", *customRoleId)
	}
	if customRoleName != nil {
		r.Params.Set("customRoleName", *customRoleName)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	outputs, err := customRolesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}

func (s *ServiceOp) ReadCustomRole(ctx context.Context, customRoleId string) (*CustomRole, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{customRoleId}", uritemplates.Values{"customRoleId": customRoleId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	customRole, err := customRolesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomRole)
	if len(customRole) > 0 {
		output = customRole[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateCustomRole(ctx context.Context, id string, input *CustomRole) (*CustomRole, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{customRoleId}", uritemplates.Values{"customRoleId": id})
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

	customRole, err := customRolesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomRole)

	if len(customRole) > 0 {
		output = customRole[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteCustomRole(ctx context.Context, id string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand(baseUrl+endpointUrl+"/{customRoleId}", uritemplates.Values{"customRoleId": id})
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

func customRoleFromJSON(in []byte) (*CustomRole, error) {
	b := new(CustomRole)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func customRolesFromJSON(in []byte) ([]*CustomRole, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*CustomRole, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := customRoleFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func customRolesFromHttpResponse(resp *http.Response) ([]*CustomRole, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return customRolesFromJSON(body)
}

//endregion

//region Setters

func (o CustomRole) MarshalJSON() ([]byte, error) {
	type noMethod CustomRole
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CustomRole) SetID(v *string) *CustomRole {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *CustomRole) SetName(v *string) *CustomRole {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *CustomRole) SetDescription(v *string) *CustomRole {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *CustomRole) SetPermissions(v []*Permission) *CustomRole {
	if o.Permissions = v; o.Permissions == nil {
		o.nullFields = append(o.nullFields, "Permissions")
	}
	return o
}

func (o *CustomRole) SetStackRestriction(v *string) *CustomRole {
	if o.StackRestriction = v; o.StackRestriction == nil {
		o.nullFields = append(o.nullFields, "StackRestriction")
	}
	return o
}

//region Permissions

func (o Permission) MarshalJSON() ([]byte, error) {
	type noMethod Permission
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Permission) SetName(v *string) *Permission {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

//endregion

//endregion

//endregion
