package stack

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

// region Dependency

// region Structure

type Dependency struct {
	ID               *string          `json:"id,omitempty"`
	StackId          *string          `json:"stackId,omitempty"`
	DependsOnStackId *string          `json:"dependsOnStackId,omitempty"`
	References       []*DependencyRef `json:"references,omitempty"`
	TriggerOption    *string          `json:"triggerOption,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DependencyRef struct {
	OutputOfStackToDependOn *string `json:"outputOfStackToDependOn,omitempty"`
	InputForStack           *string `json:"inputForStack,omitempty"`
	IncludeSensitiveOutput  *bool   `json:"includeSensitiveOutput,omitempty"`

	forceSendFields []string
	nullFields      []string
}

// endregion

// region Methods

func (s *ServiceOp) CreateDependency(ctx context.Context, input *Dependency) (*Dependency, error) {
	r := client.NewRequest(http.MethodPost, "/stack/dependency")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	items, err := dependenciesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	var out *Dependency
	if len(items) > 0 {
		out = items[0]
	} else {
		out = new(Dependency)
	}
	return out, nil
}

func (s *ServiceOp) ReadDependency(ctx context.Context, dependencyId string) (*Dependency, error) {
	path, err := uritemplates.Expand("/stack/dependency/{dependencyId}", uritemplates.Values{"dependencyId": dependencyId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	items, err := dependenciesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	var out *Dependency
	if len(items) > 0 {
		out = items[0]
	} else {
		out = new(Dependency)
	}
	return out, nil
}

func (s *ServiceOp) UpdateDependency(ctx context.Context, dependencyId string, input *Dependency) (*Dependency, error) {
	path, err := uritemplates.Expand("/stack/dependency/{dependencyId}", uritemplates.Values{"dependencyId": dependencyId})
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

	items, err := dependenciesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	var out *Dependency
	if len(items) > 0 {
		out = items[0]
	} else {
		out = new(Dependency)
	}
	return out, nil
}

func (s *ServiceOp) DeleteDependency(ctx context.Context, dependencyId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/stack/dependency/{dependencyId}", uritemplates.Values{"dependencyId": dependencyId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out := new(commons.EmptyResponse)
	return out, nil
}

// endregion

// region Private Methods

func dependencyFromJSON(in []byte) (*Dependency, error) {
	b := new(Dependency)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func dependenciesFromJSON(in []byte) ([]*Dependency, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Dependency, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := dependencyFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func dependenciesFromHttpResponse(resp *http.Response) ([]*Dependency, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return dependenciesFromJSON(body)
}

// endregion

// region Setters

func (o Dependency) MarshalJSON() ([]byte, error) {
	type noMethod Dependency
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Dependency) SetStackId(v *string) *Dependency {
	if o.StackId = v; o.StackId == nil {
		o.nullFields = append(o.nullFields, "StackId")
	}
	return o
}

func (o *Dependency) SetDependsOnStackId(v *string) *Dependency {
	if o.DependsOnStackId = v; o.DependsOnStackId == nil {
		o.nullFields = append(o.nullFields, "DependsOnStackId")
	}
	return o
}

func (o *Dependency) SetReferences(v []*DependencyRef) *Dependency {
	if o.References = v; o.References == nil {
		o.nullFields = append(o.nullFields, "References")
	}
	return o
}

func (o *Dependency) SetTriggerOption(v *string) *Dependency {
	if o.TriggerOption = v; o.TriggerOption == nil {
		o.nullFields = append(o.nullFields, "TriggerOption")
	}
	return o
}

func (o DependencyRef) MarshalJSON() ([]byte, error) {
	type noMethod DependencyRef
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DependencyRef) SetOutputOfStackToDependOn(v *string) *DependencyRef {
	if o.OutputOfStackToDependOn = v; o.OutputOfStackToDependOn == nil {
		o.nullFields = append(o.nullFields, "OutputOfStackToDependOn")
	}
	return o
}

func (o *DependencyRef) SetInputForStack(v *string) *DependencyRef {
	if o.InputForStack = v; o.InputForStack == nil {
		o.nullFields = append(o.nullFields, "InputForStack")
	}
	return o
}

func (o *DependencyRef) SetIncludeSensitiveOutput(v *bool) *DependencyRef {
	if o.IncludeSensitiveOutput = v; o.IncludeSensitiveOutput == nil {
		o.nullFields = append(o.nullFields, "IncludeSensitiveOutput")
	}
	return o
}

// endregion

// endregion
