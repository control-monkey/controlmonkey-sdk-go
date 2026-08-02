package organization_permissions

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region OrganizationPermission

//region Structure

type OrganizationPermission struct {
	UserEmail            *string `json:"userEmail,omitempty"`
	ProgrammaticUserName *string `json:"programmaticUserName,omitempty"`
	TeamId               *string `json:"teamId,omitempty"`
	CustomRoleId         *string `json:"customRoleId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateOrganizationPermission(ctx context.Context, input *OrganizationPermission) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodPost, "/iam/org/organizationPermission")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output := new(commons.EmptyResponse)
	return output, nil
}

func (s *ServiceOp) ListOrganizationPermissions(ctx context.Context, teamId string) ([]*OrganizationPermission, error) {
	r := client.NewRequest(http.MethodGet, "/iam/org/organizationPermission")
	r.Params.Set("teamId", teamId)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := organizationPermissionsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) DeleteOrganizationPermission(ctx context.Context, input *OrganizationPermission) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/iam/org/organizationPermission")
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

func organizationPermissionFromJSON(in []byte) (*OrganizationPermission, error) {
	b := new(OrganizationPermission)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func organizationPermissionsFromJSON(in []byte) ([]*OrganizationPermission, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*OrganizationPermission, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := organizationPermissionFromJSON(rb)
		if err != nil {
			return nil, err
		}

		out[i] = b
	}
	return out, nil
}

func organizationPermissionsFromHttpResponse(resp *http.Response) ([]*OrganizationPermission, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return organizationPermissionsFromJSON(body)
}

//endregion

//region Setters

func (o OrganizationPermission) MarshalJSON() ([]byte, error) {
	type noMethod OrganizationPermission
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *OrganizationPermission) SetUserEmail(v *string) *OrganizationPermission {
	if o.UserEmail = v; o.UserEmail == nil {
		o.nullFields = append(o.nullFields, "UserEmail")
	}
	return o
}

func (o *OrganizationPermission) SetProgrammaticUserName(v *string) *OrganizationPermission {
	if o.ProgrammaticUserName = v; o.ProgrammaticUserName == nil {
		o.nullFields = append(o.nullFields, "ProgrammaticUserName")
	}
	return o
}

func (o *OrganizationPermission) SetTeamId(v *string) *OrganizationPermission {
	if o.TeamId = v; o.TeamId == nil {
		o.nullFields = append(o.nullFields, "TeamId")
	}
	return o
}

func (o *OrganizationPermission) SetCustomRoleId(v *string) *OrganizationPermission {
	if o.CustomRoleId = v; o.CustomRoleId == nil {
		o.nullFields = append(o.nullFields, "CustomRoleId")
	}
	return o
}

//endregion

//endregion
