package cross_models

import "github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"

type DeploymentApprovalPolicy struct {
	Rules []*DeploymentApprovalPolicyRule `json:"rules,omitempty"`

	forceSendFields []string
	nullFields      []string
}

type DeploymentApprovalPolicyRule struct {
	Type       *string                 `json:"type,omitempty"`
	Parameters *map[string]interface{} `json:"parameters,omitempty"`

	forceSendFields []string
	nullFields      []string
}

func (o DeploymentApprovalPolicy) MarshalJSON() ([]byte, error) {
	type noMethod DeploymentApprovalPolicy
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DeploymentApprovalPolicy) SetRules(v []*DeploymentApprovalPolicyRule) *DeploymentApprovalPolicy {
	if o.Rules = v; o.Rules == nil {
		o.nullFields = append(o.nullFields, "Rules")
	}
	return o
}

func (o DeploymentApprovalPolicyRule) MarshalJSON() ([]byte, error) {
	type noMethod DeploymentApprovalPolicyRule
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *DeploymentApprovalPolicyRule) SetType(v *string) *DeploymentApprovalPolicyRule {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}

func (o *DeploymentApprovalPolicyRule) SetParameters(v *map[string]interface{}) *DeploymentApprovalPolicyRule {
	if o.Parameters = v; o.Parameters == nil {
		o.nullFields = append(o.nullFields, "Parameters")
	}
	return o
}
