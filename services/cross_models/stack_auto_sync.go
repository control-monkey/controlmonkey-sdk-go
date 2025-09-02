package cross_models

import "github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"

type AutoSync struct {
	DeployWhenDriftDetected *bool `json:"deployWhenDriftDetected,omitempty"`

	forceSendFields []string
	nullFields      []string
}

func (o AutoSync) MarshalJSON() ([]byte, error) {
	type noMethod AutoSync
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *AutoSync) SetDeployWhenDriftDetected(v *bool) *AutoSync {
	if o.DeployWhenDriftDetected = v; o.DeployWhenDriftDetected == nil {
		o.nullFields = append(o.nullFields, "DeployWhenDriftDetected")
	}
	return o
}
