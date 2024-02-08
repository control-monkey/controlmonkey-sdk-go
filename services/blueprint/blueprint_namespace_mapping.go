package blueprint

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region Blueprint

// region Structure

type BlueprintNamespaceMapping struct {
	BlueprintId *string `json:"blueprintId,omitempty"`
	NamespaceId *string `json:"namespaceId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateBlueprintNamespaceMapping(ctx context.Context, input *BlueprintNamespaceMapping) (*BlueprintNamespaceMapping, error) {
	r := client.NewRequest(http.MethodPost, "/blueprint/blueprintNamespaceMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	blueprintNamespaceMapping, err := blueprintNamespaceMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(BlueprintNamespaceMapping)
	if len(blueprintNamespaceMapping) > 0 {
		output = blueprintNamespaceMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) ListBlueprintNamespaceMappings(ctx context.Context, blueprintId string) ([]*BlueprintNamespaceMapping, error) {
	r := client.NewRequest(http.MethodGet, "/blueprint/blueprintNamespaceMapping")
	r.Params.Set("blueprintId", blueprintId)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := blueprintNamespaceMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) DeleteBlueprintNamespaceMapping(ctx context.Context, input *BlueprintNamespaceMapping) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/blueprint/blueprintNamespaceMapping")
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

func blueprintNamespaceMappingFromJSON(in []byte) (*BlueprintNamespaceMapping, error) {
	b := new(BlueprintNamespaceMapping)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func blueprintNamespaceMappingsFromJSON(in []byte) ([]*BlueprintNamespaceMapping, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*BlueprintNamespaceMapping, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := blueprintNamespaceMappingFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func blueprintNamespaceMappingsFromHttpResponse(resp *http.Response) ([]*BlueprintNamespaceMapping, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return blueprintNamespaceMappingsFromJSON(body)
}

//endregion

//region Setters

func (o BlueprintNamespaceMapping) MarshalJSON() ([]byte, error) {
	type noMethod BlueprintNamespaceMapping
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BlueprintNamespaceMapping) SetBlueprintId(v *string) *BlueprintNamespaceMapping {
	if o.BlueprintId = v; o.BlueprintId == nil {
		o.nullFields = append(o.nullFields, "BlueprintId")
	}
	return o
}

func (o *BlueprintNamespaceMapping) SetNamespaceId(v *string) *BlueprintNamespaceMapping {
	if o.NamespaceId = v; o.NamespaceId == nil {
		o.nullFields = append(o.nullFields, "NamespaceId")
	}
	return o
}

//endregion
