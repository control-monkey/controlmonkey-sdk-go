package stack

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
)

//region Stack

// region Structure

type Stack struct {
	ID          *string `json:"id,omitempty"`      // read-only
	IacType     *string `json:"iacType,omitempty"` //commons.IacTypes
	NamespaceId *string `json:"namespaceId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Data        *Data   `json:"data,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Data struct {
	DeploymentBehavior       *cross_models.DeploymentBehavior       `json:"deploymentBehavior,omitempty"`
	DeploymentApprovalPolicy *cross_models.DeploymentApprovalPolicy `json:"deploymentApprovalPolicy,omitempty"`
	VcsInfo                  *VcsInfo                               `json:"vcsInfo,omitempty"`
	RunTrigger               *cross_models.RunTrigger               `json:"runTrigger,omitempty"`
	IacConfig                *cross_models.IacConfig                `json:"iacConfig,omitempty"`
	Policy                   *Policy                                `json:"policy,omitempty"`
	RunnerConfig             *cross_models.RunnerConfig             `json:"runnerConfig,omitempty"`
	AutoSync                 *cross_models.AutoSync                 `json:"autoSync,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VcsInfo struct {
	ProviderId *string `json:"providerId,omitempty"`
	RepoName   *string `json:"repoName,omitempty"`
	Path       *string `json:"path,omitempty"`
	Branch     *string `json:"branch,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Policy struct {
	TtlConfig *TtlConfig `json:"ttlConfig,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TtlConfig struct {
	Ttl         *TtlDefinition `json:"ttl,omitempty"`
	TtlOverride *TtlOverride   `json:"ttlOverride,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TtlDefinition struct {
	Type  *string `json:"type,omitempty"` //commons.TtlTypes
	Value *int    `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TtlOverride struct {
	Type          *string    `json:"type,omitempty"` //commons.TtlTypes
	Value         *int       `json:"value,omitempty"`
	EffectiveFrom *time.Time `json:"effectiveFrom,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

type CreateStackInput struct {
	Stack *Stack `json:"stack,omitempty"`
}

func (s *ServiceOp) CreateStack(ctx context.Context, input *Stack) (*Stack, error) {
	r := client.NewRequest(http.MethodPost, "/stack")
	r.Obj = CreateStackInput{input}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	stack, err := stacksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Stack)
	if len(stack) > 0 {
		output = stack[0]
	}

	return output, nil
}

func (s *ServiceOp) ListStacks(ctx context.Context, stackId *string, stackName *string, namespaceId *string) ([]*Stack, error) {
	r := client.NewRequest(http.MethodGet, "/stack")

	if stackId != nil {
		r.Params.Set("stackId", *stackId)
	}
	if stackName != nil {
		r.Params.Set("stackName", *stackName)
	}
	if namespaceId != nil {
		r.Params.Set("namespaceId", *namespaceId)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := stacksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) ReadStack(ctx context.Context, stackId string) (*Stack, error) {
	path, err := uritemplates.Expand("/stack/{stackId}", uritemplates.Values{
		"stackId": stackId,
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

	stack, err := stacksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Stack)
	if len(stack) > 0 {
		output = stack[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateStack(ctx context.Context, stackId string, input *Stack) (*Stack, error) {
	path, err := uritemplates.Expand("/stack/{stackId}", uritemplates.Values{"stackId": stackId})
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

	stack, err := stacksFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Stack)
	if len(stack) > 0 {
		output = stack[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteStack(ctx context.Context, stackId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/stack/{stackId}", uritemplates.Values{"stackId": stackId})
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

func stackFromJSON(in []byte) (*Stack, error) {
	b := new(Stack)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func stacksFromJSON(in []byte) ([]*Stack, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Stack, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := stackFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func stacksFromHttpResponse(resp *http.Response) ([]*Stack, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return stacksFromJSON(body)
}

//endregion

//region Setters

//region Stack

func (o Stack) MarshalJSON() ([]byte, error) {
	type noMethod Stack
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Stack) SetIacType(v *string) *Stack {
	if o.IacType = v; o.IacType == nil {
		o.nullFields = append(o.nullFields, "IacType")
	}
	return o
}

func (o *Stack) SetNamespaceId(v *string) *Stack {
	if o.NamespaceId = v; o.NamespaceId == nil {
		o.nullFields = append(o.nullFields, "NamespaceId")
	}
	return o
}

func (o *Stack) SetName(v *string) *Stack {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Stack) SetDescription(v *string) *Stack {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *Stack) SetData(v *Data) *Stack {
	if o.Data = v; o.Data == nil {
		o.nullFields = append(o.nullFields, "Data")
	}
	return o
}

//endregion

//region Data

func (o Data) MarshalJSON() ([]byte, error) {
	type noMethod Data
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Data) SetDeploymentBehavior(v *cross_models.DeploymentBehavior) *Data {
	if o.DeploymentBehavior = v; o.DeploymentBehavior == nil {
		o.nullFields = append(o.nullFields, "DeploymentBehavior")
	}
	return o
}

func (o *Data) SetDeploymentApprovalPolicy(v *cross_models.DeploymentApprovalPolicy) *Data {
	if o.DeploymentApprovalPolicy = v; o.DeploymentApprovalPolicy == nil {
		o.nullFields = append(o.nullFields, "DeploymentApprovalPolicy")
	}
	return o
}

func (o *Data) SetVcsInfo(v *VcsInfo) *Data {
	if o.VcsInfo = v; o.VcsInfo == nil {
		o.nullFields = append(o.nullFields, "VcsInfo")
	}
	return o
}

func (o *Data) SetRunTrigger(v *cross_models.RunTrigger) *Data {
	if o.RunTrigger = v; o.RunTrigger == nil {
		o.nullFields = append(o.nullFields, "RunTrigger")
	}
	return o
}

func (o *Data) SetIacConfig(v *cross_models.IacConfig) *Data {
	if o.IacConfig = v; o.IacConfig == nil {
		o.nullFields = append(o.nullFields, "IacConfig")
	}
	return o
}

func (o *Data) SetPolicy(v *Policy) *Data {
	if o.Policy = v; o.Policy == nil {
		o.nullFields = append(o.nullFields, "Policy")
	}
	return o
}

func (o *Data) SetRunnerConfig(v *cross_models.RunnerConfig) *Data {
	if o.RunnerConfig = v; o.RunnerConfig == nil {
		o.nullFields = append(o.nullFields, "RunnerConfig")
	}
	return o
}

func (o *Data) SetAutoSync(v *cross_models.AutoSync) *Data {
	if o.AutoSync = v; o.AutoSync == nil {
		o.nullFields = append(o.nullFields, "AutoSync")
	}
	return o
}

//endregion

//region VCS Info

func (o VcsInfo) MarshalJSON() ([]byte, error) {
	type noMethod VcsInfo
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *VcsInfo) SetProviderId(v *string) *VcsInfo {
	if o.ProviderId = v; o.ProviderId == nil {
		o.nullFields = append(o.nullFields, "ProviderId")
	}
	return o
}

func (o *VcsInfo) SetRepoName(v *string) *VcsInfo {
	if o.RepoName = v; o.RepoName == nil {
		o.nullFields = append(o.nullFields, "RepoName")
	}
	return o
}

func (o *VcsInfo) SetPath(v *string) *VcsInfo {
	if o.Path = v; o.Path == nil {
		o.nullFields = append(o.nullFields, "Path")
	}
	return o
}

func (o *VcsInfo) SetBranch(v *string) *VcsInfo {
	if o.Branch = v; o.Branch == nil {
		o.nullFields = append(o.nullFields, "Branch")
	}
	return o
}

//endregion

//region Policy

func (o Policy) MarshalJSON() ([]byte, error) {
	type noMethod Policy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Policy) SetTtlConfig(v *TtlConfig) *Policy {
	if o.TtlConfig = v; o.TtlConfig == nil {
		o.nullFields = append(o.nullFields, "TtlConfig")
	}
	return o
}

func (o TtlConfig) MarshalJSON() ([]byte, error) {
	type noMethod TtlConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TtlConfig) SetTtl(v *TtlDefinition) *TtlConfig {
	if o.Ttl = v; o.Ttl == nil {
		o.nullFields = append(o.nullFields, "Ttl")
	}
	return o
}

func (o *TtlConfig) SetTtlOverride(v *TtlOverride) *TtlConfig {
	if o.TtlOverride = v; o.TtlOverride == nil {
		o.nullFields = append(o.nullFields, "TtlOverride")
	}
	return o
}

func (o TtlDefinition) MarshalJSON() ([]byte, error) {
	type noMethod TtlDefinition
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TtlDefinition) SetType(v *string) *TtlDefinition {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *TtlDefinition) SetValue(v *int) *TtlDefinition {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

func (o TtlOverride) MarshalJSON() ([]byte, error) {
	type noMethod TtlOverride
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TtlOverride) SetType(v *string) *TtlOverride {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *TtlOverride) SetValue(v *int) *TtlOverride {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//endregion

//endregion

// region Plan

type Plan struct {
	ID       *string `json:"id,omitempty"`
	Status   *string `json:"status,omitempty"`
	IsActive *bool   `json:"isActive,omitempty"`

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

type CreatePlanInput struct {
	StackId    *string `json:"stackId,omitempty"`
	HeadBranch *string `json:"headBranch,omitempty"`
	HeadSha    *string `json:"headSha,omitempty"`
}

type CreatePlanOutput struct {
	Plan *Plan `json:"plan,omitempty"`
}

type ReadPlanInput struct {
	PlanId *string `json:"planId,omitempty"`
}

type ReadPlanOutput struct {
	Plan *Plan `json:"plan,omitempty"`
}

func planFromJSON(in []byte) (*Plan, error) {
	b := new(Plan)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func plansFromJSON(in []byte) ([]*Plan, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Plan, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := planFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func plansFromHttpResponse(resp *http.Response) ([]*Plan, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return plansFromJSON(body)
}

func (s *ServiceOp) CreatePlan(ctx context.Context, input *CreatePlanInput) (*CreatePlanOutput, error) {
	r := client.NewRequest(http.MethodPost, "/stack/plan")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := plansFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreatePlanOutput)
	if len(gs) > 0 {
		output.Plan = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadPlan(ctx context.Context, input *ReadPlanInput) (*ReadPlanOutput, error) {
	path, err := uritemplates.Expand("/stack/plan/{planId}", uritemplates.Values{
		"planId": controlmonkey.StringValue(input.PlanId),
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

	gs, err := plansFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadPlanOutput)
	if len(gs) > 0 {
		output.Plan = gs[0]
	}

	return output, nil
}

func (o Plan) MarshalJSON() ([]byte, error) {
	type noMethod Plan
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Plan) SetID(v *string) *Plan {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

// endregion

//region Deployment

type Deployment struct {
	ID       *string `json:"id,omitempty"`
	Status   *string `json:"status,omitempty"`
	IsActive *bool   `json:"isActive,omitempty"`

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

type CreateDeploymentInput struct {
	StackId *string `json:"stackId,omitempty"`
}

type CreateDeploymentOutput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

type ReadDeploymentInput struct {
	DeploymentId *string `json:"deploymentId,omitempty"`
}

type ReadDeploymentOutput struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

func deploymentFromJSON(in []byte) (*Deployment, error) {
	b := new(Deployment)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func deploymentsFromJSON(in []byte) ([]*Deployment, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Deployment, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := deploymentFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func deploymentsFromHttpResponse(resp *http.Response) ([]*Deployment, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return deploymentsFromJSON(body)
}

func (s *ServiceOp) CreateDeployment(ctx context.Context, input *CreateDeploymentInput) (*CreateDeploymentOutput, error) {
	r := client.NewRequest(http.MethodPost, "/stack/deployment")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gs, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateDeploymentOutput)
	if len(gs) > 0 {
		output.Deployment = gs[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadDeployment(ctx context.Context, input *ReadDeploymentInput) (*ReadDeploymentOutput, error) {
	path, err := uritemplates.Expand("/stack/deployment/{deploymentId}", uritemplates.Values{
		"deploymentId": controlmonkey.StringValue(input.DeploymentId),
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

	gs, err := deploymentsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadDeploymentOutput)
	if len(gs) > 0 {
		output.Deployment = gs[0]
	}

	return output, nil
}

func (o Deployment) MarshalJSON() ([]byte, error) {
	type noMethod Deployment
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Deployment) SetID(v *string) *Deployment {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

//endregion
