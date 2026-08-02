package commons

const (
	OrganizationScope      = "organization"
	NamespaceScope         = "namespace"
	TemplateScope          = "template"
	BlueprintScope         = "blueprint"
	StackScope             = "stack"
	AwsAccountScope        = "awsAccount"
	AzureSubscriptionScope = "azureSubscription"
	GcpProjectScope        = "gcpProject"

	TfTVar = "tfVar"
	EnvVar = "envVar"

	RequireApproval     = "requireApproval"
	AutoApprove         = "autoApprove"
	RequireTwoApprovals = "requireTwoApprovals"

	GcpServiceAccount     = "gcpServiceAccount"
	AzureServicePrincipal = "azureServicePrincipal"
	AwsAssumeRole         = "awsAssumeRole"

	Terraform  = "terraform"
	Terragrunt = "terragrunt"
	Opentofu   = "opentofu"

	Allow    = "allow"
	Deny     = "deny"
	Extended = "extended"

	Managed    = "managed"
	SelfHosted = "selfHosted"

	Hours = "hours"
	Days  = "days"

	Ne         = "ne"
	Gt         = "gt"
	Gte        = "gte"
	Lt         = "lt"
	Lte        = "lte"
	In         = "in"
	StartsWith = "startsWith"
	Contains   = "contains"

	StackTargetType     = "stack"
	NamespaceTargetType = "namespace"

	AllTargetsIdentifier = "ALL"

	PostPlanStage = "postPlan"

	Warning       = "warning"
	SoftMandatory = "softMandatory"
	HardMandatory = "hardMandatory"
	BySeverity    = "bySeverity"

	NamespaceRoleViewer = "viewer"
	NamespaceRoleAdmin  = "admin"

	SlackProtocol    = "slack"
	SlackAppProtocol = "slackApp"
	TeamsProtocol    = "teams"
	EmailProtocol    = "email"

	WeeklyReportType = "weeklyReport"

	SeverityLow      = "low"
	SeverityMedium   = "medium"
	SeverityHigh     = "high"
	SeverityCritical = "critical"

	RoleViewer = "viewer"
	RoleAdmin  = "admin"

	NamespaceRoleType    = "namespaceRole"
	OrganizationRoleType = "organizationRole"

	RestrictReadToOwnStacks  = "restrictReadToOwnStacks"
	RestrictWriteToOwnStacks = "restrictWriteToOwnStacks"

	Default = "default"
	Manual  = "manual"

	BlueprintVariableManagedByStack  = "stack"
	BlueprintVariableManagedByInCode = "inCode"
)
