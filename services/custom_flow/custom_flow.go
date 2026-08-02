package custom_flow

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

//region CustomFlow

//region Structure

type CustomFlow struct {
	ID        *string `json:"id,omitempty"` // read-only
	Name      *string `json:"name,omitempty"`
	FlowYaml  *string `json:"flowYaml,omitempty"`
	IsEnabled *bool   `json:"isEnabled,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateCustomFlow(ctx context.Context, input *CustomFlow) (*CustomFlow, error) {
	r := client.NewRequest(http.MethodPost, "/customFlow")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	customFlows, err := customFlowsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomFlow)
	if len(customFlows) > 0 {
		output = customFlows[0]
	}

	return output, nil
}

func (s *ServiceOp) ListCustomFlows(ctx context.Context, customFlowId *string, customFlowName *string,
	namespaceId *string, stackId *string) ([]*CustomFlow, error) {
	r := client.NewRequest(http.MethodGet, "/customFlow")

	if customFlowId != nil {
		r.Params.Set("customFlowId", *customFlowId)
	}
	if customFlowName != nil {
		r.Params.Set("customFlowName", *customFlowName)
	}
	if namespaceId != nil {
		r.Params.Set("namespaceId", *namespaceId)
	}
	if stackId != nil {
		r.Params.Set("stackId", *stackId)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return customFlowsFromHttpResponse(resp)
}

func (s *ServiceOp) ReadCustomFlow(ctx context.Context, customFlowId string) (*CustomFlow, error) {
	path, err := uritemplates.Expand("/customFlow/{customFlowId}", uritemplates.Values{"customFlowId": customFlowId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	customFlows, err := customFlowsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomFlow)
	if len(customFlows) > 0 {
		output = customFlows[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateCustomFlow(ctx context.Context, customFlowId string, input *CustomFlow) (*CustomFlow, error) {
	path, err := uritemplates.Expand("/customFlow/{customFlowId}", uritemplates.Values{"customFlowId": customFlowId})
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

	customFlows, err := customFlowsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CustomFlow)
	if len(customFlows) > 0 {
		output = customFlows[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteCustomFlow(ctx context.Context, customFlowId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/customFlow/{customFlowId}", uritemplates.Values{"customFlowId": customFlowId})
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

func customFlowFromJSON(in []byte) (*CustomFlow, error) {
	b := new(CustomFlow)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func customFlowsFromJSON(in []byte) ([]*CustomFlow, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*CustomFlow, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := customFlowFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func customFlowsFromHttpResponse(resp *http.Response) ([]*CustomFlow, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return customFlowsFromJSON(body)
}

//endregion

//region Setters

func (o CustomFlow) MarshalJSON() ([]byte, error) {
	type noMethod CustomFlow
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *CustomFlow) SetName(v *string) *CustomFlow {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *CustomFlow) SetFlowYaml(v *string) *CustomFlow {
	if o.FlowYaml = v; o.FlowYaml == nil {
		o.nullFields = append(o.nullFields, "FlowYaml")
	}
	return o
}

func (o *CustomFlow) SetIsEnabled(v *bool) *CustomFlow {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

//endregion

//endregion
