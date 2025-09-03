package notification

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

//region NotificationSlackApp

//region Structure

type NotificationSlackApp struct {
	ID           *string `json:"id,omitempty"` // read-only
	Name         *string `json:"name,omitempty"`
	BotAuthToken *string `json:"botAuthToken,omitempty"`

	forceSendFields []string
	nullFields      []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateNotificationSlackApp(ctx context.Context, input *NotificationSlackApp) (*NotificationSlackApp, error) {
	r := client.NewRequest(http.MethodPost, baseUrl+slackAppUrl)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apps, err := notificationSlackAppsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(NotificationSlackApp)
	if len(apps) > 0 {
		output = apps[0]
	}

	return output, nil
}

func (s *ServiceOp) ListNotificationSlackApps(ctx context.Context, slackAppId *string, slackAppName *string) ([]*NotificationSlackApp, error) {
	r := client.NewRequest(http.MethodGet, baseUrl+slackAppUrl)

	if slackAppId != nil {
		r.Params.Set("slackAppId", *slackAppId)
	}
	if slackAppName != nil {
		r.Params.Set("slackAppName", *slackAppName)
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := notificationSlackAppsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) UpdateNotificationSlackApp(ctx context.Context, slackAppId string, input *NotificationSlackApp) (*NotificationSlackApp, error) {
	path, err := uritemplates.Expand(baseUrl+slackAppUrl+"/{slackAppId}", uritemplates.Values{"slackAppId": slackAppId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodPut, path)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apps, err := notificationSlackAppsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(NotificationSlackApp)
	if len(apps) > 0 {
		output = apps[0]
	}

	return output, nil
}

func (s *ServiceOp) DeleteNotificationSlackApp(ctx context.Context, slackAppId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand(baseUrl+slackAppUrl+"/{slackAppId}", uritemplates.Values{"slackAppId": slackAppId})
	if err != nil {
		return nil, err
	}

	r := client.NewRequest(http.MethodDelete, path)
	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output := new(commons.EmptyResponse)
	return output, nil
}

//endregion

//region Private Methods

func notificationSlackAppFromJSON(in []byte) (*NotificationSlackApp, error) {
	b := new(NotificationSlackApp)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func notificationSlackAppsFromJSON(in []byte) ([]*NotificationSlackApp, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*NotificationSlackApp, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := notificationSlackAppFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func notificationSlackAppsFromHttpResponse(resp *http.Response) ([]*NotificationSlackApp, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return notificationSlackAppsFromJSON(body)
}

//endregion

//region Setters

func (o NotificationSlackApp) MarshalJSON() ([]byte, error) {
	type noMethod NotificationSlackApp
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *NotificationSlackApp) SetName(v *string) *NotificationSlackApp {
	if o.Name = v; o.Name == nil {
		o.nullFields = append(o.nullFields, "Name")
	}
	return o
}

func (o *NotificationSlackApp) SetBotAuthToken(v *string) *NotificationSlackApp {
	if o.BotAuthToken = v; o.BotAuthToken == nil {
		o.nullFields = append(o.nullFields, "BotAuthToken")
	}
	return o
}

//endregion

//endregion
