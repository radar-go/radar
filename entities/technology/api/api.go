// Package api defines the protocol for the technology entity.
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
	"github.com/radar-go/radar/entities/technology"
)

// Technology entity represents a technology used in one project, by one member
// or in one resource entity.
type Technology interface {
	Name() string
	Type() string
	Level() int

	SetName(newName string)
	SetType(newType string)
	SetLevel(newLevel int)

	Equals(tech interface{}) bool
}

// New returns a new Technology object.
func New(name, techType string, level int) Technology {
	tech := &technology.Technology{}

	tech.SetName(name)
	tech.SetType(techType)
	tech.SetLevel(level)

	return tech
}
