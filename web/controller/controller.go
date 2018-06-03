// Package controller implements the Radar Web ui controller.
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
	"github.com/buaazp/fasthttprouter"
	"github.com/golang/glog"
	"github.com/tdewolff/minify"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/web/actionsprovider"
	_ "github.com/radar-go/radar/web/actionsprovider/actions"
	"github.com/radar-go/radar/web/controller/page"
	"github.com/radar-go/radar/web/templates"
)

// Controller struct to manager the Radar web ui Controller.
type Controller struct {
	Router        *fasthttprouter.Router
	minify        *minify.M
	staticHandler fasthttp.RequestHandler
	cfg           *config.Config
}

// New creates and return a new Controller object.
func New(cfg *config.Config, m *minify.M) *Controller {
	fs := &fasthttp.FS{
		Root:               cfg.StaticDir,
		GenerateIndexPages: false,
		Compress:           true,
		AcceptByteRange:    false,
	}

	c := &Controller{
		Router:        fasthttprouter.New(),
		minify:        m,
		staticHandler: fs.NewRequestHandler(),
		cfg:           cfg,
	}
	c.register()

	return c
}

// logPaths logs the requested path to the info log.
func logPath(path []byte) {
	glog.Infof("Request path: %s", path)
}

// register defines all the router paths the Web implements.
func (c *Controller) register() {
	c.Router.HandleMethodNotAllowed = true
	c.Router.NotFound = c.static
	c.Router.MethodNotAllowed = c.methodNotAllowed
	c.Router.PanicHandler = c.panic

	c.Router.GET("/healthcheck", c.healthcheck)
	c.Router.GET("/", c.home)

	paths, err := actionsprovider.Paths("GET")
	if err == nil {
		for _, path := range paths {
			c.Router.GET(path, c.webHandler)
		}
	}

	paths, err = actionsprovider.Paths("POST")
	if err == nil {
		for _, path := range paths {
			c.Router.POST(path, c.webHandler)
		}
	}
}

// panic handles when the server have a fatal error.
func (c *Controller) panic(ctx *fasthttp.RequestCtx, from interface{}) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
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

// static serves the static pages (css, js, imgs, ...) of the site.
func (c *Controller) static(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	c.staticHandler(ctx)
	resp := &ctx.Response
	if resp.StatusCode() == fasthttp.StatusNotFound ||
		resp.StatusCode() == fasthttp.StatusForbidden {
		c.notFound(ctx)
	}
}

// healthcheck handler.
func (c *Controller) healthcheck(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetBodyString(`{"status": "ok"}`)
}

// home handler.
func (c *Controller) home(ctx *fasthttp.RequestCtx) {
	logPath(ctx.Path())
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("text/html; charset=utf-8")

	writer := c.minify.Writer("text/html", ctx)
	defer writer.Close()
	p := page.New("home", "Radar", c.cfg)

	templates.WritePageTemplate(writer, p.Get())
}

func (c *Controller) webHandler(ctx *fasthttp.RequestCtx) {
	var args *fasthttp.Args
	logPath(ctx.Path())

	action, err := actionsprovider.GetAction(string(ctx.Path()[:]), c.cfg)
	if err != nil {
		glog.Error("Unknown action for path %s", ctx.Path())
		return
	}

	if ctx.IsGet() {
		args = ctx.QueryArgs()
	} else if ctx.IsPost() || ctx.IsPut() {
		args = ctx.PostArgs()
	}

	args.VisitAll(action.AddParam)
	resp, err := action.Run(ctx)
	if err != nil {
		glog.Errorf("Error running the action: %s", err)
		c.panic(ctx, err)
	}

	glog.Infof("Redirecting enabled? %b", resp.IsRedirect())
	if resp.IsRedirect() {
		glog.Infof("Redirecting to %s", resp.RedirectionURL())
		ctx.Redirect(resp.RedirectionURL(), 301)
		return
	}

	c.response(ctx, resp.Page())
}

// response to the client.
func (c *Controller) response(ctx *fasthttp.RequestCtx, p *page.Page) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("text/html; charset=utf-8")
	writer := c.minify.Writer("text/html", ctx)
	defer writer.Close()
	templates.WritePageTemplate(writer, p.Get())
}
