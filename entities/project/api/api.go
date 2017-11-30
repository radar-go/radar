// Package api defines the protocol for the project entity.
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
	member "github.com/radar-go/radar/entities/member/api"
	technology "github.com/radar-go/radar/entities/technology/api"
)

// Project entity represents a project done in the organization.
type Project interface {
	Name() string
	Members() []member.Member
	Technologies() []technology.Technology

	SetName(name string)
	AddMember(newMember member.Member)
	AddTechnology(newTechnology technology.Technology)
	DeleteMember(member member.Member) error
	DeleteTechnology(tech technology.Technology) error
}
