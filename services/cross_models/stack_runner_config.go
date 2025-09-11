package cross_models

import (
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
)

type RunnerConfig struct {
	Mode   *string   `json:"mode,omitempty"` //commons.RunnerConfigModeTypes
	Groups []*string `json:"groups,omitempty"`

	forceSendFields []string
	nullFields      []string
}

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
