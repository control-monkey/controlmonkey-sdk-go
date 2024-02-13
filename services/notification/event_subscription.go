package notification

import (
	"context"
	"encoding/json"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/client"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/uritemplates"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
	"io"
	"net/http"
)

//region EventSubscription

//region Structure

type EventSubscription struct {
	ID                     *string `json:"id,omitempty"` // read-only
	NotificationEndpointId *string `json:"notificationEndpointId,omitempty"`
	Scope                  *string `json:"scope,omitempty"`
	ScopeId                *string `json:"scopeId,omitempty"`
	EventType              *string `json:"eventType,omitempty"`

	// forceSendFields is a read of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	forceSendFields []string

	// nullFields is a read of field names (e.g. "Keys") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this read has a non-empty value.
	// This may be used to include null fields in Patch requests.
	nullFields []string
}

//endregion

//region Methods

func (s *ServiceOp) CreateEventSubscription(ctx context.Context, input *EventSubscription) (*EventSubscription, error) {
	r := client.NewRequest(http.MethodPost, baseUrl+subscriptionUrl)
	r.Obj = input

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	endpoint, err := eventSubscriptionsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	output := new(EventSubscription)
	if len(endpoint) > 0 {
		output = endpoint[0]
	}

	return output, nil
}

func (s *ServiceOp) ListEventSubscriptions(ctx context.Context, scope string, scopeId *string) ([]*EventSubscription, error) {
	r := client.NewRequest(http.MethodGet, baseUrl+subscriptionUrl)
	if scope == commons.OrganizationScope {
		r.Params.Set("orgOnly", "true")
	} else if scopeId != nil {
		if scope == commons.NamespaceScope {
			r.Params.Set("namespaceId", *scopeId)
		}
	}

	resp, err := client.RequireOK(s.Client.Do(ctx, r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := eventSubscriptionsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *ServiceOp) DeleteEventSubscription(ctx context.Context, subscriptionId string) (*commons.EmptyResponse, error) {
	path, err := uritemplates.Expand(baseUrl+subscriptionUrl+"/{subscriptionId}", uritemplates.Values{"subscriptionId": subscriptionId})
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

func eventSubscriptionFromJSON(in []byte) (*EventSubscription, error) {
	b := new(EventSubscription)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func eventSubscriptionsFromJSON(in []byte) ([]*EventSubscription, error) {
	var rw client.Response
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*EventSubscription, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := eventSubscriptionFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}

	return out, nil
}

func eventSubscriptionsFromHttpResponse(resp *http.Response) ([]*EventSubscription, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return eventSubscriptionsFromJSON(body)
}

//endregion

//region Setters

func (o EventSubscription) MarshalJSON() ([]byte, error) {
	type noMethod EventSubscription
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *EventSubscription) SetID(v *string) *EventSubscription {
	if o.ID = v; o.ID == nil {
		o.nullFields = append(o.nullFields, "ID")
	}
	return o
}

func (o *EventSubscription) SetNotificationEndpointId(v *string) *EventSubscription {
	if o.NotificationEndpointId = v; o.NotificationEndpointId == nil {
		o.nullFields = append(o.nullFields, "NotificationEndpointId")
	}
	return o
}

func (o *EventSubscription) SetScope(v *string) *EventSubscription {
	if o.Scope = v; o.Scope == nil {
		o.nullFields = append(o.nullFields, "Scope")
	}
	return o
}

func (o *EventSubscription) SetScopeId(v *string) *EventSubscription {
	if o.ScopeId = v; o.ScopeId == nil {
		o.nullFields = append(o.nullFields, "ScopeId")
	}
	return o
}

func (o *EventSubscription) SetEventType(v *string) *EventSubscription {
	if o.EventType = v; o.EventType == nil {
		o.nullFields = append(o.nullFields, "EventType")
	}
	return o
}

//endregion

//endregion
