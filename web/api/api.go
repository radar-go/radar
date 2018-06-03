// Package api is an interface to do requests to the API
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
	"strconv"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"
)

// API object to connect to the radar API.
type API struct {
	host   string
	port   int
	client *fasthttp.Client
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
func (a *API) NewRequest(path, method string) (*Request, error) {
	if a.client == nil {
		a.Connect()
	}

	req := &Request{
		api:        a,
		method:     "GET",
		path:       "/",
		parameters: make(map[string]interface{}),
		req:        fasthttp.AcquireRequest(),
	}

	req.Path(path)
	err := req.Method(method)

	return req, err
}

// SessionIsValid checks against the API if the session id is still valid and
// belongs to the user.
func (a *API) SessionIsValid(session, id []byte) bool {
	if a.client == nil {
		a.Connect()
	}

	req, err := a.NewRequest("/account/session", "POST")
	if err != nil {
		glog.Errorf("Error creating a new request for the API: %s", err)
		return false
	}

	fmtID, err := strconv.Atoi(string(id))
	if err != nil {
		glog.Errorf("Id %s in the wrong format: %s", id, err)
		return false
	}

	req.AddParameter("id", fmtID)
	req.AddParameter("session", session)
	apiResponse, err := req.Do()
	if err != nil {
		glog.Errorf("Login error from the API: %s", err)
		return false
	}

	glog.Infof("Api response %d %s", apiResponse.Code(), apiResponse.Raw())

	return apiResponse.Code() == 200
}
