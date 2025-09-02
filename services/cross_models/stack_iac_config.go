package cross_models

import "github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"

type IacConfig struct {
	TerraformVersion   *string   `json:"terraformVersion,omitempty"`
	TerragruntVersion  *string   `json:"terragruntVersion,omitempty"`
	OpentofuVersion    *string   `json:"opentofuVersion,omitempty"`
	IsTerragruntRunAll *bool     `json:"isTerragruntRunAll,omitempty"`
	VarFiles           []*string `json:"varFiles,omitempty"`

	forceSendFields []string
	nullFields      []string
}

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

func (o *IacConfig) SetIsTerragruntRunAll(v *bool) *IacConfig {
	if o.IsTerragruntRunAll = v; o.IsTerragruntRunAll == nil {
		o.nullFields = append(o.nullFields, "IsTerragruntRunAll")
	}
	return o
}

func (o *IacConfig) SetVarFiles(v []*string) *IacConfig {
	if o.VarFiles = v; o.VarFiles == nil {
		o.nullFields = append(o.nullFields, "VarFiles")
	}
	return o
}
