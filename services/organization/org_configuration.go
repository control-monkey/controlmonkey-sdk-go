package organization

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region OrgConfiguration

// region Structure

type OrgConfiguration struct {
	IacConfig             *IacConfig              `json:"iacConfig,omitempty"`
	S3StateFilesLocations []*S3StateFilesLocation `json:"tfStateFilesS3Locations,omitempty"`
	RunnerConfig          *RunnerConfig           `json:"runnerConfig,omitempty"`
	SuppressedResources   *SuppressedResources    `json:"suppressedResources,omitempty"`
	ReportConfigurations  []*ReportConfiguration  `json:"reportConfigurations,omitempty"`

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

type S3StateFilesLocation struct {
	BucketName   *string `json:"bucketName,omitempty"`
	BucketRegion *string `json:"bucketRegion,omitempty"`
	AwsAccountId *string `json:"awsAccountId,omitempty"`

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

type SuppressedResources struct {
	ManagedByTags []*TagProperties `json:"managedByTags,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type TagProperties struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ReportConfiguration struct {
	Type       *string           `json:"type,omitempty"` //commons.ReportTypes
	Recipients *ReportRecipients `json:"recipients,omitempty"`
	Enabled    *bool             `json:"enabled,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type ReportRecipients struct {
	AllAdmins               *bool     `json:"allAdmins,omitempty"`
	EmailAddresses          []*string `json:"emailAddresses,omitempty"`
	EmailAddressesToExclude []*string `json:"emailAddressesToExclude,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) ReadOrgConfiguration(ctx context.Context) (*OrgConfiguration, error) {
	r := client.NewRequest(http.MethodGet, baseUrl+configurationUrl)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	orgConfiguration, err := orgConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(OrgConfiguration)
	if len(orgConfiguration) > 0 {
		output = orgConfiguration[0]
	}

	return output, nil
}

func (s *ServiceOp) UpsertOrgConfiguration(ctx context.Context, input *OrgConfiguration) (*OrgConfiguration, error) {
	r := client.NewRequest(http.MethodPut, baseUrl+configurationUrl)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	orgConfiguration, err := orgConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(OrgConfiguration)
	if len(orgConfiguration) > 0 {
		output = orgConfiguration[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteOrgConfiguration(ctx context.Context) (*commons.EmptyResponse, error) {
	r := client.NewRequest(http.MethodDelete, baseUrl+configurationUrl)
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

func orgConfigurationFromJSON(in []byte) (*OrgConfiguration, error) {
	b := new(OrgConfiguration)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func orgConfigurationsFromJSON(in []byte) ([]*OrgConfiguration, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*OrgConfiguration, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := orgConfigurationFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func orgConfigurationsFromHttpResponse(resp *http.Response) ([]*OrgConfiguration, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return orgConfigurationsFromJSON(body)
}

//endregion

//region Setters

//region OrgConfiguration

func (o OrgConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod OrgConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *OrgConfiguration) SetIacConfig(v *IacConfig) *OrgConfiguration {
	if o.IacConfig = v; o.IacConfig == nil {
		o.nullFields = append(o.nullFields, "IacConfig")
	}
	return o
}

func (o *OrgConfiguration) SetTfStateFilesS3Locations(v []*S3StateFilesLocation) *OrgConfiguration {
	if o.S3StateFilesLocations = v; o.S3StateFilesLocations == nil {
		o.nullFields = append(o.nullFields, "S3StateFilesLocations")
	}
	return o
}

func (o *OrgConfiguration) SetRunnerConfig(v *RunnerConfig) *OrgConfiguration {
	if o.RunnerConfig = v; o.RunnerConfig == nil {
		o.nullFields = append(o.nullFields, "RunnerConfig")
	}
	return o
}

func (o *OrgConfiguration) SetSuppressedResources(v *SuppressedResources) *OrgConfiguration {
	if o.SuppressedResources = v; o.SuppressedResources == nil {
		o.nullFields = append(o.nullFields, "SuppressedResources")
	}
	return o
}

func (o *OrgConfiguration) SetReportConfigurations(v []*ReportConfiguration) *OrgConfiguration {
	if o.ReportConfigurations = v; o.ReportConfigurations == nil {
		o.nullFields = append(o.nullFields, "ReportConfigurations")
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

//region S3StateFilesLocation

func (o S3StateFilesLocation) MarshalJSON() ([]byte, error) {
	type noMethod S3StateFilesLocation
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *S3StateFilesLocation) SetBucketName(v *string) *S3StateFilesLocation {
	if o.BucketName = v; o.BucketName == nil {
		o.nullFields = append(o.nullFields, "BucketName")
	}
	return o
}

func (o *S3StateFilesLocation) SetBucketRegion(v *string) *S3StateFilesLocation {
	if o.BucketRegion = v; o.BucketRegion == nil {
		o.nullFields = append(o.nullFields, "BucketRegion")
	}
	return o
}

func (o *S3StateFilesLocation) SetAwsAccountId(v *string) *S3StateFilesLocation {
	if o.AwsAccountId = v; o.AwsAccountId == nil {
		o.nullFields = append(o.nullFields, "AwsAccountId")
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

//region SuppressedResources

func (o SuppressedResources) MarshalJSON() ([]byte, error) {
	type noMethod SuppressedResources
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *SuppressedResources) SetManagedByTags(v []*TagProperties) *SuppressedResources {
	if o.ManagedByTags = v; o.ManagedByTags == nil {
		o.nullFields = append(o.nullFields, "ManagedByTags")
	}
	return o
}

//region TagProperties

func (o TagProperties) MarshalJSON() ([]byte, error) {
	type noMethod TagProperties
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *TagProperties) SetKey(v *string) *TagProperties {
	if o.Key = v; o.Key == nil {
		o.nullFields = append(o.nullFields, "Key")
	}
	return o
}

func (o *TagProperties) SetValue(v *string) *TagProperties {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//endregion

//endregion

//region ReportConfiguration

func (o ReportConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod ReportConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ReportConfiguration) SetType(v *string) *ReportConfiguration {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *ReportConfiguration) SetRecipients(v *ReportRecipients) *ReportConfiguration {
	if o.Recipients = v; o.Recipients == nil {
		o.nullFields = append(o.nullFields, "Recipients")
	}
	return o
}

func (o *ReportConfiguration) SetEnabled(v *bool) *ReportConfiguration {
	if o.Enabled = v; o.Enabled == nil {
		o.nullFields = append(o.nullFields, "Enabled")
	}
	return o
}

//region ReportRecipients

func (o ReportRecipients) MarshalJSON() ([]byte, error) {
	type noMethod ReportRecipients
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *ReportRecipients) SetAllAdmins(v *bool) *ReportRecipients {
	if o.AllAdmins = v; o.AllAdmins == nil {
		o.nullFields = append(o.nullFields, "AllAdmins")
	}
	return o
}

func (o *ReportRecipients) SetEmailAddresses(v []*string) *ReportRecipients {
	if o.EmailAddresses = v; o.EmailAddresses == nil {
		o.nullFields = append(o.nullFields, "EmailAddresses")
	}
	return o
}

func (o *ReportRecipients) SetEmailAddressesToExclude(v []*string) *ReportRecipients {
	if o.EmailAddressesToExclude = v; o.EmailAddressesToExclude == nil {
		o.nullFields = append(o.nullFields, "EmailAddressesToExclude")
	}
	return o
}

//endregion

//endregion

//endregion

//endregion
