package template

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region Template

// region Structure

type TemplateNamespaceMapping struct {
	TemplateId  *string `json:"templateId,omitempty"`
	NamespaceId *string `json:"namespaceId,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateTemplateNamespaceMapping(ctx context.Context, input *TemplateNamespaceMapping) (*TemplateNamespaceMapping, error) {
	r := client.NewRequest(http.MethodPost, "/template/templateNamespaceMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	templateNamespaceMapping, err := templateNamespaceMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(TemplateNamespaceMapping)
	if len(templateNamespaceMapping) > 0 {
		output = templateNamespaceMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) ListTemplateNamespaceMappings(ctx context.Context, templateId string) ([]*TemplateNamespaceMapping, error) {
	r := client.NewRequest(http.MethodGet, "/template/templateNamespaceMapping")
	r.Params.Set("templateId", templateId)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := templateNamespaceMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) DeleteTemplateNamespaceMapping(ctx context.Context, input *TemplateNamespaceMapping) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/template/templateNamespaceMapping")
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

func templateNamespaceMappingFromJSON(in []byte) (*TemplateNamespaceMapping, error) {
	b := new(TemplateNamespaceMapping)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func templateNamespaceMappingsFromJSON(in []byte) ([]*TemplateNamespaceMapping, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*TemplateNamespaceMapping, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := templateNamespaceMappingFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func templateNamespaceMappingsFromHttpResponse(resp *http.Response) ([]*TemplateNamespaceMapping, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return templateNamespaceMappingsFromJSON(body)
}

//endregion

//region Setters

func (o TemplateNamespaceMapping) MarshalJSON() ([]byte, error) {
	type noMethod TemplateNamespaceMapping
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TemplateNamespaceMapping) SetTemplateId(v *string) *TemplateNamespaceMapping {
	if o.TemplateId = v; o.TemplateId == nil {
		o.nullFields = append(o.nullFields, "TemplateId")
	}
	return o
}

func (o *TemplateNamespaceMapping) SetNamespaceId(v *string) *TemplateNamespaceMapping {
	if o.NamespaceId = v; o.NamespaceId == nil {
		o.nullFields = append(o.nullFields, "NamespaceId")
	}
	return o
}

//endregion
