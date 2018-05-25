// Package api defines the protocol for the member entity.
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
	"github.com/radar-go/radar/api/entities/member"
	role "github.com/radar-go/radar/api/entities/role/api"
	technology "github.com/radar-go/radar/api/entities/technology/api"
)

// Member entity represents a member of the organization (employee, partner,
// associate, ...) and his relation  with the rest of entities.
type Member interface {
	Name() string
	Roles() []role.Role
	CurrentRole() role.Role
	Technologies() []technology.Technology

	SetName(string)
	Equals(interface{}) bool
	AddRole(role.Role)
	AddTechnology(technology.Technology)
	DeleteRole(role.Role) error
	DeleteTechnology(technology.Technology) error
}

// New creates a new Member object.
func New(name string) Member {
	member := &member.Member{}

	member.SetName(name)

	return member
}
