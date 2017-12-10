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
)

type API struct {
}

func New() *API {
	return &API{}
}

func (a *API) Start() {
	cfg := config.New()
	c := newController()
	server := fasthttp.Server{
		Handler:           fasthttp.CompressHandler(c.router.Handler),
		ReadBufferSize:    1024 * 64,
		WriteBufferSize:   1024 * 64,
		ReduceMemoryUsage: true,
	}

	glog.Infof("Starting server on port %d...", cfg.APIPort)
	err := server.ListenAndServe(fmt.Sprint(":", cfg.APIPort))
	if err != nil {
		glog.Exit(err)
	}
}
