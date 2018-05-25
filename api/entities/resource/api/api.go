// Package api defines the protocol for the resource entity.
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
	"github.com/radar-go/radar/api/entities/resource"
	technology "github.com/radar-go/radar/api/entities/technology/api"
)

// Resource entity represents a resource to learn one or several technologies.
type Resource interface {
	Name() string
	URL() string
	Technologies() []technology.Technology
	Rate() float64

	SetName(name string)
	SetURL(url string)

	AddRate(newRate float64)
	AddTechnology(newTechnology technology.Technology)
	DeleteRate(rate float64) error
	DeleteTechnology(tech technology.Technology) error
}

// New creates a new Resource object.
func New(name, url string) Resource {
	res := &resource.Resource{}
	res.SetName(name)
	res.SetURL(url)

	return res
}
