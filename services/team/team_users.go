package team

import (
	"context"
	"encoding/json"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"io"
	"net/http"
)

//region TeamUser

//region Structure

type TeamUser struct {
	TeamId    *string `json:"teamId,omitempty"`
	UserEmail *string `json:"userEmail,omitempty"`

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

//region Requests & Responses

type teamUserResponse struct {
	Email *string `json:"email,omitempty"`
}

//endregion

//region Methods

func (s *ServiceOp) CreateTeamUser(ctx context.Context, input *TeamUser) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodPost, "/iam/org/teamUser")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output := new(commons.EmptyResponse)
	return output, nil
}

func (s *ServiceOp) ListTeamUsers(ctx context.Context, teamId string) ([]*TeamUser, error) {
	r := client.NewRequest(http.MethodGet, "/iam/org/teamUser")
	r.Params.Set("teamId", teamId)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := teamUsersFromHttpResponse(resp, teamId)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) DeleteTeamUser(ctx context.Context, input *TeamUser) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/iam/org/teamUser")
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

func teamUserFromJSON(in []byte) (*teamUserResponse, error) {
	b := new(teamUserResponse)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func teamUsersFromJSON(in []byte, teamId string) ([]*TeamUser, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*TeamUser, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := teamUserFromJSON(rb)
		if err != nil {
			return nil, err
		}

		//convert response object to TeamUser object
		user := new(TeamUser)
		user.UserEmail = b.Email
		user.TeamId = &teamId

		out[i] = user
	}
	return out, nil
}

func teamUsersFromHttpResponse(resp *http.Response, teamId string) ([]*TeamUser, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return teamUsersFromJSON(body, teamId)
}

//endregion

//region Setters

func (o TeamUser) MarshalJSON() ([]byte, error) {
	type noMethod TeamUser
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TeamUser) SetTeamId(v *string) *TeamUser {
	if o.TeamId = v; o.TeamId == nil {
		o.nullFields = append(o.nullFields, "TeamId")
	}
	return o
}

func (o *TeamUser) SetUserEmail(v *string) *TeamUser {
	if o.UserEmail = v; o.UserEmail == nil {
		o.nullFields = append(o.nullFields, "UserEmail")
	}
	return o
}

//endregion

//endregion
