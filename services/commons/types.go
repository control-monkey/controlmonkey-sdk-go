package commons

var (
	VariableScopeTypes                = []string{OrganizationScope, NamespaceScope, TemplateScope, BlueprintScope, StackScope}
	VariableTypes                     = []string{TfTVar, EnvVar}
	DeploymentApprovalPolicyRuleTypes = []string{RequireApproval, AutoApprove, RequireTwoApprovals}
	ExternalCredentialTypes           = []string{AwsAssumeRole, GcpServiceAccount, AzureServicePrincipal}
	IacTypes                          = []string{Terraform, Terragrunt}
	OverrideBehaviorTypes             = []string{Allow, Deny, Extended}
	RunnerConfigModeTypes             = []string{Managed, SelfHosted}
	TtlTypes                          = []string{Hours, Days}
	VariableConditionOperatorTypes    = []string{Ne, Gt, Gte, Lt, Lte, In, StartsWith, Contains}
)
