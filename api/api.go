// Package api contains the api to access the radar test cases.
package api

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
	"fmt"
	"net"
	"time"

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/api/controller"
	"github.com/radar-go/radar/config"
)

// API structure to manage the Radar API.
type API struct {
	listener net.Listener
}

// New creates and returns a new API object.
func New() *API {
	return &API{}
}

// Start starts the Radar API.
func (a *API) Start() error {
	var err error
	cfg := config.New()
	c := controller.New()

	server := fasthttp.Server{
		Handler:           fasthttp.CompressHandler(c.Router.Handler),
		ReadBufferSize:    1024 * 64,
		WriteBufferSize:   1024 * 64,
		ReduceMemoryUsage: true,
	}

	a.listener, err = net.Listen("tcp4", fmt.Sprint(":", cfg.APIPort))
	if err != nil {
		return err
	}

	go func() {
		glog.Infof("Starting api on port %d...", cfg.APIPort)
		err := server.Serve(a.listener)
		if err != nil {
			glog.Exit("Error starting the API: %s", err)
		}
	}()

	return err
}

// Stop the Radar API.
func (a *API) Stop() error {
	var err error

	if a.listener != nil {
		err = a.listener.Close()
		time.Sleep(time.Second)
		a.listener = nil
	}

	return err
}

// Reload the Radar API.
func (a *API) Reload() error {
	err := a.Stop()
	if err != nil {
		glog.Errorf("Error stoping the api server: %s", err)
		return err
	}

	err = a.Start()
	if err != nil {
		glog.Exitf("Error starting the api server: %s", err)
	}

	return err
}
