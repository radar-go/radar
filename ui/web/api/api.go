// Package api is an interface to do requests to the API
package api

import (
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
	parameters map[string]string
	path       string
	req        *fasthttp.Request
}

// Response from the radar API.
type Response struct {
	code   int
	parsed map[string]interface{}
	raw    string
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
		parameters: make(map[string]string),
		req:        fasthttp.AcquireRequest(),
	}
}

// Path sets the request path to the Radar API.
func (r *Request) Path(p string) {
	// XXX: Check the path is correct.
	r.path = p
}

// Method sets the method for the request to the Radar API.
func (r *Request) Method(m string) {
	// XXX: Check the method is correct.
	r.method = m
}

// AddParameter adds a new parameter to the request to the Radar API.
func (r *Request) AddParameter(key, value string) {
	r.parameters[key] = value
}

// Do the request to the Radar API.
func (r *Request) Do() (*Response, error) {
	var err error

	resp := &Response{
		parsed: make(map[string]interface{}),
		resp:   fasthttp.AcquireResponse(),
	}

	// r.SetRequestURI(url)
	err = r.api.client.Do(r.req, resp.resp)
	// bodyBytes := resp.Body()

	return resp, err
}

// Code returns the response code from the request to the Radar API.
func (r *Response) Code() int {
	return r.code
}

// Raw returns the raw response from the Radar API.
func (r *Response) Raw() string {
	return r.raw
}

// Parsed returns the parsed response from the Radar API.
func (r *Response) Parsed() map[string]interface{} {
	return r.parsed
}
