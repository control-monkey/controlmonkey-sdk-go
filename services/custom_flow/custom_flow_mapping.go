package custom_flow

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region CustomFlowMapping

// region Structure

type CustomFlowMapping struct {
	CustomFlowId *string `json:"customFlowId,omitempty"`
	TargetId     *string `json:"targetId,omitempty"`
	TargetType   *string `json:"targetType,omitempty"` //commons.PolicyMappingTargetTypes

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateCustomFlowMapping(ctx context.Context, input *CustomFlowMapping) (*CustomFlowMapping, error) {
	r := client.NewRequest(http.MethodPost, "/customFlow/customFlowMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	customFlowMapping, err := customFlowMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomFlowMapping)
	if len(customFlowMapping) > 0 {
		output = customFlowMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) ListCustomFlowMappings(ctx context.Context) ([]*CustomFlowMapping, error) {
	r := client.NewRequest(http.MethodGet, "/customFlow/customFlowMapping")

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := customFlowMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) DeleteCustomFlowMapping(ctx context.Context, input *CustomFlowMapping) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/customFlow/customFlowMapping")
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

func customFlowMappingFromJSON(in []byte) (*CustomFlowMapping, error) {
	b := new(CustomFlowMapping)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func customFlowMappingsFromJSON(in []byte) ([]*CustomFlowMapping, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*CustomFlowMapping, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := customFlowMappingFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func customFlowMappingsFromHttpResponse(resp *http.Response) ([]*CustomFlowMapping, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return customFlowMappingsFromJSON(body)
}

//endregion

//region Setters

func (o CustomFlowMapping) MarshalJSON() ([]byte, error) {
	type noMethod CustomFlowMapping
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CustomFlowMapping) SetCustomFlowId(v *string) *CustomFlowMapping {
	if o.CustomFlowId = v; o.CustomFlowId == nil {
		o.nullFields = append(o.nullFields, "CustomFlowId")
	}
	return o
}

func (o *CustomFlowMapping) SetTargetId(v *string) *CustomFlowMapping {
	if o.TargetId = v; o.TargetId == nil {
		o.nullFields = append(o.nullFields, "TargetId")
	}
	return o
}

func (o *CustomFlowMapping) SetTargetType(v *string) *CustomFlowMapping {
	if o.TargetType = v; o.TargetType == nil {
		o.nullFields = append(o.nullFields, "TargetType")
	}
	return o
}

//endregion

//endregion
