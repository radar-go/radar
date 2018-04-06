package web

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
	"net"
	"time"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/svg"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/ui/web/controller"
)

// Web structure to manage the Radar web interface.
type Web struct {
	listener net.Listener
	mins     *minify.M
}

// New creates and returns a new Web object.
func New() *Web {
	w := &Web{
		mins: minify.New(),
	}

	w.mins.AddFunc("text/css", css.Minify)
	w.mins.AddFunc("text/html", html.Minify)
	w.mins.AddFunc("text/javascript", js.Minify)
	w.mins.AddFunc("image/svg+xml", svg.Minify)

	return w
}

// Start the Radar web interface.
func (w *Web) Start() error {
	var err error
	cfg := config.New()
	c := controller.New(cfg, w.mins)
	server := fasthttp.Server{
		Handler:           fasthttp.CompressHandler(c.Router.Handler),
		ReadBufferSize:    1024 * 64,
		WriteBufferSize:   1024 * 64,
		ReduceMemoryUsage: true,
	}

	w.listener, err = net.Listen("tcp4", fmt.Sprint(":", cfg.WebPort))
	if err != nil {
		return err
	}

	go func() {
		glog.Infof("Starting web on port %d...", cfg.WebPort)
		err := server.Serve(w.listener)
		if err != nil {
			glog.Exitf("Error starting the server: %s", err)
		}
	}()

	return err
}

// Stop the Radar web interface.
func (w *Web) Stop() error {
	var err error

	if w.listener != nil {
		err = w.listener.Close()
		time.Sleep(time.Second)
		w.listener = nil
	}

	return err
}

// Reload the Radar web interface.
func (w *Web) Reload() error {
	err := w.Stop()
	if err != nil {
		glog.Errorf("Error stoping the web server: %s", err)
		return err
	}

	err = w.Start()
	if err != nil {
		glog.Exitf("Error starting the web server: %s", err)
	}

	return err
}
