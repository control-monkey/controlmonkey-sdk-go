package variable

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
	ListVariables(context.Context, *ListVariablesInput) (*ListVariablesOutput, error)
	CreateVariable(context.Context, *Variable) (*CreateVariableOutput, error)
	ReadVariable(context.Context, *ReadVariableInput) (*ReadVariableOutput, error)
	UpdateVariable(context.Context, *string, *Variable) (*UpdateVariableOutput, error)
	DeleteVariable(context.Context, *DeleteVariableInput) (*commons.EmptyResponse, error)
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
