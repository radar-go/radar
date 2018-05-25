// Package register implements the register action for the web.
package register

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
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/web/actionsprovider"
	"github.com/radar-go/radar/web/actionsprovider/actions/action"
	"github.com/radar-go/radar/web/api"
	"github.com/radar-go/radar/web/controller/page"
)

// Register action.
type Register struct {
	action.Action
}

// New creates and returns a new login action object.
func New(cfg *config.Config) *Register {
	action := &Register{
		action.Action{
			Cfg:    cfg,
			Path:   "/register",
			Params: make(map[string]string),
		},
	}

	return action
}

// New creates and returns a new login action object.
func (r *Register) New(cfg *config.Config) actionsprovider.Action {
	return New(cfg)
}

// Run the register action.
func (r *Register) Run(ctx *fasthttp.RequestCtx) (actionsprovider.ActionResponse, error) {
	var p *page.Page
	var err error
	resp := action.NewResponse()

	if ctx.IsGet() {
		resp.SetPage(page.New("register", "Radar - Register", r.Cfg))
	} else if ctx.IsPost() {
		err := r.checkParams()
		if err != nil {
			p = page.New("register", "Radar - Register", r.Cfg)
			p.AddError("params", fmt.Sprint(err))
			resp.SetPage(p)

			return resp, nil
		}

		glog.Infof("Calling the API...")
		a := api.New(r.Cfg.APIHost, r.Cfg.APIPort)
		a.Connect()
		req, err := a.NewRequest("/account/register", "POST")
		if err != nil {
			glog.Errorf("Error creating a new request for the API: %s", err)
			return nil, fmt.Errorf("Error calling the API: %s", err)
		}

		req.Referer(ctx.Referer())
		req.AddParameter("username", r.Params["username"])
		req.AddParameter("name", r.Params["name"])
		req.AddParameter("email", r.Params["email"])
		req.AddParameter("password", r.Params["password"])
		apiResponse, err := req.Do()
		if err != nil {
			glog.Errorf("Error registering an account: %s", err)
			return nil, fmt.Errorf("Error calling the API: %s", err)
		}

		glog.Infof("API response: %s", apiResponse.Raw())
		if apiResponse.Code() != 200 {
			p = page.New("register", "Radar - Register", r.Cfg)
			p.AddError("register", apiResponse.Parsed()["error"].(string))
		} else {
			glog.Infof("API response: %s", apiResponse.Raw())
			glog.Infof("Referer: %s", ctx.Referer())
			p = page.New("register", "Radar - Register", r.Cfg)
		}

		resp.SetPage(p)
	}

	return resp, err
}

// checkParams checks the parameters passed to the register action.
func (r *Register) checkParams() error {
	var err error

	glog.Infof("Checking %d params", len(r.Params))
	if len(r.Params["username"]) == 0 {
		err = errors.New("The username can not be empty")
	} else if len(r.Params["email"]) == 0 {
		err = errors.New("The email can not be empty")
	} else if len(r.Params["password"]) == 0 {
		err = errors.New("The password can not be empty")
	} else if r.Params["password"] != r.Params["repeat-password"] {
		err = errors.New("The password and the password confirmation must match")
	}

	return err
}
