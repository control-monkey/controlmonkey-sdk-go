package variable

import (
	"context"
	"encoding/json"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
)

// region Variable

type Variable struct {
	ID            *string `json:"id,omitempty"`
	Scope         *string `json:"scope,omitempty"`
	ScopeId       *string `json:"scopeId,omitempty"`
	Key           *string `json:"key,omitempty"`
	Value         *string `json:"value,omitempty"`
	Type          *string `json:"type,omitempty"`
	IsSensitive   *bool   `json:"isSensitive,omitempty"`
	IsOverridable *bool   `json:"isOverridable,omitempty"`
	IsRequired    *bool   `json:"isRequired,omitempty"`
	Description   *string `json:"description,omitempty"`

	// Read-only fields.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

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

type ListVariablesInput struct {
	StackId     *string `json:"stackId,omitempty"`
	NamespaceId *string `json:"namespaceId,omitempty"`
	TemplateId  *string `json:"templateId,omitempty"`
	StackRunId  *string `json:"stackRunId,omitempty"`
	OrgOnly     *bool   `json:"orgOnly,omitempty"`
}

type ListVariablesOutput struct {
	Variables []*Variable `json:"variables,omitempty"`
}

type CreateVariableOutput struct {
	Variable *Variable `json:"variable,omitempty"`
}

type ReadVariableInput struct {
	VariableId *string `json:"variableId,omitempty"`
}

type ReadVariableOutput struct {
	Variable *Variable `json:"variable,omitempty"`
}

type UpdateVariableOutput struct {
	Variable *Variable `json:"variable,omitempty"`
}

type DeleteVariableInput struct {
	VariableId *string `json:"variableId"`
}

type DeleteVariableOutput struct{}

func variableFromJSON(in []byte) (*Variable, error) {
	b := new(Variable)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func variablesFromJSON(in []byte) ([]*Variable, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Variable, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := variableFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func variablesFromHttpResponse(resp *http.Response) ([]*Variable, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return variablesFromJSON(body)
}

func (s *ServiceOp) ListVariables(ctx context.Context, input *ListVariablesInput) (*ListVariablesOutput, error) {
	r := client.NewRequest(http.MethodGet, "/variable")

	if input.StackId != nil {
		r.Params.Set("stackId", controlmonkey.StringValue(input.StackId))
	}
	if input.StackRunId != nil {
		r.Params.Set("stackRunId", controlmonkey.StringValue(input.StackRunId))
	}
	if input.NamespaceId != nil {
		r.Params.Set("namespaceId", controlmonkey.StringValue(input.NamespaceId))
	}
	if input.TemplateId != nil {
		r.Params.Set("templateId", controlmonkey.StringValue(input.TemplateId))
	}
	if input.OrgOnly != nil {
		r.Params.Set("orgOnly", strconv.FormatBool(controlmonkey.BoolValue(input.OrgOnly)))
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	variables, err := variablesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListVariablesOutput{Variables: variables}, nil
}

func (s *ServiceOp) CreateVariable(ctx context.Context, input *Variable) (*CreateVariableOutput, error) {
	r := client.NewRequest(http.MethodPost, "/variable")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	variable, err := variablesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateVariableOutput)
	if len(variable) > 0 {
		output.Variable = variable[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadVariable(ctx context.Context, input *ReadVariableInput) (*ReadVariableOutput, error) {
	path, err := uritemplates.Expand("/variable/{variableId}", uritemplates.Values{
		"variableId": controlmonkey.StringValue(input.VariableId),
	})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	variable, err := variablesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadVariableOutput)
	if len(variable) > 0 {
		output.Variable = variable[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateVariable(ctx context.Context, variableId *string, input *Variable) (*UpdateVariableOutput, error) {
	path, err := uritemplates.Expand("/variable/{variableId}", uritemplates.Values{
		"variableId": controlmonkey.StringValue(variableId),
	})
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

	variable, err := variablesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateVariableOutput)
	if len(variable) > 0 {
		output.Variable = variable[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteVariable(ctx context.Context, input *DeleteVariableInput) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/variable/{variableId}", uritemplates.Values{
		"variableId": controlmonkey.StringValue(input.VariableId),
	})
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

func (o Variable) MarshalJSON() ([]byte, error) {
	type noMethod Variable
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Variable) SetScope(v *string) *Variable {
	o.Scope = v
	return o
}

func (o *Variable) SetScopeId(v *string) *Variable {
	o.ScopeId = v
	return o
}

func (o *Variable) SetKey(v *string) *Variable {
	o.Key = v
	return o
}

func (o *Variable) SetType(v *string) *Variable {
	o.Type = v
	return o
}

func (o *Variable) SetValue(v *string) *Variable {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

func (o *Variable) SetIsSensitive(v *bool) *Variable {
	o.IsSensitive = v
	return o
}

func (o *Variable) SetIsOverridable(v *bool) *Variable {
	o.IsOverridable = v
	return o
}

func (o *Variable) SetIsRequired(v *bool) *Variable {
	if o.IsRequired = v; o.IsRequired == nil {
		o.nullFields = append(o.nullFields, "IsRequired")
	}
	return o
}

func (o *Variable) SetDescription(v *string) *Variable {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

// endregion
