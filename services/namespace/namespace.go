package namespace

import (
	"context"
	"encoding/json"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"
	"io"
	"net/http"
)

//region Namespace

// region Structure

type Namespace struct {
	ID                       *string                   `json:"id,omitempty"` // read-only
	Name                     *string                   `json:"name,omitempty"`
	Description              *string                   `json:"description,omitempty"`
	ExternalCredentials      []*ExternalCredentials    `json:"externalCredentials,omitempty"`
	Policy                   *Policy                   `json:"policy,omitempty"`
	IacConfig                *IacConfig                `json:"iacConfig,omitempty"`
	RunnerConfig             *RunnerConfig             `json:"runnerConfig,omitempty"`
	DeploymentApprovalPolicy *DeploymentApprovalPolicy `json:"deploymentApprovalPolicy,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ExternalCredentials struct {
	Type                  *string `json:"type,omitempty"` //commons.ExternalCredentialTypes
	ExternalCredentialsId *string `json:"externalCredentialsId,omitempty"`
	AwsProfileName        *string `json:"awsProfileName,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type Policy struct {
	TtlConfig *TtlConfig `json:"ttlConfig,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TtlConfig struct {
	MaxTtl     *TtlDefinition `json:"maxTtl,omitempty"`
	DefaultTtl *TtlDefinition `json:"defaultTtl,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TtlDefinition struct {
	Type  *string `json:"type,omitempty"` //commons.TtlTypes
	Value *int    `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type IacConfig struct {
	TerraformVersion  *string `json:"terraformVersion,omitempty"`
	TerragruntVersion *string `json:"terragruntVersion,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type RunnerConfig struct {
	Mode          *string   `json:"mode,omitempty"` //commons.RunnerConfigModeTypes
	Groups        []*string `json:"groups,omitempty"`
	IsOverridable *bool     `json:"isOverridable,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DeploymentApprovalPolicy struct {
	Rules            []*cross_models.DeploymentApprovalPolicyRule `json:"rules,omitempty"`
	OverrideBehavior *string                                      `json:"overrideBehavior,omitempty"` //commons.OverrideBehaviorTypes

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Requests & Responses

type ListNamespacesOutput struct {
	Namespaces []*Namespace `json:"namespaces,omitempty"`
}

type CreateNamespaceOutput struct {
	Namespace *Namespace `json:"namespace,omitempty"`
}

type ReadNamespaceOutput struct {
	Namespace *Namespace `json:"namespace,omitempty"`
}

type UpdateNamespaceOutput struct {
	Namespace *Namespace `json:"namespace,omitempty"`
}

//endregion

//region Methods

func (s *ServiceOp) ListNamespaces(ctx context.Context) (*ListNamespacesOutput, error) {
	r := client.NewRequest(http.MethodGet, "/namespace")

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	namespaces, err := namespacesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListNamespacesOutput{Namespaces: namespaces}, nil
}

func (s *ServiceOp) CreateNamespace(ctx context.Context, input *Namespace) (*CreateNamespaceOutput, error) {
	r := client.NewRequest(http.MethodPost, "/namespace")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	namespace, err := namespacesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(CreateNamespaceOutput)
	if len(namespace) > 0 {
		output.Namespace = namespace[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadNamespace(ctx context.Context, namespaceId string) (*ReadNamespaceOutput, error) {
	path, err := uritemplates.Expand("/namespace/{namespaceId}", uritemplates.Values{"namespaceId": namespaceId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	namespace, err := namespacesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(ReadNamespaceOutput)
	if len(namespace) > 0 {
		output.Namespace = namespace[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateNamespace(ctx context.Context, namespaceId string, input *Namespace) (*UpdateNamespaceOutput, error) {
	path, err := uritemplates.Expand("/namespace/{namespaceId}", uritemplates.Values{"namespaceId": namespaceId})
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

	namespace, err := namespacesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(UpdateNamespaceOutput)
	if len(namespace) > 0 {
		output.Namespace = namespace[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteNamespace(ctx context.Context, namespaceId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/namespace/{namespaceId}", uritemplates.Values{"namespaceId": namespaceId})
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

func namespaceFromJSON(in []byte) (*Namespace, error) {
	b := new(Namespace)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func namespacesFromJSON(in []byte) ([]*Namespace, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Namespace, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := namespaceFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func namespacesFromHttpResponse(resp *http.Response) ([]*Namespace, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return namespacesFromJSON(body)
}

//endregion

//region Setters

//region Namespace

func (o Namespace) MarshalJSON() ([]byte, error) {
	type noMethod Namespace
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Namespace) SetID(v *string) *Namespace {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *Namespace) SetName(v *string) *Namespace {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Namespace) SetDescription(v *string) *Namespace {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *Namespace) SetExternalCredentials(v []*ExternalCredentials) *Namespace {
	if o.ExternalCredentials = v; o.ExternalCredentials == nil {
		o.nullFields = append(o.nullFields, "ExternalCredentials")
	}
	return o
}

func (o *Namespace) SetPolicy(v *Policy) *Namespace {
	if o.Policy = v; o.Policy == nil {
		o.nullFields = append(o.nullFields, "Policy")
	}
	return o
}

func (o *Namespace) SetIacConfig(v *IacConfig) *Namespace {
	if o.IacConfig = v; o.IacConfig == nil {
		o.nullFields = append(o.nullFields, "IacConfig")
	}
	return o
}

func (o *Namespace) SetRunnerConfig(v *RunnerConfig) *Namespace {
	if o.RunnerConfig = v; o.RunnerConfig == nil {
		o.nullFields = append(o.nullFields, "RunnerConfig")
	}
	return o
}

func (o *Namespace) SetDeploymentApprovalPolicy(v *DeploymentApprovalPolicy) *Namespace {
	if o.DeploymentApprovalPolicy = v; o.DeploymentApprovalPolicy == nil {
		o.nullFields = append(o.nullFields, "DeploymentApprovalPolicy")
	}
	return o
}

//endregion

//region ExternalCredentials

func (o ExternalCredentials) MarshalJSON() ([]byte, error) {
	type noMethod ExternalCredentials
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ExternalCredentials) SetType(v *string) *ExternalCredentials {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *ExternalCredentials) SetExternalCredentialsId(v *string) *ExternalCredentials {
	if o.ExternalCredentialsId = v; o.ExternalCredentialsId == nil {
		o.nullFields = append(o.nullFields, "ExternalCredentialsId")
	}
	return o
}

func (o *ExternalCredentials) SetAwsProfileName(v *string) *ExternalCredentials {
	if o.AwsProfileName = v; o.AwsProfileName == nil {
		o.nullFields = append(o.nullFields, "AwsProfileName")
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

//region Iac Config

func (o IacConfig) MarshalJSON() ([]byte, error) {
	type noMethod IacConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *IacConfig) SetTerraformVersion(v *string) *IacConfig {
	if o.TerraformVersion = v; o.TerraformVersion == nil {
		o.nullFields = append(o.nullFields, "TerraformVersion")
	}
	return o
}

func (o *IacConfig) SetTerragruntVersion(v *string) *IacConfig {
	if o.TerragruntVersion = v; o.TerragruntVersion == nil {
		o.nullFields = append(o.nullFields, "TerragruntVersion")
	}
	return o
}

//endregion

//region Runner Config

func (o RunnerConfig) MarshalJSON() ([]byte, error) {
	type noMethod RunnerConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RunnerConfig) SetMode(v *string) *RunnerConfig {
	if o.Mode = v; o.Mode == nil {
		o.nullFields = append(o.nullFields, "Mode")
	}
	return o
}

func (o *RunnerConfig) SetGroups(v []*string) *RunnerConfig {
	if o.Groups = v; o.Groups == nil {
		o.nullFields = append(o.nullFields, "Groups")
	}
	return o
}

func (o *RunnerConfig) SetIsOverridable(v *bool) *RunnerConfig {
	if o.IsOverridable = v; o.IsOverridable == nil {
		o.nullFields = append(o.nullFields, "IsOverridable")
	}
	return o
}

//endregion

//region Deployment Approval Policy

func (o DeploymentApprovalPolicy) MarshalJSON() ([]byte, error) {
	type noMethod DeploymentApprovalPolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DeploymentApprovalPolicy) SetRules(v []*cross_models.DeploymentApprovalPolicyRule) *DeploymentApprovalPolicy {
	if o.Rules = v; o.Rules == nil {
		o.nullFields = append(o.nullFields, "Rules")
	}
	return o
}

func (o *DeploymentApprovalPolicy) SetOverrideBehavior(v *string) *DeploymentApprovalPolicy {
	if o.OverrideBehavior = v; o.OverrideBehavior == nil {
		o.nullFields = append(o.nullFields, "OverrideBehavior")
	}
	return o
}

//endregion

//endregion

//endregion
