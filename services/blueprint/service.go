package blueprint

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
	CreateBlueprint(context.Context, *Blueprint) (*Blueprint, error)
	ListBlueprints(context.Context, *string, *string) ([]*Blueprint, error)
	ReadBlueprint(context.Context, string) (*Blueprint, error)
	UpdateBlueprint(context.Context, string, *Blueprint) (*Blueprint, error)
	DeleteBlueprint(context.Context, string) (*commons.EmptyResponse, error)

	ListBlueprintNamespaceMappings(context.Context, string) ([]*BlueprintNamespaceMapping, error)
	CreateBlueprintNamespaceMapping(context.Context, *BlueprintNamespaceMapping) (*BlueprintNamespaceMapping, error)
	DeleteBlueprintNamespaceMapping(context.Context, *BlueprintNamespaceMapping) (*commons.EmptyResponse, error)
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
