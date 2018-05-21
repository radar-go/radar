package controller

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
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/ui/web/controller/page"
)

// accountLogin handler.
func (c *Controller) accountLogin(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	var p *page.Page

	if ctx.IsGet() {
		p = page.New("login", "Radar - Login", c.cfg)
	} else if ctx.IsPost() {
		args := ctx.PostArgs()
		err := c.checkParams(ctx, "email", "password")
		if err != nil {
			p.AddError("params", fmt.Sprint(err))
			c.response(ctx, p)

			return
		}

		req, err := c.api.NewRequest("/account/login", "POST")
		if err != nil {
			glog.Errorf("Error setting the method: %s", err)
			c.panic(ctx, "Error calling the API")
		}

		req.Referer(ctx.Referer())
		req.AddParameter("login", args.Peek("email"))
		req.AddParameter("password", args.Peek("password"))
		resp, err := req.Do()
		if err != nil {
			glog.Errorf("Error login: %s", err)
			c.panic(ctx, "Error calling the API")
		}

		if resp.Code() != 200 {
			title := fmt.Sprintf("Radar - Login")
			p = page.New("login", title, c.cfg)
			p.AddError("login", resp.Parsed()["error"].(string))
		} else {
			c.setCookie(ctx, "id", resp.Parsed()["id"].(string), 24*time.Hour)
			c.setCookie(ctx, "username", resp.Parsed()["username"].(string), 24*time.Hour)
			c.setCookie(ctx, "name", resp.Parsed()["name"].(string), 24*time.Hour)
			c.setCookie(ctx, "email", resp.Parsed()["email"].(string), 24*time.Hour)
			c.setCookie(ctx, "session", resp.Parsed()["token"].(string), 24*time.Hour)

			glog.Infof("Referer: %s", ctx.Referer())
			if bytes.Contains(ctx.Referer(), []byte("/login")) {
				redirect := bytes.Replace(ctx.Referer(), []byte("/login"), []byte("/account"), 1)
				ctx.Redirect(fmt.Sprintf("%s", redirect), 301)
			} else {
				ctx.Redirect(fmt.Sprintf("%s", ctx.Referer()), 301)
			}
		}
	}

	c.response(ctx, p)
}

// accountRegister handler.
func (c *Controller) accountRegister(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	var p *page.Page

	if ctx.IsGet() {
		p = page.New("register", "Radar - Register", c.cfg)
	} else if ctx.IsPost() {
		p = page.New("register", "Radar - Register post", c.cfg)
	}

	c.response(ctx, p)
}
