package template

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
	CreateTemplate(context.Context, *Template) (*Template, error)
	ListTemplates(context.Context, *string, *string) ([]*Template, error)
	ReadTemplate(context.Context, string) (*Template, error)
	UpdateTemplate(context.Context, string, *Template) (*Template, error)
	DeleteTemplate(context.Context, string) (*commons.EmptyResponse, error)

	ListTemplateNamespaceMappings(context.Context, string) ([]*TemplateNamespaceMapping, error)
	CreateTemplateNamespaceMapping(context.Context, *TemplateNamespaceMapping) (*TemplateNamespaceMapping, error)
	DeleteTemplateNamespaceMapping(context.Context, *TemplateNamespaceMapping) (*commons.EmptyResponse, error)
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
