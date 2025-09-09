package cross_models

import (
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
)

type DeploymentBehavior struct {
	DeployOnPush    *bool `json:"deployOnPush,omitempty"`
	WaitForApproval *bool `json:"waitForApproval,omitempty"`

	forceSendFields []string
	nullFields      []string
}

func (o DeploymentBehavior) MarshalJSON() ([]byte, error) {
	type noMethod DeploymentBehavior
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DeploymentBehavior) SetDeployOnPush(v *bool) *DeploymentBehavior {
	if o.DeployOnPush = v; o.DeployOnPush == nil {
		o.nullFields = append(o.nullFields, "DeployOnPush")
	}
	return o
}

func (o *DeploymentBehavior) SetWaitForApproval(v *bool) *DeploymentBehavior {
	if o.WaitForApproval = v; o.WaitForApproval == nil {
		o.nullFields = append(o.nullFields, "WaitForApproval")
	}
	return o
}
