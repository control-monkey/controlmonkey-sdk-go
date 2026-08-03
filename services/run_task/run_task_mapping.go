package run_task

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region RunTaskMapping

// region Structure

type RunTaskMapping struct {
	RunTaskId        *string `json:"runTaskId,omitempty"`
	TargetId         *string `json:"targetId,omitempty"`
	TargetType       *string `json:"targetType,omitempty"`       //commons.PolicyMappingTargetTypes
	EnforcementLevel *string `json:"enforcementLevel,omitempty"` //commons.EnforcementLevelTypes
	Stage            *string `json:"stage,omitempty"`            //commons.RunTaskStageTypes

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateRunTaskMapping(ctx context.Context, input *RunTaskMapping) (*RunTaskMapping, error) {
	r := client.NewRequest(http.MethodPost, "/runTask/runTaskMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	runTaskMapping, err := runTaskMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(RunTaskMapping)
	if len(runTaskMapping) > 0 {
		output = runTaskMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) ListRunTaskMappings(ctx context.Context, runTaskId string) ([]*RunTaskMapping, error) {
	r := client.NewRequest(http.MethodGet, "/runTask/runTaskMapping")
	r.Params.Set("runTaskId", runTaskId)

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := runTaskMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) UpdateRunTaskMapping(ctx context.Context, input *RunTaskMapping) (*RunTaskMapping, error) {
	r := client.NewRequest(http.MethodPut, "/runTask/runTaskMapping")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	runTaskMapping, err := runTaskMappingsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(RunTaskMapping)

	if len(runTaskMapping) > 0 {
		output = runTaskMapping[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteRunTaskMapping(ctx context.Context, input *RunTaskMapping) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, "/runTask/runTaskMapping")
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

func runTaskMappingFromJSON(in []byte) (*RunTaskMapping, error) {
	b := new(RunTaskMapping)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func runTaskMappingsFromJSON(in []byte) ([]*RunTaskMapping, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*RunTaskMapping, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := runTaskMappingFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func runTaskMappingsFromHttpResponse(resp *http.Response) ([]*RunTaskMapping, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return runTaskMappingsFromJSON(body)
}

//endregion

//region Setters

func (o RunTaskMapping) MarshalJSON() ([]byte, error) {
	type noMethod RunTaskMapping
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RunTaskMapping) SetRunTaskId(v *string) *RunTaskMapping {
	if o.RunTaskId = v; o.RunTaskId == nil {
		o.nullFields = append(o.nullFields, "RunTaskId")
	}
	return o
}

func (o *RunTaskMapping) SetTargetId(v *string) *RunTaskMapping {
	if o.TargetId = v; o.TargetId == nil {
		o.nullFields = append(o.nullFields, "TargetId")
	}
	return o
}

func (o *RunTaskMapping) SetTargetType(v *string) *RunTaskMapping {
	if o.TargetType = v; o.TargetType == nil {
		o.nullFields = append(o.nullFields, "TargetType")
	}
	return o
}

func (o *RunTaskMapping) SetEnforcementLevel(v *string) *RunTaskMapping {
	if o.EnforcementLevel = v; o.EnforcementLevel == nil {
		o.nullFields = append(o.nullFields, "EnforcementLevel")
	}
	return o
}

func (o *RunTaskMapping) SetStage(v *string) *RunTaskMapping {
	if o.Stage = v; o.Stage == nil {
		o.nullFields = append(o.nullFields, "Stage")
	}
	return o
}

//endregion

//endregion
