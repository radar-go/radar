// Package main implements the radar command.
package main

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
	"flag"

	"github.com/golang/glog"

	"github.com/radar-go/radar/ui/api"
)

func main() {
	/* Parse the arguments. */
	flag.Parse()

	/* Starts the radar API. */
	a := api.New()
	err := a.Start()
	if err != nil {
		glog.Exit(err)
	}
}
