package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
)

type Request struct {
	Obj    interface{}
	Entity interface{}
	Params url.Values
	url    *url.URL
	method string
	body   io.Reader
	header http.Header
}

// toHTTP converts the request to an HTTP request.
func (r *Request) toHTTP(ctx context.Context, cfg *controlmonkey.Config, shouldWrapWithEntity bool) (*http.Request, error) {
	// Set the user credentials.
	creds, err := cfg.Credentials.Get()
	if err != nil {
		return nil, err
	}
	if creds.Token != "" {
		r.header.Set("Authorization", "Bearer "+creds.Token)
	}

	// Encode the query parameters.
	r.url.RawQuery = r.Params.Encode()

	// Check if we should encode the body.
	if r.body == nil && r.Obj != nil {
		var body io.Reader
		var err error

		if shouldWrapWithEntity {
			entity := map[string]interface{}{
				"entity": r.Obj,
			}
			body, err = EncodeBody(entity)
		} else {
			body, err = EncodeBody(r.Obj)
		}

		if err != nil {
			return nil, err
		} else {
			r.body = body
		}
	}

	// Create the HTTP request.
	req, err := http.NewRequest(r.method, r.url.RequestURI(), r.body)
	if err != nil {
		return nil, err
	}

	// Set request base URL.
	req.URL.Host = cfg.BaseURL.Host
	req.URL.Scheme = cfg.BaseURL.Scheme

	// Set request headers.
	req.Host = cfg.BaseURL.Host
	req.Header = r.header
	req.Header.Set("Content-Type", cfg.ContentType)
	req.Header.Add("Accept", cfg.ContentType)
	req.Header.Add("User-Agent", cfg.UserAgent)

	return req.WithContext(ctx), nil
}

// EncodeBody is used to encode a request body
func EncodeBody(obj interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(obj); err != nil {
		return nil, err
	}
	return buf, nil
}
