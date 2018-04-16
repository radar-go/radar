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
	"crypto/sha512"
	"encoding/base64"
	"io/ioutil"

	"github.com/buaazp/fasthttprouter"
	"github.com/golang/glog"
	"github.com/tdewolff/minify"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/ui/web/templates"
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

// register defines all the router paths the API implements.
func (c *Controller) register() {
	c.Router.HandleMethodNotAllowed = true
	c.Router.NotFound = c.static
	c.Router.MethodNotAllowed = c.methodNotAllowed
	c.Router.PanicHandler = c.panic

	c.Router.GET("/healthcheck", c.healthcheck)
	c.Router.GET("/", c.home)
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

	p := &templates.BasePage{}
	p.TemplateName = "home"
	p.TitleStr = "Radar"
	p.Copyright = "2017-2018 Radar authors"

	p.CSSArr = append(p.CSSArr, c.getSha386Sum("bootstrap.min.css", "css"))
	p.CSSArr = append(p.CSSArr, c.getSha386Sum("radar.css", "css"))

	p.JavascriptArr = append(p.JavascriptArr,
		c.getSha386Sum("jquery-3.2.1.slim.min.js", "js"))
	p.JavascriptArr = append(p.JavascriptArr,
		c.getSha386Sum("popper.min.js", "js"))
	p.JavascriptArr = append(p.JavascriptArr,
		c.getSha386Sum("bootstrap.min.js", "js"))

	templates.WritePageTemplate(writer, p)
}

func (c *Controller) getSha386Sum(file string, fileType string) templates.File {
	b, err := ioutil.ReadFile(c.cfg.StaticDir + "/" + fileType + "/" + file)
	if err != nil {
		glog.Errorf("Error reading the file %s: %s", file, err)
		return templates.File{
			Name: file,
		}
	}

	sum := sha512.Sum384(b)

	return templates.File{
		Name: file,
		Sum:  base64.StdEncoding.EncodeToString(sum[:]),
	}
}
