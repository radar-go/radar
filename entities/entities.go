// Package entities contains the entities interface definitions.
package entities

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
	"time"
)

// Technology entity represents a technology used in one project, by one member
// or in one resource entity.
type Technology interface {
	GetName() string
	GetType() string
	GetLevel() string
}

// Role entity defines the role that a member have in the organization
// (developer, product owner, collaborator, ...)
type Role interface {
	GetTitle() string
	GetExperience() time.Time
}

// Member entity represents a member of the organization (employee, partner,
// associate, ...) and his relation  with the rest of entities.
type Member interface {
	GetName() string
	GetRoles() []Role
	GetCurrentRole() Role
	GetTechnologies() []Technology

	AddRole(newRole Role)
	AddTechnology(newTechnology Technology)
}

// Project entity represents a project done in the organization.
type Project interface {
	GetName() string
	GetMembers() []Member
	GetTechnologies() []Technology

	AddMember(newMember Member)
	AddTechnology(newTechnology Technology)
}

// Resource entity represents a resource to learn one or several technologies.
type Resource interface {
	GetName() string
	GetURL() string
	GetTechnologies() []Technology
	GetRate() float32

	AddTechnology(newTechnology Technology)
	AddRate(newRate float32)
}
