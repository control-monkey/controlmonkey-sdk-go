package namespace_permissions

import (
	"context"
	"encoding/json"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"io"
	"net/http"
)

//region NamespacePermission

//region Structure

type NamespacePermission struct {
	NamespaceId          *string `json:"namespaceId,omitempty"`
	UserEmail            *string `json:"userEmail,omitempty"`
	ProgrammaticUserName *string `json:"programmaticUserName,omitempty"`
	TeamId               *string `json:"teamId,omitempty"`
	Role                 *string `json:"role,omitempty"`
	CustomRoleId         *string `json:"customRoleId,omitempty"`

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

func (s *ServiceOp) CreateNamespacePermission(ctx context.Context, input *NamespacePermission) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodPost, "/iam/org/namespacePermission")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output := new(commons.EmptyResponse)
	return output, nil
}

func (s *ServiceOp) ListNamespacePermissions(ctx context.Context, namespaceId string) ([]*NamespacePermission, error) {
	r := client.NewRequest(http.MethodGet, "/iam/org/namespacePermission")
	r.Params.Set("namespaceId", namespaceId)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := namespaceUsersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) DeleteNamespacePermission(ctx context.Context, input *NamespacePermission) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/iam/org/namespacePermission")
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

func namespacePermissionFromJSON(in []byte) (*NamespacePermission, error) {
	b := new(NamespacePermission)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func namespacePermissionsFromJSON(in []byte) ([]*NamespacePermission, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*NamespacePermission, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := namespacePermissionFromJSON(rb)
		if err != nil {
			return nil, err
		}

		out[i] = b
	}
	return out, nil
}

func namespaceUsersFromHttpResponse(resp *http.Response) ([]*NamespacePermission, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return namespacePermissionsFromJSON(body)
}

//endregion

//region Setters

func (o NamespacePermission) MarshalJSON() ([]byte, error) {
	type noMethod NamespacePermission
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NamespacePermission) SetNamespaceId(v *string) *NamespacePermission {
	if o.NamespaceId = v; o.NamespaceId == nil {
		o.nullFields = append(o.nullFields, "NamespaceId")
	}
	return o
}

func (o *NamespacePermission) SetUserEmail(v *string) *NamespacePermission {
	if o.UserEmail = v; o.UserEmail == nil {
		o.nullFields = append(o.nullFields, "UserEmail")
	}
	return o
}

func (o *NamespacePermission) SetProgrammaticUserName(v *string) *NamespacePermission {
	if o.ProgrammaticUserName = v; o.ProgrammaticUserName == nil {
		o.nullFields = append(o.nullFields, "ProgrammaticUserName")
	}
	return o
}

func (o *NamespacePermission) SetTeamId(v *string) *NamespacePermission {
	if o.TeamId = v; o.TeamId == nil {
		o.nullFields = append(o.nullFields, "TeamId")
	}
	return o
}

func (o *NamespacePermission) SetRole(v *string) *NamespacePermission {
	if o.Role = v; o.Role == nil {
		o.nullFields = append(o.nullFields, "Role")
	}
	return o
}

func (o *NamespacePermission) SetCustomRoleId(v *string) *NamespacePermission {
	if o.CustomRoleId = v; o.CustomRoleId == nil {
		o.nullFields = append(o.nullFields, "CustomRoleId")
	}
	return o
}

//endregion

//endregion
