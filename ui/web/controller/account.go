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
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/ui/web/controller/page"
	"github.com/radar-go/radar/ui/web/templates"
)

// accountLogin handler.
func (c *Controller) accountLogin(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("text/html; charset=utf-8")

	writer := c.minify.Writer("text/html", ctx)
	defer writer.Close()
	p := page.New("login", "Radar - Login", c.cfg)

	templates.WritePageTemplate(writer, p.Get())
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