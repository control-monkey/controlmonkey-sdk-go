package blueprint

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
)

//region Blueprint

// region Structure

type Blueprint struct {
	ID   *string `json:"id,omitempty"` // read-only
	Name *string `json:"name,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) ListBlueprints(ctx context.Context, blueprintId *string, blueprintName *string) ([]*Blueprint, error) {
	r := client.NewRequest(http.MethodGet, "/blueprint")

	if blueprintId != nil {
		r.Params.Set("blueprintId", *blueprintId)
	}
	if blueprintName != nil {
		r.Params.Set("blueprintName", *blueprintName)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	blueprints, err := blueprintsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return blueprints, nil
}

//endregion

//region Private Methods

func blueprintFromJSON(in []byte) (*Blueprint, error) {
	b := new(Blueprint)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func blueprintsFromJSON(in []byte) ([]*Blueprint, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Blueprint, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := blueprintFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func blueprintsFromHttpResponse(resp *http.Response) ([]*Blueprint, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return blueprintsFromJSON(body)
}

//endregion

//region Setters

func (o Blueprint) MarshalJSON() ([]byte, error) {
	type noMethod Blueprint
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Blueprint) SetName(v *string) *Blueprint {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

//endregion
