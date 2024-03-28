package control_policy_group

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
	CreateControlPolicyGroup(context.Context, *ControlPolicyGroup) (*ControlPolicyGroup, error)
	ReadControlPolicyGroup(context.Context, string) (*ControlPolicyGroup, error)
	UpdateControlPolicyGroup(context.Context, string, *ControlPolicyGroup) (*ControlPolicyGroup, error)
	DeleteControlPolicyGroup(context.Context, string) (*commons.EmptyResponse, error)

	CreateControlPolicyGroupMapping(context.Context, *ControlPolicyGroupMapping) (*ControlPolicyGroupMapping, error)
	ListControlPolicyGroupMappings(context.Context, string) ([]*ControlPolicyGroupMapping, error)
	UpdateControlPolicyGroupMapping(context.Context, *ControlPolicyGroupMapping) (*ControlPolicyGroupMapping, error)
	DeleteControlPolicyGroupMapping(context.Context, *ControlPolicyGroupMapping) (*commons.EmptyResponse, error)
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
