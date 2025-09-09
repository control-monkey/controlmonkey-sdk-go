package stack_discovery_configuration

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"
)

// region StackDiscoveryConfiguration

// region Structure

type StackDiscoveryConfiguration struct {
	ID          *string       `json:"id,omitempty"` // read-only
	Name        *string       `json:"name,omitempty"`
	NamespaceId *string       `json:"namespaceId,omitempty"`
	Description *string       `json:"description,omitempty"`
	VcsPatterns []*VcsPattern `json:"vcsPatterns,omitempty"`
	StackConfig *StackConfig  `json:"stackConfig,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VcsPattern struct {
	ProviderId          *string   `json:"providerId,omitempty"`
	RepoName            *string   `json:"repoName,omitempty"`
	PathPatterns        []*string `json:"pathPatterns,omitempty"`
	ExcludePathPatterns []*string `json:"excludePathPatterns,omitempty"`
	Branch              *string   `json:"branch,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type StackConfig struct {
	IacType                  *string                                `json:"iacType,omitempty"` //commons.IacTypes
	DeploymentBehavior       *cross_models.DeploymentBehavior       `json:"deploymentBehavior,omitempty"`
	DeploymentApprovalPolicy *cross_models.DeploymentApprovalPolicy `json:"deploymentApprovalPolicy,omitempty"`
	RunTrigger               *cross_models.RunTrigger               `json:"runTrigger,omitempty"`
	IacConfig                *cross_models.IacConfig                `json:"iacConfig,omitempty"`
	RunnerConfig             *cross_models.RunnerConfig             `json:"runnerConfig,omitempty"`
	AutoSync                 *cross_models.AutoSync                 `json:"autoSync,omitempty"`

	forceSendFields []string
	nullFields      []string
}

// endregion

// region Methods

func (s *ServiceOp) CreateStackDiscoveryConfiguration(ctx context.Context, input *StackDiscoveryConfiguration) (*StackDiscoveryConfiguration, error) {
	r := client.NewRequest(http.MethodPost, "/stackDiscoveryConfiguration")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	items, err := stackDiscoveryConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	var out *StackDiscoveryConfiguration
	if len(items) > 0 {
		out = items[0]
	} else {
		out = new(StackDiscoveryConfiguration)
	}
	return out, nil
}

func (s *ServiceOp) ReadStackDiscoveryConfiguration(ctx context.Context, configId string) (*StackDiscoveryConfiguration, error) {
	path, err := uritemplates.Expand("/stackDiscoveryConfiguration/{configId}", uritemplates.Values{"configId": configId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	items, err := stackDiscoveryConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	var out *StackDiscoveryConfiguration
	if len(items) > 0 {
		out = items[0]
	} else {
		out = new(StackDiscoveryConfiguration)
	}
	return out, nil
}

func (s *ServiceOp) UpdateStackDiscoveryConfiguration(ctx context.Context, configId string, input *StackDiscoveryConfiguration) (*StackDiscoveryConfiguration, error) {
	path, err := uritemplates.Expand("/stackDiscoveryConfiguration/{configId}", uritemplates.Values{"configId": configId})
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

	items, err := stackDiscoveryConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	var out *StackDiscoveryConfiguration
	if len(items) > 0 {
		out = items[0]
	} else {
		out = new(StackDiscoveryConfiguration)
	}
	return out, nil
}

func (s *ServiceOp) DeleteStackDiscoveryConfiguration(ctx context.Context, configId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/stackDiscoveryConfiguration/{configId}", uritemplates.Values{"configId": configId})
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

func stackDiscoveryConfigurationFromJSON(in []byte) (*StackDiscoveryConfiguration, error) {
	b := new(StackDiscoveryConfiguration)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func stackDiscoveryConfigurationsFromJSON(in []byte) ([]*StackDiscoveryConfiguration, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*StackDiscoveryConfiguration, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := stackDiscoveryConfigurationFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func stackDiscoveryConfigurationsFromHttpResponse(resp *http.Response) ([]*StackDiscoveryConfiguration, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return stackDiscoveryConfigurationsFromJSON(body)
}

// endregion

// region Setters

func (o StackDiscoveryConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod StackDiscoveryConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *StackDiscoveryConfiguration) SetName(v *string) *StackDiscoveryConfiguration {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *StackDiscoveryConfiguration) SetNamespaceId(v *string) *StackDiscoveryConfiguration {
	if o.NamespaceId = v; o.NamespaceId == nil {
		o.nullFields = append(o.nullFields, "NamespaceId")
	}
	return o
}

func (o *StackDiscoveryConfiguration) SetDescription(v *string) *StackDiscoveryConfiguration {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *StackDiscoveryConfiguration) SetVcsPatterns(v []*VcsPattern) *StackDiscoveryConfiguration {
	if o.VcsPatterns = v; o.VcsPatterns == nil {
		o.nullFields = append(o.nullFields, "VcsPatterns")
	}
	return o
}

func (o *StackDiscoveryConfiguration) SetStackConfig(v *StackConfig) *StackDiscoveryConfiguration {
	if o.StackConfig = v; o.StackConfig == nil {
		o.nullFields = append(o.nullFields, "StackConfig")
	}
	return o
}

// region VcsPattern

func (o VcsPattern) MarshalJSON() ([]byte, error) {
	type noMethod VcsPattern
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *VcsPattern) SetProviderId(v *string) *VcsPattern {
	if o.ProviderId = v; o.ProviderId == nil {
		o.nullFields = append(o.nullFields, "ProviderId")
	}
	return o
}

func (o *VcsPattern) SetRepoName(v *string) *VcsPattern {
	if o.RepoName = v; o.RepoName == nil {
		o.nullFields = append(o.nullFields, "RepoName")
	}
	return o
}

func (o *VcsPattern) SetPathPatterns(v []*string) *VcsPattern {
	if o.PathPatterns = v; o.PathPatterns == nil {
		o.nullFields = append(o.nullFields, "PathPatterns")
	}
	return o
}

func (o *VcsPattern) SetExcludePathPatterns(v []*string) *VcsPattern {
	if o.ExcludePathPatterns = v; o.ExcludePathPatterns == nil {
		o.nullFields = append(o.nullFields, "ExcludePathPatterns")
	}
	return o
}

func (o *VcsPattern) SetBranch(v *string) *VcsPattern {
	if o.Branch = v; o.Branch == nil {
		o.nullFields = append(o.nullFields, "Branch")
	}
	return o
}

// endregion

// region StackConfig

func (o StackConfig) MarshalJSON() ([]byte, error) {
	type noMethod StackConfig
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *StackConfig) SetIacType(v *string) *StackConfig {
	if o.IacType = v; o.IacType == nil {
		o.nullFields = append(o.nullFields, "IacType")
	}
	return o
}

func (o *StackConfig) SetDeploymentBehavior(v *cross_models.DeploymentBehavior) *StackConfig {
	if o.DeploymentBehavior = v; o.DeploymentBehavior == nil {
		o.nullFields = append(o.nullFields, "DeploymentBehavior")
	}
	return o
}

func (o *StackConfig) SetDeploymentApprovalPolicy(v *cross_models.DeploymentApprovalPolicy) *StackConfig {
	if o.DeploymentApprovalPolicy = v; o.DeploymentApprovalPolicy == nil {
		o.nullFields = append(o.nullFields, "DeploymentApprovalPolicy")
	}
	return o
}

func (o *StackConfig) SetRunTrigger(v *cross_models.RunTrigger) *StackConfig {
	if o.RunTrigger = v; o.RunTrigger == nil {
		o.nullFields = append(o.nullFields, "RunTrigger")
	}
	return o
}

func (o *StackConfig) SetIacConfig(v *cross_models.IacConfig) *StackConfig {
	if o.IacConfig = v; o.IacConfig == nil {
		o.nullFields = append(o.nullFields, "IacConfig")
	}
	return o
}

func (o *StackConfig) SetRunnerConfig(v *cross_models.RunnerConfig) *StackConfig {
	if o.RunnerConfig = v; o.RunnerConfig == nil {
		o.nullFields = append(o.nullFields, "RunnerConfig")
	}
	return o
}

func (o *StackConfig) SetAutoSync(v *cross_models.AutoSync) *StackConfig {
	if o.AutoSync = v; o.AutoSync == nil {
		o.nullFields = append(o.nullFields, "AutoSync")
	}
	return o
}

// endregion

// endregion

// endregion
