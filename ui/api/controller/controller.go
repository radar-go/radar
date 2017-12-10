// Package controller implements the Radar API controller.
package controller

/* Copyright (C) 2017 Radar team (see AUTHORS)

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
	"github.com/buaazp/fasthttprouter"
	"github.com/golang/glog"
	"github.com/valyala/fasthttp"
)

// Controller struct to manager the Radar API Controller.
type Controller struct {
	Router *fasthttprouter.Router
}

// New creates and return a new Controller object.
func New() *Controller {
	c := &Controller{
		Router: fasthttprouter.New(),
	}
	c.register()

	return c
}

// logPaths logs the requested path to the info log.
func logPath(path []byte) {
	glog.Infof("Request path: %s", path)
}

// register defines all the router paths the API implements.
func (c *Controller) register() {
	c.Router.HandleMethodNotAllowed = true
	c.Router.NotFound = c.notFound
	c.Router.MethodNotAllowed = c.methodNotAllowed
	c.Router.PanicHandler = c.panic

	c.Router.GET("/healthcheck", c.healthcheck)
}

// panic handles when the server have a fatal error.
func (c *Controller) panic(ctx *fasthttp.RequestCtx, from interface{}) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}

// methodNotAllowed handles the response when a method call is not allowed from
// the client.
func (c *Controller) methodNotAllowed(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
}

// notFound handles the response when a path have not been found.
func (c *Controller) notFound(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusNotFound)
}

// healthcheck handler.
func (c *Controller) healthcheck(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusOK)
}
