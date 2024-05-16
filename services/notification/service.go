package notification

import (
	"context"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/session"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

const (
	baseUrl         = "/notification"
	endpointUrl     = "/endpoint"
	subscriptionUrl = "/subscription"
)

// Service provides the API operation methods for making requests to endpoints
// of the ControlMonkey API. See this package's package overview docs for details on
// the service.
type Service interface {
	CreateNotificationEndpoint(context.Context, *Endpoint) (*Endpoint, error)
	ListNotificationEndpoints(context.Context, *string, *string) ([]*Endpoint, error)
	ReadNotificationEndpoint(context.Context, string) (*Endpoint, error)
	UpdateNotificationEndpoint(context.Context, string, *Endpoint) (*Endpoint, error)
	DeleteNotificationEndpoint(context.Context, string) (*commons.EmptyResponse, error)

	ListEventSubscriptions(context.Context, string, *string) ([]*EventSubscription, error)
	CreateEventSubscription(context.Context, *EventSubscription) (*EventSubscription, error)
	DeleteEventSubscription(context.Context, string) (*commons.EmptyResponse, error)
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
