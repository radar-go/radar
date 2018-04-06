// Package main implements the radar command.
package main

/* Copyright (C) 2017-2018 Radar team (see AUTHORS)

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
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"

	"github.com/radar-go/radar/ui/api"
	"github.com/radar-go/radar/ui/web"
)

func main() {
	/* Parse the arguments. */
	flag.Parse()

	/* Start the radar API. */
	a := api.New()
	err := a.Start()
	if err != nil {
		glog.Exit(err)
	}

	/* Start the radar web interface. */
	w := web.New()
	err = w.Start()
	if err != nil {
		glog.Exit(err)
	}

	exit := make(chan os.Signal, 1)
	reload := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(reload, syscall.SIGUSR1)

	for {
		select {
		case <-exit:
			glog.Info("Stoping servers...")
			err := a.Stop()
			if err != nil {
				glog.Exitf("Error stoping the api server: %s", err)
			}

			err = w.Stop()
			if err != nil {
				glog.Exitf("Error stoping the web server: %s", err)
			}

			os.Exit(0)
		case <-reload:
			glog.Info("Reloading servers...")
			err := a.Reload()
			if err != nil {
				glog.Exitf("Error reloading the api server: %s", err)
			}

			err = w.Reload()
			if err != nil {
				glog.Exitf("Error reloading the web server: %s", err)
			}

			glog.Info("Servers reloaded")
		}
	}

}
