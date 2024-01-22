package stack

import (
	"context"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

// Service provides the API operation methods for making requests to endpoints
// of the ControlMonkey API. See this package's package overview docs for details on
// the service.
type Service interface {
	ListStacks(context.Context, *ListStacksParams) (*ListStacksOutput, error)
	CreateStack(context.Context, *CreateStackInput) (*CreateStackOutput, error)
	ReadStack(context.Context, string) (*ReadStackOutput, error)
	UpdateStack(context.Context, string, *Stack) (*UpdateStackOutput, error)
	DeleteStack(context.Context, string) (*commons.EmptyResponse, error)

	CreateDeployment(context.Context, *CreateDeploymentInput) (*CreateDeploymentOutput, error)
	ReadDeployment(context.Context, *ReadDeploymentInput) (*ReadDeploymentOutput, error)

	CreatePlan(context.Context, *CreatePlanInput) (*CreatePlanOutput, error)
	ReadPlan(context.Context, *ReadPlanInput) (*ReadPlanOutput, error)
}

type ServiceOp struct {
	Client *client.Client
}

var _ Service = &ServiceOp{}

func New(sess *session.Session, cfgs ...*controlmonkey.Config) Service {
	cfg := &controlmonkey.Config{}
	cfg.Merge(sess.Config)
	cfg.Merge(cfgs...)

	return &ServiceOp{
		Client: client.New(sess.Config),
	}
}
