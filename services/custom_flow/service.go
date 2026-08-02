package custom_flow

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
	CreateCustomFlow(context.Context, *CustomFlow) (*CustomFlow, error)
	ListCustomFlows(context.Context, *string, *string, *string, *string) ([]*CustomFlow, error)
	ReadCustomFlow(context.Context, string) (*CustomFlow, error)
	UpdateCustomFlow(context.Context, string, *CustomFlow) (*CustomFlow, error)
	DeleteCustomFlow(context.Context, string) (*commons.EmptyResponse, error)

	CreateCustomFlowMapping(context.Context, *CustomFlowMapping) (*CustomFlowMapping, error)
	ListCustomFlowMappings(context.Context) ([]*CustomFlowMapping, error)
	DeleteCustomFlowMapping(context.Context, *CustomFlowMapping) (*commons.EmptyResponse, error)
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
