// Package project implements the project entity.
package project

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

	member "github.com/radar-go/radar/api/entities/member/api"
	technology "github.com/radar-go/radar/api/entities/technology/api"
)

// Project entity represents a project done in the organization.
type Project struct {
	name         string
	members      []member.Member
	technologies []technology.Technology
}

// Name returns the name of the project.
func (p *Project) Name() string {
	return p.name
}

// Members return the list of members belonging to this project.
func (p *Project) Members() []member.Member {
	return p.members
}

// Technologies return the list of technologies used in this project.
func (p *Project) Technologies() []technology.Technology {
	return p.technologies
}

// SetName sets the project name.
func (p *Project) SetName(name string) {
	p.name = name
}

// AddMember adds a new member to the project.
func (p *Project) AddMember(newMember member.Member) {
	p.members = append(p.members, newMember)
}

// AddTechnology adds a new technology to the project.
func (p *Project) AddTechnology(newTechnology technology.Technology) {
	p.technologies = append(p.technologies, newTechnology)
}

// DeleteMember deletes a member from the project.
func (p *Project) DeleteMember(member member.Member) error {
	for i, m := range p.members {
		if m.Equals(member) {
			p.members = append(p.members[:i], p.members[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No member %s is present in the project %s", member.Name(),
		p.Name())
}

// DeleteTechnology deletes a technology from the project.
func (p *Project) DeleteTechnology(tech technology.Technology) error {
	for i, t := range p.technologies {
		if t.Equals(tech) {
			p.technologies = append(p.technologies[:i], p.technologies[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No technology %s is present in the project %s", tech.Name(),
		p.Name())
}
