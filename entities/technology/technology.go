// Package technology implements the technology entity.
package technology

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

// Technology represents a technology used in a project, by an user or in a
// resource.
type Technology struct {
	name     string
	techType string
	level    int
}

// New returns a new Technology object.
func New(name, techType string, level int) *Technology {
	return &Technology{
		name:     name,
		techType: techType,
		level:    level,
	}
}

// GetName returns the name of the technology.
func (t *Technology) GetName() string {
	return t.name
}

// GetType returns the type of the technology.
func (t *Technology) GetType() string {
	return t.techType
}

// GetLevel returns the level of the technology.
func (t *Technology) GetLevel() int {
	return t.level
}
