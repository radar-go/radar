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

// register defines all the router paths the API implements.
func (c *Controller) register() {
}
