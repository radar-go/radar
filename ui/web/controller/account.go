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
	"fmt"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/ui/web/controller/page"
	"github.com/radar-go/radar/ui/web/templates"
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

		title := fmt.Sprintf("Radar - Login - %s", resp.Raw())
		p = page.New("login", title, c.cfg)
		if resp.Code() != 200 {
			p.AddError("login", resp.Parsed()["error"].(string))
		}
	}

	c.response(ctx, p)
}

// accountRegister handler.
func (c *Controller) accountRegister(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("text/html; charset=utf-8")

	writer := c.minify.Writer("text/html", ctx)
	defer writer.Close()
	p := page.New("register", "Radar - Register", c.cfg)

	templates.WritePageTemplate(writer, p.Get())
}
