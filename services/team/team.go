package team

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

//region Team

//region Structure

type Team struct {
	ID          *string `json:"id,omitempty"` // read-only
	Name        *string `json:"name,omitempty"`
	CustomIdpId *string `json:"customIdpId,omitempty"`

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

func (s *ServiceOp) CreateTeam(ctx context.Context, input *Team) (*Team, error) {
	r := client.NewRequest(http.MethodPost, "/iam/org/team")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	team, err := teamsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Team)
	if len(team) > 0 {
		output = team[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadTeam(ctx context.Context, teamId string) (*Team, error) {
	path, err := uritemplates.Expand("/iam/org/team/{teamId}", uritemplates.Values{"teamId": teamId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	team, err := teamsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Team)
	if len(team) > 0 {
		output = team[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateTeam(ctx context.Context, teamId string, input *Team) (*Team, error) {
	path, err := uritemplates.Expand("/iam/org/team/{teamId}", uritemplates.Values{"teamId": teamId})
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

	team, err := teamsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Team)
	if len(team) > 0 {
		output = team[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteTeam(ctx context.Context, teamId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/iam/org/team/{teamId}", uritemplates.Values{"teamId": teamId})
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

func teamFromJSON(in []byte) (*Team, error) {
	b := new(Team)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func teamsFromJSON(in []byte) ([]*Team, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Team, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := teamFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func teamsFromHttpResponse(resp *http.Response) ([]*Team, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return teamsFromJSON(body)
}

//endregion

//region Setters

func (o Team) MarshalJSON() ([]byte, error) {
	type noMethod Team
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Team) SetName(v *string) *Team {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Team) SetCustomIdpId(v *string) *Team {
	if o.CustomIdpId = v; o.CustomIdpId == nil {
		o.nullFields = append(o.nullFields, "CustomIdpId")
	}
	return o
}

//endregion

//endregion
