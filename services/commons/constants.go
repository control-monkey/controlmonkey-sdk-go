package commons

const (
	OrganizationScope = "organization"
	NamespaceScope    = "namespace"
	TemplateScope     = "template"
	BlueprintScope    = "blueprint"
	StackScope        = "stack"

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

	Warning       = "warning"
	SoftMandatory = "softMandatory"
	HardMandatory = "hardMandatory"
	BySeverity    = "bySeverity"

	NamespaceRoleViewer   = "viewer"
	NamespaceRoleDeployer = "deployer"
	NamespaceRoleAdmin    = "admin"

	SlackProtocol = "slack"
	TeamsProtocol = "teams"
)
