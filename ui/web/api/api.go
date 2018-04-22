// Package api is an interface to do requests to the API
package api

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// API object to connect to the radar API.
type API struct {
	host   string
	port   int
	client *fasthttp.Client
}

// Request to the Radar API.
type Request struct {
	api        *API
	method     string
	parameters map[string]interface{}
	path       string
	referer    []byte
	req        *fasthttp.Request
}

// Response from the radar API.
type Response struct {
	code   int
	parsed map[string]interface{}
	raw    []byte
	resp   *fasthttp.Response
}

// New returns a new API struct.
func New(host string, port int) *API {
	return &API{
		host: host,
		port: port,
	}
}

// Connect to the Radar API.
func (a *API) Connect() {
	a.client = &fasthttp.Client{}
}

// NewRequest creates a new Request to the Radar API.
func (a *API) NewRequest() *Request {
	return &Request{
		api:        a,
		method:     "GET",
		path:       "/",
		parameters: make(map[string]interface{}),
		req:        fasthttp.AcquireRequest(),
	}
}

func (r *Request) Referer(referer []byte) {
	// XXX: Check the referer is well formed.
	r.referer = referer
}

// Path sets the request path to the Radar API.
func (r *Request) Path(p string) {
	// XXX: Check the path is well formed.
	r.path = p
}

// Method sets the method for the request to the Radar API.
func (r *Request) Method(m string) error {
	if m != "GET" && m != "POST" && m != "PUT" {
		return fmt.Errorf("Unknown HTTP method: %s", m)
	}

	r.method = m

	return nil
}

// AddParameter adds a new parameter to the request to the Radar API.
func (r *Request) AddParameter(key string, value interface{}) {
	r.parameters[key] = value
}

// Do the request to the Radar API.
func (r *Request) Do() (*Response, error) {
	var err error

	body, err := json.Marshal(r.parameters)
	if err != nil {
		return nil, err
	}

	resp := &Response{
		parsed: make(map[string]interface{}),
		resp:   fasthttp.AcquireResponse(),
	}

	r.req.Reset()
	// XXX: Set web host
	r.req.SetHost(r.api.host)
	r.req.SetRequestURI(fmt.Sprintf("http://%s:%d/%s", r.api.host, r.api.port, r.path))
	r.req.Header.SetContentType("application/json; charset=utf-8")
	r.req.Header.SetMethod(r.method)
	r.req.Header.SetRefererBytes(r.referer)
	r.req.SetBody(body)
	err = r.api.client.Do(r.req, resp.resp)

	// Get the response
	resp.raw = resp.resp.Body()
	resp.code = resp.resp.StatusCode()
	json.Unmarshal(resp.raw, &resp.parsed)

	return resp, err
}

// Code returns the response code from the request to the Radar API.
func (r *Response) Code() int {
	return r.code
}

// Raw returns the raw response from the Radar API.
func (r *Response) Raw() []byte {
	return r.raw
}

// Parsed returns the parsed response from the Radar API.
func (r *Response) Parsed() map[string]interface{} {
	return r.parsed
}
