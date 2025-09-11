package template

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/services/cross_models"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region Template

// region Structure

type Template struct {
	ID                        *string                    `json:"id,omitempty"` // read-only
	Name                      *string                    `json:"name,omitempty"`
	IacType                   *string                    `json:"iacType,omitempty"` //commons.IacTypes
	Description               *string                    `json:"description,omitempty"`
	VcsInfo                   *VcsInfo                   `json:"vcsInfo,omitempty"`
	Policy                    *Policy                    `json:"policy,omitempty"`
	SkipStateRefreshOnDestroy *bool                      `json:"skipStateRefreshOnDestroy,omitempty"`
	IacConfig                 *IacConfig                 `json:"iacConfig,omitempty"`
	RunnerConfig              *cross_models.RunnerConfig `json:"runnerConfig,omitempty"`

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
	OpentofuVersion   *string `json:"opentofuVersion,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Requests & Responses

type ListTemplatesParams struct {
	NamespaceId *string
}

//endregion

//region Methods

func (s *ServiceOp) CreateTemplate(ctx context.Context, input *Template) (*Template, error) {
	r := client.NewRequest(http.MethodPost, "/template")
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	template, err := templatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Template)
	if len(template) > 0 {
		output = template[0]
	}

	return output, nil
}

func (s *ServiceOp) ListTemplates(ctx context.Context, templateId *string, templateName *string) ([]*Template, error) {
	r := client.NewRequest(http.MethodGet, "/template")

	if templateId != nil {
		r.Params.Set("templateId", *templateId)
	}
	if templateName != nil {
		r.Params.Set("templateName", *templateName)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	templates, err := templatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return templates, nil
}

func (s *ServiceOp) ReadTemplate(ctx context.Context, templateId string) (*Template, error) {
	path, err := uritemplates.Expand("/template/{templateId}", uritemplates.Values{
		"templateId": templateId,
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

	template, err := templatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Template)
	if len(template) > 0 {
		output = template[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateTemplate(ctx context.Context, templateId string, input *Template) (*Template, error) {
	path, err := uritemplates.Expand("/template/{templateId}", uritemplates.Values{"templateId": templateId})
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

	template, err := templatesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(Template)
	if len(template) > 0 {
		output = template[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteTemplate(ctx context.Context, templateId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand("/template/{templateId}", uritemplates.Values{"templateId": templateId})
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

func templateFromJSON(in []byte) (*Template, error) {
	b := new(Template)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func templatesFromJSON(in []byte) ([]*Template, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Template, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := templateFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func templatesFromHttpResponse(resp *http.Response) ([]*Template, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return templatesFromJSON(body)
}

//endregion

//region Setters

//region Template

func (o Template) MarshalJSON() ([]byte, error) {
	type noMethod Template
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Template) SetName(v *string) *Template {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *Template) SetIacType(v *string) *Template {
	if o.IacType = v; o.IacType == nil {
		o.nullFields = append(o.nullFields, "IacType")
	}
	return o
}

func (o *Template) SetDescription(v *string) *Template {
	if o.Description = v; o.Description == nil {
		o.nullFields = append(o.nullFields, "Description")
	}
	return o
}

func (o *Template) SetVcsInfo(v *VcsInfo) *Template {
	if o.VcsInfo = v; o.VcsInfo == nil {
		o.nullFields = append(o.nullFields, "VcsInfo")
	}
	return o
}

func (o *Template) SetPolicy(v *Policy) *Template {
	if o.Policy = v; o.Policy == nil {
		o.nullFields = append(o.nullFields, "Policy")
	}
	return o
}

func (o *Template) SetSkipStateRefreshOnDestroy(v *bool) *Template {
	if o.SkipStateRefreshOnDestroy = v; o.SkipStateRefreshOnDestroy == nil {
		o.nullFields = append(o.nullFields, "SkipStateRefreshOnDestroy")
	}
	return o
}

func (o *Template) SetIacConfig(v *IacConfig) *Template {
	if o.IacConfig = v; o.IacConfig == nil {
		o.nullFields = append(o.nullFields, "IacConfig")
	}
	return o
}

func (o *Template) SetRunnerConfig(v *cross_models.RunnerConfig) *Template {
	if o.RunnerConfig = v; o.RunnerConfig == nil {
		o.nullFields = append(o.nullFields, "RunnerConfig")
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

func (o *IacConfig) SetOpentofuVersion(v *string) *IacConfig {
	if o.OpentofuVersion = v; o.OpentofuVersion == nil {
		o.nullFields = append(o.nullFields, "OpentofuVersion")
	}
	return o
}

//endregion

//endregion
