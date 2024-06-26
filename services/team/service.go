package team

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
	CreateTeam(context.Context, *Team) (*Team, error)
	ListTeams(context.Context, *string, *string) ([]*Team, error)
	ReadTeam(context.Context, string) (*Team, error)
	UpdateTeam(context.Context, string, *Team) (*Team, error)
	DeleteTeam(context.Context, string) (*commons.EmptyResponse, error)

	ListTeamUsers(context.Context, string) ([]*TeamUser, error)
	CreateTeamUser(context.Context, *TeamUser) (*commons.EmptyResponse, error)
	DeleteTeamUser(context.Context, *TeamUser) (*commons.EmptyResponse, error)
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
