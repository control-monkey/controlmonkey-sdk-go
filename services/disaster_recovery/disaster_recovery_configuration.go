package disaster_recovery

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

const (
	configurationUrl = baseUrl + "/configuration"
)

//region DisasterRecoveryConfiguration

// region Structure

type DisasterRecoveryConfiguration struct {
	ID             *string         `json:"id,omitempty"` // read-only
	Scope          *string         `json:"scope,omitempty"`
	CloudAccountId *string         `json:"cloudAccountId,omitempty"`
	BackupStrategy *BackupStrategy `json:"backupStrategy,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type BackupStrategy struct {
	IncludeManagedResources *bool                     `json:"includeManagedResources,omitempty"`
	Mode                    *string                   `json:"mode,omitempty"`
	VcsInfo                 *VcsInfo                  `json:"vcsInfo,omitempty"`
	Groups                  []*map[string]interface{} `json:"groups,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type VcsInfo struct {
	ProviderId *string `json:"vcsProviderId,omitempty"`
	RepoName   *string `json:"repoName,omitempty"`
	Branch     *string `json:"branch,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateDisasterRecoveryConfiguration(ctx context.Context, input *DisasterRecoveryConfiguration) (*DisasterRecoveryConfiguration, error) {
	r := client.NewRequest(http.MethodPost, configurationUrl)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	disasterRecoveryConfiguration, err := disasterRecoveryConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(DisasterRecoveryConfiguration)
	if len(disasterRecoveryConfiguration) > 0 {
		output = disasterRecoveryConfiguration[0]
	}

	return output, nil
}

func (s *ServiceOp) ReadDisasterRecoveryConfiguration(ctx context.Context, disasterRecoveryConfigurationId string) (*DisasterRecoveryConfiguration, error) {
	path, err := uritemplates.Expand(configurationUrl+"/{disasterRecoveryConfigurationId}", uritemplates.Values{"disasterRecoveryConfigurationId": disasterRecoveryConfigurationId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodGet, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	disasterRecoveryConfiguration, err := disasterRecoveryConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(DisasterRecoveryConfiguration)
	if len(disasterRecoveryConfiguration) > 0 {
		output = disasterRecoveryConfiguration[0]
	}

	return output, nil
}

func (s *ServiceOp) UpdateDisasterRecoveryConfiguration(ctx context.Context, id string, input *DisasterRecoveryConfiguration) (*DisasterRecoveryConfiguration, error) {
	path, err := uritemplates.Expand(configurationUrl+"/{disasterRecoveryConfigurationId}", uritemplates.Values{"disasterRecoveryConfigurationId": id})
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

	disasterRecoveryConfiguration, err := disasterRecoveryConfigurationsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(DisasterRecoveryConfiguration)

	if len(disasterRecoveryConfiguration) > 0 {
		output = disasterRecoveryConfiguration[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteDisasterRecoveryConfiguration(ctx context.Context, id string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand(configurationUrl+"/{disasterRecoveryConfigurationId}", uritemplates.Values{"disasterRecoveryConfigurationId": id})
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

func disasterRecoveryConfigurationFromJSON(in []byte) (*DisasterRecoveryConfiguration, error) {
	b := new(DisasterRecoveryConfiguration)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func disasterRecoveryConfigurationsFromJSON(in []byte) ([]*DisasterRecoveryConfiguration, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*DisasterRecoveryConfiguration, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := disasterRecoveryConfigurationFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func disasterRecoveryConfigurationsFromHttpResponse(resp *http.Response) ([]*DisasterRecoveryConfiguration, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return disasterRecoveryConfigurationsFromJSON(body)
}

//endregion

//region Setters

func (o DisasterRecoveryConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod DisasterRecoveryConfiguration
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DisasterRecoveryConfiguration) SetID(v *string) *DisasterRecoveryConfiguration {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *DisasterRecoveryConfiguration) SetScope(v *string) *DisasterRecoveryConfiguration {
	if o.Scope = v; o.Scope == nil {
		o.nullFields = append(o.nullFields, "Scope")
	}
	return o
}

func (o *DisasterRecoveryConfiguration) SetCloudAccountId(v *string) *DisasterRecoveryConfiguration {
	if o.CloudAccountId = v; o.CloudAccountId == nil {
		o.nullFields = append(o.nullFields, "CloudAccountId")
	}
	return o
}

func (o *DisasterRecoveryConfiguration) SetBackupStrategy(v *BackupStrategy) *DisasterRecoveryConfiguration {
	if o.BackupStrategy = v; o.BackupStrategy == nil {
		o.nullFields = append(o.nullFields, "BackupStrategy")
	}
	return o
}

//region BackupStrategy

func (o BackupStrategy) MarshalJSON() ([]byte, error) {
	type noMethod BackupStrategy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *BackupStrategy) SetIncludeManagedResources(v *bool) *BackupStrategy {
	if o.IncludeManagedResources = v; o.IncludeManagedResources == nil {
		o.nullFields = append(o.nullFields, "IncludeManagedResources")
	}
	return o
}

func (o *BackupStrategy) SetMode(v *string) *BackupStrategy {
	if o.Mode = v; o.Mode == nil {
		o.nullFields = append(o.nullFields, "Mode")
	}
	return o
}

func (o *BackupStrategy) SetVcsInfo(v *VcsInfo) *BackupStrategy {
	if o.VcsInfo = v; o.VcsInfo == nil {
		o.nullFields = append(o.nullFields, "VcsInfo")
	}
	return o
}

func (o *BackupStrategy) SetGroups(v []*map[string]interface{}) *BackupStrategy {
	if o.Groups = v; o.Groups == nil {
		o.nullFields = append(o.nullFields, "Groups")
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

func (o *VcsInfo) SetBranch(v *string) *VcsInfo {
	if o.Branch = v; o.Branch == nil {
		o.nullFields = append(o.nullFields, "Branch")
	}
	return o
}

//endregion

//endregion

//endregion

//endregion
