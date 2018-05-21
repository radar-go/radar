// Package login implements the login action for the web.
package login

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
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/ui/web/actionsprovider"
	"github.com/radar-go/radar/ui/web/actionsprovider/actions/action"
	"github.com/radar-go/radar/ui/web/api"
	"github.com/radar-go/radar/ui/web/controller/page"
)

// Login action.
type Login struct {
	action.Action
}

// New creates and returns a new login action object.
func New(cfg *config.Config) *Login {
	action := &Login{
		action.Action{
			Cfg:    cfg,
			Path:   "/login",
			Params: make(map[string]string),
		},
	}

	return action
}

// New creates and returns a new login action object.
func (l *Login) New(cfg *config.Config) actionsprovider.Action {
	return New(cfg)
}

// Run runs the login action.
func (l *Login) Run(ctx *fasthttp.RequestCtx) (actionsprovider.ActionResponse, error) {
	var err error
	resp := action.NewResponse()

	if ctx.IsGet() {
		resp.SetPage(page.New("login", "Radar - Login", l.Cfg))
	} else if ctx.IsPost() {
		err := l.checkParams()
		if err != nil {
			p := page.New("login", "Radar - Register", l.Cfg)
			p.AddError("params", fmt.Sprint(err))
			resp.SetPage(p)

			return resp, nil
		}

		a := api.New(l.Cfg.APIHost, l.Cfg.APIPort)
		a.Connect()
		req, err := a.NewRequest("/account/login", "POST")
		if err != nil {
			glog.Errorf("Error creating a new request for the API: %s", err)
			return nil, fmt.Errorf("Error calling the API: %s", err)
		}

		req.Referer(ctx.Referer())
		req.AddParameter("login", l.Params["username"])
		req.AddParameter("password", l.Params["password"])
		apiResponse, err := req.Do()
		if err != nil {
			glog.Errorf("Login error from the API: %s", err)
			return nil, fmt.Errorf("Error calling the API: %s", err)
		}

		if apiResponse.Code() != 200 {
			title := fmt.Sprintf("Radar - Login")
			p := page.New("login", title, l.Cfg)
			p.AddError("login", apiResponse.Parsed()["error"].(string))
			resp.SetPage(p)
		} else {
			parsed := apiResponse.Parsed()
			actionsprovider.SetCookie(ctx, "id",
				fmt.Sprintf("%f", parsed["id"].(float64)), 24*time.Hour)
			actionsprovider.SetCookie(ctx, "username", parsed["username"].(string),
				24*time.Hour)
			actionsprovider.SetCookie(ctx, "name", parsed["name"].(string),
				24*time.Hour)
			actionsprovider.SetCookie(ctx, "email", parsed["email"].(string),
				24*time.Hour)
			actionsprovider.SetCookie(ctx, "session", parsed["token"].(string),
				24*time.Hour)

			glog.Infof("Referer: %s", ctx.Referer())
			if bytes.Contains(ctx.Referer(), []byte("/login")) {
				redirection := bytes.Replace(ctx.Referer(), []byte("/login"),
					[]byte("/account"), 1)
				glog.Infof("Redirecting to %s", redirection)
				resp.SetRedirectionURL(string(redirection[:]))

				return resp, err
			}

			glog.Infof("Redirecting to the referer page: %s", ctx.Referer())
			redirection := ctx.Referer()
			glog.Infof("Redirecting to %s", redirection)
			resp.SetRedirectionURL(string(redirection[:]))

			return resp, err
		}
	}

	return resp, err
}

// checkParams checks the parameters passed to the login action.
func (l *Login) checkParams() error {
	var err error

	if len(l.Params["username"]) == 0 {
		err = errors.New("The username can not be empty")
	} else if len(l.Params["password"]) == 0 {
		err = errors.New("The password can not be empty")
	}

	return err
}
