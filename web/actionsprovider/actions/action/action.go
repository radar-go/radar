// Package action implements a generic web action.
package action

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

	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/web/actionsprovider"
	"github.com/radar-go/radar/web/controller/page"
)

// Response represents a generic user case response.
type Response struct {
	page        *page.Page
	redirect    bool
	redirectURL string
}

// NewResponse creates a new response object.
func NewResponse() *Response {
	return &Response{}
}

// Page returns the resulting page for the action response.
func (r *Response) Page() *page.Page {
	return r.page
}

// IsRedirect returns true if the response of the action is a redirection or false
// otherwise.
func (r *Response) IsRedirect() bool {
	return r.redirect
}

// RedirectionURL returns the url to redirect as an action response.
func (r *Response) RedirectionURL() string {
	return r.redirectURL
}

// SetPage sets the response page.
func (r *Response) SetPage(newPage *page.Page) {
	r.page = newPage
}

// SetRedirectionURL sets the url for the redirection. If the url is not empty,
// it sets IsRedirect to true, otherwise it sets it to false.
func (r *Response) SetRedirectionURL(url string) {
	r.redirect = (len(url) > 0)
	r.redirectURL = url
}

// Action represents a generic use action.
type Action struct {
	Cfg    *config.Config
	Path   string
	Params map[string]string
}

// New returns a new Action object.
func (a *Action) New(cfg *config.Config) actionsprovider.Action {
	return &Action{
		Cfg:    cfg,
		Path:   "/",
		Params: make(map[string]string),
	}
}

// GetPath returns the action path.
func (a *Action) GetPath() string {
	return a.Path
}

// AddParam adds a new ad param to the web action.
func (a *Action) AddParam(newKey, newValue []byte) {
	key := string(newKey[:])
	value := string(newValue[:])

	a.Params[key] = value
}

// Run executes an action based on the method.
func (a *Action) Run(ctx *fasthttp.RequestCtx) (actionsprovider.ActionResponse, error) {
	return nil, fmt.Errorf("Function Run not implemented")
}

// ErrorURL returns the error (404) url.
func (a *Action) ErrorURL() string {
	return fmt.Sprintf("http://%s:%d/404", a.Cfg.WebHost, a.Cfg.WebPort)
}
