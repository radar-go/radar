package api

/* Copyright (C) 2018 Radar team (see AUTHORS)

   This file is part of radar.

   radar is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   radar is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with radar. If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"
)

// Request to the Radar API.
type Request struct {
	api        *API
	method     string
	parameters map[string]interface{}
	path       string
	referer    []byte
	req        *fasthttp.Request
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
	resp := &Response{
		parsed: make(map[string]interface{}),
		resp:   fasthttp.AcquireResponse(),
	}

	// Call the API.
	r.prepareRequest()
	glog.Infof("Calling %s with the method %s", r.req.URI(), r.method)
	err := r.api.client.Do(r.req, resp.resp)
	if err != nil {
		glog.Errorf("Error calling the API: %s", err)
		return resp, err
	}

	// Get the response
	resp.raw = resp.resp.Body()
	resp.code = resp.resp.StatusCode()
	err = json.Unmarshal(resp.raw, &resp.parsed)
	if err != nil {
		glog.Errorf("Error unmarshaling the response: %s", err)
	}

	return resp, err
}

func (r *Request) prepareRequest() error {
	body, err := r.marshalParameters()
	if err != nil {
		return err
	}

	r.req.Reset()
	// XXX: Set web host
	//r.req.SetHost(r.api.host)
	uri := fmt.Sprintf("http://%s:%d%s", r.api.host, r.api.port, r.path)
	r.req.SetRequestURI(uri)
	r.req.Header.SetContentType("application/json; charset=utf-8")
	r.req.Header.SetMethod(r.method)
	r.req.Header.SetRefererBytes(r.referer)
	r.req.SetBody(body)

	return nil
}

// marshalParameters marshal the parameters into a json []byte.
func (r *Request) marshalParameters() ([]byte, error) {
	var err error

	marshaled := []byte("{")

	for name, param := range r.parameters {
		glog.Infof("Parameter %s with type %T", name, param)
		var content []byte
		element := []byte(`"` + name + `": `)

		switch param.(type) {
		case string:
			content = []byte(fmt.Sprintf(`"%s"`, param.(string)))
		case []byte:
			content = []byte(fmt.Sprintf(`"%s"`, param.([]byte)))
		case int:
			content = []byte(fmt.Sprintf("%d", param.(int)))
		case float32:
			content = []byte(fmt.Sprintf("%f", param.(float32)))
		case float64:
			content = []byte(fmt.Sprintf("%f", param.(float64)))
		case bool:
			content = []byte(fmt.Sprintf("%t", param.(bool)))
		default:
			err = fmt.Errorf("Unknown param type for %s", name)
			return marshaled, err
		}

		if len(marshaled) > 1 {
			marshaled = append(marshaled, []byte(", ")...)
		}

		element = append(element, content...)
		marshaled = append(marshaled, element...)
	}

	marshaled = append(marshaled, []byte("}")...)

	return marshaled, err
}
