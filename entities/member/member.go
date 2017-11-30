// Package member implements the member entity.
package member

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

	"github.com/radar-go/radar/entities/member/api"
	role "github.com/radar-go/radar/entities/role/api"
	technology "github.com/radar-go/radar/entities/technology/api"
)

// Member entity represents a member of the organization (employee, partner,
// associate, ...) and his relation with the rest of entities.
type Member struct {
	name         string
	roles        []role.Role
	technologies []technology.Technology
}

// New creates a new Member object.
func New(name string) *Member {
	return &Member{
		name:         name,
		roles:        make([]role.Role, 0),
		technologies: make([]technology.Technology, 0),
	}
}

// Name returns the member name.
func (m *Member) Name() string {
	return m.name
}

// Roles returns all the roles that the member have had.
func (m *Member) Roles() []role.Role {
	return m.roles
}

// CurrentRole returns the role that the member have.
func (m *Member) CurrentRole() role.Role {
	var r role.Role
	if len(m.roles) == 0 {
		return nil
	}

	for _, role := range m.roles {
		if role.IsActive() {
			if r == nil {
				r = role
			} else if r.Started().Before(role.Started()) {
				r = role
			}
		}
	}

	if r == nil {
		r = m.roles[0]
	}

	return r
}

// Technologies returns the list of technologies the member knows.
func (m *Member) Technologies() []technology.Technology {
	return m.technologies
}

// SetName sets the member name.
func (m *Member) SetName(name string) {
	m.name = name
}

// Equals compares two member objects to check if they're the same one.
func (m *Member) Equals(member api.Member) bool {
	return m.Name() == member.Name()
}

// AddRole adds a new role to the member.
func (m *Member) AddRole(newRole role.Role) {
	m.roles = append(m.roles, newRole)
}

// AddTechnology adds a new technology to the known technologies of the member.
func (m *Member) AddTechnology(newTechnology technology.Technology) {
	m.technologies = append(m.technologies, newTechnology)
}

// DeleteRole deletes a role from the member.
func (m *Member) DeleteRole(role role.Role) error {
	for i, r := range m.roles {
		if r.Equals(role) {
			m.roles = append(m.roles[:i], m.roles[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No role %s is present in the member %s", role.Title(), m.Name())
}

// DeleteTechnology deletes a technology from the known technologies of the member.
func (m *Member) DeleteTechnology(tech technology.Technology) error {
	for i, t := range m.technologies {
		if t.Equals(tech) {
			m.technologies = append(m.technologies[:i], m.technologies[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No technology %s is present in the member %s", tech.Name(),
		m.Name())
}
