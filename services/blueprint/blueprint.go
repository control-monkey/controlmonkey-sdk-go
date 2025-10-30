package blueprint

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
)

//region Blueprint

// region Structure

type Blueprint struct {
	ID                               *string                `json:"id,omitempty"` // read-only
	Name                             *string                `json:"name,omitempty"`
	Description                      *string                `json:"description,omitempty"`
	BlueprintVcsInfo                 *VcsInfo               `json:"vcsInfo,omitempty"`
	StackConfiguration               *StackConfiguration    `json:"stackConfiguration,omitempty"`
	SubstituteParameters             []*SubstituteParameter `json:"substituteParameters,omitempty"`
	Policy                           *Policy                `json:"policy,omitempty"`
	SkipPlanOnStackInitialization    *bool                  `json:"skipPlanOnStackInitialization,omitempty"`
	AutoApproveApplyOnInitialization *bool                  `json:"autoApproveApplyOnInitialization,omitempty"`

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

type StackConfiguration struct {
	NamePattern              *string                                `json:"name,omitempty"`
	IacType                  *string                                `json:"iacType,omitempty"`
	VcsInfoWithPatterns      *StackVcsInfoWithPatterns              `json:"vcsInfo,omitempty"`
	DeploymentApprovalPolicy *cross_models.DeploymentApprovalPolicy `json:"deploymentApprovalPolicy,omitempty"`
	RunTrigger               *cross_models.RunTrigger               `json:"runTrigger,omitempty"`
	IacConfig                *cross_models.IacConfig                `json:"iacConfig,omitempty"`
	AutoSync                 *cross_models.AutoSync                 `json:"autoSync,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type StackVcsInfoWithPatterns struct {
	ProviderId    *string `json:"providerId,omitempty"`
	RepoName      *string `json:"repoName,omitempty"`
	PathPattern   *string `json:"path,omitempty"`
	BranchPattern *string `json:"branch,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type SubstituteParameter struct {
	Key             *string                   `json:"key,omitempty"`
	Description     *string                   `json:"description,omitempty"`
	ValueConditions []*cross_models.Condition `json:"valueConditions,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Policy struct {
	TtlConfig *TtlConfig `json:"ttlConfig,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TtlConfig struct {
	MaxTtl                        *TtlDefinition `json:"maxTtl,omitempty"`
	DefaultTtl                    *TtlDefinition `json:"defaultTtl,omitempty"`
	OpenCleanupPrOnTtlTermination *bool          `json:"openCleanupPrOnTtlTermination,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TtlDefinition struct {
	Type  *string `json:"type,omitempty"` //commons.TtlTypes
	Value *int    `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateBlueprint(ctx context.Context, input *Blueprint) (*Blueprint, error) {
	r := client.NewRequest(http.MethodPost, "/blueprint")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	blueprint, err := blueprintsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Blueprint)
	if len(blueprint) > 0 {
		output = blueprint[0]
	}

	return output, nil
}

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

func (s *ServiceOp) ReadBlueprint(ctx context.Context, blueprintId string) (*Blueprint, error) {
	path, err := uritemplates.Expand("/blueprint/{blueprintId}", uritemplates.Values{
		"blueprintId": blueprintId,
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

	blueprint, err := blueprintsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Blueprint)
	if len(blueprint) > 0 {
		output = blueprint[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateBlueprint(ctx context.Context, blueprintId string, input *Blueprint) (*Blueprint, error) {
	path, err := uritemplates.Expand("/blueprint/{blueprintId}", uritemplates.Values{"blueprintId": blueprintId})
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

	blueprint, err := blueprintsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Blueprint)
	if len(blueprint) > 0 {
		output = blueprint[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteBlueprint(ctx context.Context, blueprintId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/blueprint/{blueprintId}", uritemplates.Values{"blueprintId": blueprintId})
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

func (o *Blueprint) SetDescription(v *string) *Blueprint {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *Blueprint) SetBlueprintVcsInfo(v *VcsInfo) *Blueprint {
	if o.BlueprintVcsInfo = v; o.BlueprintVcsInfo == nil {
		o.nullFields = append(o.nullFields, "BlueprintVcsInfo")
	}
	return o
}

func (o *Blueprint) SetStackConfiguration(v *StackConfiguration) *Blueprint {
	if o.StackConfiguration = v; o.StackConfiguration == nil {
		o.nullFields = append(o.nullFields, "StackConfiguration")
	}
	return o
}

func (o *Blueprint) SetSubstituteParameters(v []*SubstituteParameter) *Blueprint {
	if o.SubstituteParameters = v; o.SubstituteParameters == nil {
		o.nullFields = append(o.nullFields, "SubstituteParameters")
	}
	return o
}

func (o *Blueprint) SetSkipPlanOnStackInitialization(v *bool) *Blueprint {
	if o.SkipPlanOnStackInitialization = v; o.SkipPlanOnStackInitialization == nil {
		o.nullFields = append(o.nullFields, "SkipPlanOnStackInitialization")
	}
	return o
}

func (o *Blueprint) SetAutoApproveApplyOnInitialization(v *bool) *Blueprint {
	if o.AutoApproveApplyOnInitialization = v; o.AutoApproveApplyOnInitialization == nil {
		o.nullFields = append(o.nullFields, "AutoApproveApplyOnInitialization")
	}
	return o
}

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

//region StackConfiguration

func (o StackConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod StackConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *StackConfiguration) SetNamePattern(v *string) *StackConfiguration {
	if o.NamePattern = v; o.NamePattern == nil {
		o.nullFields = append(o.nullFields, "NamePattern")
	}
	return o
}

func (o *StackConfiguration) SetIacType(v *string) *StackConfiguration {
	if o.IacType = v; o.IacType == nil {
		o.nullFields = append(o.nullFields, "IacType")
	}
	return o
}

func (o *StackConfiguration) SetVcsInfoWithPatterns(v *StackVcsInfoWithPatterns) *StackConfiguration {
	if o.VcsInfoWithPatterns = v; o.VcsInfoWithPatterns == nil {
		o.nullFields = append(o.nullFields, "VcsInfoWithPatterns")
	}
	return o
}

func (o *StackConfiguration) SetDeploymentApprovalPolicy(v *cross_models.DeploymentApprovalPolicy) *StackConfiguration {
	if o.DeploymentApprovalPolicy = v; o.DeploymentApprovalPolicy == nil {
		o.nullFields = append(o.nullFields, "DeploymentApprovalPolicy")
	}
	return o
}

func (o *StackConfiguration) SetRunTrigger(v *cross_models.RunTrigger) *StackConfiguration {
	if o.RunTrigger = v; o.RunTrigger == nil {
		o.nullFields = append(o.nullFields, "RunTrigger")
	}
	return o
}

func (o *StackConfiguration) SetIacConfig(v *cross_models.IacConfig) *StackConfiguration {
	if o.IacConfig = v; o.IacConfig == nil {
		o.nullFields = append(o.nullFields, "IacConfig")
	}
	return o
}

func (o *StackConfiguration) SetAutoSync(v *cross_models.AutoSync) *StackConfiguration {
	if o.AutoSync = v; o.AutoSync == nil {
		o.nullFields = append(o.nullFields, "AutoSync")
	}
	return o
}

//region StackVcsInfoWithPatterns

func (o StackVcsInfoWithPatterns) MarshalJSON() ([]byte, error) {
	type noMethod StackVcsInfoWithPatterns
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *StackVcsInfoWithPatterns) SetProviderId(v *string) *StackVcsInfoWithPatterns {
	if o.ProviderId = v; o.ProviderId == nil {
		o.nullFields = append(o.nullFields, "ProviderId")
	}
	return o
}

func (o *StackVcsInfoWithPatterns) SetRepoName(v *string) *StackVcsInfoWithPatterns {
	if o.RepoName = v; o.RepoName == nil {
		o.nullFields = append(o.nullFields, "RepoName")
	}
	return o
}

func (o *StackVcsInfoWithPatterns) SetPathPattern(v *string) *StackVcsInfoWithPatterns {
	if o.PathPattern = v; o.PathPattern == nil {
		o.nullFields = append(o.nullFields, "PathPattern")
	}
	return o
}

func (o *StackVcsInfoWithPatterns) SetBranchPattern(v *string) *StackVcsInfoWithPatterns {
	if o.BranchPattern = v; o.BranchPattern == nil {
		o.nullFields = append(o.nullFields, "BranchPattern")
	}
	return o
}

//endregion

//endregion

//region SubstituteParameter

func (o SubstituteParameter) MarshalJSON() ([]byte, error) {
	type noMethod SubstituteParameter
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *SubstituteParameter) SetKey(v *string) *SubstituteParameter {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *SubstituteParameter) SetDescription(v *string) *SubstituteParameter {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *SubstituteParameter) SetValueConditions(v []*cross_models.Condition) *SubstituteParameter {
	if o.ValueConditions = v; o.ValueConditions == nil {
		o.nullFields = append(o.nullFields, "ValueConditions")
	}
	return o
}

//endregion

// region Policy
func (o *Blueprint) SetPolicy(v *Policy) *Blueprint {
	if o.Policy = v; o.Policy == nil {
		o.nullFields = append(o.nullFields, "Policy")
	}
	return o
}

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

func (o *TtlConfig) SetMaxTtl(v *TtlDefinition) *TtlConfig {
	if o.MaxTtl = v; o.MaxTtl == nil {
		o.nullFields = append(o.nullFields, "MaxTtl")
	}
	return o
}

func (o *TtlConfig) SetDefaultTtl(v *TtlDefinition) *TtlConfig {
	if o.DefaultTtl = v; o.DefaultTtl == nil {
		o.nullFields = append(o.nullFields, "DefaultTtl")
	}
	return o
}

func (o *TtlConfig) SetOpenCleanupPrOnTtlTermination(v *bool) *TtlConfig {
	if o.OpenCleanupPrOnTtlTermination = v; o.OpenCleanupPrOnTtlTermination == nil {
		o.nullFields = append(o.nullFields, "OpenCleanupPrOnTtlTermination")
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

//endregion

//endregion
