// Package actionsprovider provides the actions handled by the web controller.
package actionsprovider

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
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/web/controller/page"
)

// ActionResponse defines the behaviour of a response from a web action.
type ActionResponse interface {
	Page() *page.Page
	IsRedirect() bool
	RedirectionURL() string
}

// Action defines the behaviour of a web action.
type Action interface {
	AddParam([]byte, []byte)
	GetPath() string
	New(*config.Config) Action
	Run(*fasthttp.RequestCtx) (ActionResponse, error)
}

// WebActions struct to register and call the different web actions.
type WebActions struct {
	action map[string]Action
	method map[string][]string
}

var actions = &WebActions{
	action: make(map[string]Action),
	method: make(map[string][]string),
}

// Register a new web action.
func Register(action Action, method string) {
	if _, ok := actions.action[action.GetPath()]; ok {
		glog.Errorf("Action %s already register", action.GetPath())
	}

	actions.action[action.GetPath()] = action

	if _, ok := actions.method[method]; !ok {
		actions.method[method] = make([]string, 0)
	}

	actions.method[method] = append(actions.method[method], action.GetPath())
}

// GetAction returns a registered action by its name.
func GetAction(name string, cfg *config.Config) (Action, error) {
	action, ok := actions.action[name]
	if !ok {
		return nil, fmt.Errorf("Action %s is not registered", name)
	}

	return action.New(cfg), nil
}

// GetPaths return the list of paths based on a method (GET, POST, PUT, ...).
func GetPaths(method string) ([]string, error) {
	paths, ok := actions.method[method]
	if !ok {
		return nil, fmt.Errorf("No actions registered for the method %s", method)
	}

	return paths, nil
}

// ActionsList returns the list of all actions registered.
func ActionsList() []string {
	list := make([]string, 0, len(actions.action))
	for k := range actions.action {
		list = append(list, k)
	}

	return list
}

// SetCookie sets a new cookie in the client.
func SetCookie(ctx *fasthttp.RequestCtx, name, value string, t time.Duration) {
	glog.Infof("Getting cookie")
	cookie := fasthttp.AcquireCookie()
	cookie.SetKey(name)
	cookie.SetValue(value)
	// cookie.SetPath("/")
	// cookie.SetDomain("mydomain.com")
	cookie.SetExpire(time.Now().Add(t))
	// cookie.SetSecure(true)
	ctx.Response.Header.Cookie(cookie)
	glog.Infof("Cookie %s setted to %s", name, value)
}
