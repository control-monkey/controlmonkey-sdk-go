package cross_models

type DeploymentApprovalPolicyRule struct {
	Type *string `json:"type,omitempty"` //commons.DeploymentApprovalPolicyRuleTypes

	forceSendFields []string
	nullFields      []string
}

func (o *DeploymentApprovalPolicyRule) SetType(v *string) *DeploymentApprovalPolicyRule {
	if o.Type = v; o.Type == nil {
		o.nullFields = append(o.nullFields, "Type")
	}
	return o
}
