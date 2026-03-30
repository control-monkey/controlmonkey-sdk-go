package run_task

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

//region RunTask

//region Structure

type RunTask struct {
	ID                  *string `json:"id,omitempty"`
	Name                *string `json:"name,omitempty"`
	Url                 *string `json:"url,omitempty"`
	IsEnabled           *bool   `json:"isEnabled,omitempty"`
	HmacKey             *string `json:"hmacKey,omitempty"`
	IsHmacKeyConfigured *bool   `json:"isHmacKeyConfigured,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateRunTask(ctx context.Context, input *RunTask) (*RunTask, error) {
	r := client.NewRequest(http.MethodPost, "/runTask")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	runTasks, err := runTasksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(RunTask)
	if len(runTasks) > 0 {
		output = runTasks[0]
	}

	return output, nil
}

func (s *ServiceOp) ListRunTasks(ctx context.Context, runTaskId *string, runTaskName *string) ([]*RunTask, error) {
	r := client.NewRequest(http.MethodGet, "/runTask")

	if runTaskId != nil {
		r.Params.Set("runTaskId", *runTaskId)
	}
	if runTaskName != nil {
		r.Params.Set("runTaskName", *runTaskName)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return runTasksFromHttpResponse(resp)
}

func (s *ServiceOp) ReadRunTask(ctx context.Context, runTaskId string) (*RunTask, error) {
	path, err := uritemplates.Expand("/runTask/{runTaskId}", uritemplates.Values{"runTaskId": runTaskId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	runTasks, err := runTasksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(RunTask)
	if len(runTasks) > 0 {
		output = runTasks[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateRunTask(ctx context.Context, runTaskId string, input *RunTask) (*RunTask, error) {
	path, err := uritemplates.Expand("/runTask/{runTaskId}", uritemplates.Values{"runTaskId": runTaskId})
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

	runTasks, err := runTasksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(RunTask)
	if len(runTasks) > 0 {
		output = runTasks[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteRunTask(ctx context.Context, runTaskId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/runTask/{runTaskId}", uritemplates.Values{"runTaskId": runTaskId})
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

func runTaskFromJSON(in []byte) (*RunTask, error) {
	b := new(RunTask)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func runTasksFromJSON(in []byte) ([]*RunTask, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*RunTask, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := runTaskFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func runTasksFromHttpResponse(resp *http.Response) ([]*RunTask, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return runTasksFromJSON(body)
}

//endregion

//region Setters

//region RunTask

func (o RunTask) MarshalJSON() ([]byte, error) {
	type noMethod RunTask
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RunTask) SetName(v *string) *RunTask {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *RunTask) SetUrl(v *string) *RunTask {
	if o.Url = v; o.Url == nil {
		o.nullFields = append(o.nullFields, "Url")
	}
	return o
}

func (o *RunTask) SetIsEnabled(v *bool) *RunTask {
	if o.IsEnabled = v; o.IsEnabled == nil {
		o.nullFields = append(o.nullFields, "IsEnabled")
	}
	return o
}

func (o *RunTask) SetHmacKey(v *string) *RunTask {
	if o.HmacKey = v; o.HmacKey == nil {
		o.nullFields = append(o.nullFields, "HmacKey")
	}
	return o
}

//endregion

//endregion
