package account

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
	"bytes"
	"errors"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/web/actionsprovider"
	"github.com/radar-go/radar/web/actionsprovider/actions/action"
	"github.com/radar-go/radar/web/api"
	"github.com/radar-go/radar/web/controller/page"
)

// Account action.
type Account struct {
	action.Action
}

// New creates and returns a new account action object.
func New(cfg *config.Config) *Account {
	action := &Account{
		action.Action{
			Cfg:    cfg,
			Path:   "/account",
			Params: make(map[string]string),
		},
	}

	return action
}

// New creates and returns a new account action object.
func (a *Account) New(cfg *config.Config) actionsprovider.Action {
	return New(cfg)
}

// Run runs the account action.
func (a *Account) Run(ctx *fasthttp.RequestCtx) (actionsprovider.ActionResponse, error) {
	var err error
	resp := action.NewResponse()

	if !ctx.IsGet() {
		resp.SetRedirectionURL(a.ErrorURL())
		err = errors.New("Only GET method allowed")
	} else if ctx.Request.Header.Cookie("session") == nil {
		glog.Info("User not logged in, redirecting to the login page")
		redirection := bytes.Replace(ctx.Request.RequestURI(), []byte("/account"),
			[]byte("/login"), 1)
		resp.SetRedirectionURL(string(redirection[:]))
	} else {
		resp.SetPage(page.New("account", "Radar - Account", a.Cfg))

		apiClient := api.New(a.Cfg.APIHost, a.Cfg.APIPort)
		session := ctx.Request.Header.Cookie("session")
		id := ctx.Request.Header.Cookie("id")
		if !apiClient.SessionIsValid(session, id) {
			redirection := bytes.Replace(ctx.Request.RequestURI(), []byte("/account"),
				[]byte("/login"), 1)
			resp.SetRedirectionURL(string(redirection[:]))
		}
	}

	return resp, err
}
