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

	"github.com/golang/glog"
	"github.com/valyala/fasthttp"

	"github.com/radar-go/radar/config"
	"github.com/radar-go/radar/ui/api/controller"
)

// API structure to manage the Radar API.
type API struct {
}

// New creates and returns a new API object.
func New() *API {
	return &API{}
}

// Start starts the Radar API.
func (a *API) Start() {
	cfg := config.New()
	c := controller.New()
	server := fasthttp.Server{
		Handler:           fasthttp.CompressHandler(c.Router.Handler),
		ReadBufferSize:    1024 * 64,
		WriteBufferSize:   1024 * 64,
		ReduceMemoryUsage: true,
	}

	glog.Infof("Starting api on port %d...", cfg.APIPort)
	err := server.ListenAndServe(fmt.Sprint(":", cfg.APIPort))
	if err != nil {
		glog.Exit(err)
	}
}
