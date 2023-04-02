package client

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
)

// Client provides a client to the API.
type Client struct {
	config *controlmonkey.Config
}

// New returns a new client.
func New(cfg *controlmonkey.Config) *Client {
	if cfg == nil {
		cfg = controlmonkey.DefaultConfig()
	}
	return &Client{cfg}
}

// NewRequest is used to create a new request.
func NewRequest(method, path string) *Request {
	return &Request{
		method: method,
		url: &url.URL{
			Path: path,
		},
		header: make(http.Header),
		Params: make(url.Values),
	}
}

func (c *Client) Do(ctx context.Context, r *Request) (*http.Response, error) {
	return c.Do2(ctx, r, true)
}

// Do2 runs a request with our client.
func (c *Client) Do2(ctx context.Context, r *Request, shouldWrapWithEntity bool) (*http.Response, error) {
	req, err := r.toHTTP(ctx, c.config, shouldWrapWithEntity)
	if err != nil {
		return nil, err
	}
	c.logRequest(req)
	resp, err := c.config.HTTPClient.Do(req)
	c.logResponse(resp)
	return resp, err
}

func (c *Client) logf(format string, args ...interface{}) {
	if c.config.Logger != nil {
		c.config.Logger.Printf(format, args...)
	}
}

const logReqMsg = `CONTROL MONKEY: Request "%s %s" details:
---[ REQUEST ]---------------------------------------
%s
-----------------------------------------------------`

func (c *Client) logRequest(req *http.Request) {
	if c.config.Logger != nil && req != nil {
		out, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			c.logf(logReqMsg, req.Method, req.URL, string(out))
		}
	}
}

const logRespMsg = `CONTROL MONKEY: Response "%s %s" details:
---[ RESPONSE ]----------------------------------------
%s
-------------------------------------------------------`

func (c *Client) logResponse(resp *http.Response) {
	if c.config.Logger != nil && resp != nil {
		out, err := httputil.DumpResponse(resp, true)
		if err == nil {
			c.logf(logRespMsg, resp.Request.Method, resp.Request.URL, string(out))
		}
	}
}
