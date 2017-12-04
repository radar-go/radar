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

// Name returns the name of the technology.
func (t *Technology) Name() string {
	return t.name
}

// Type returns the type of the technology.
func (t *Technology) Type() string {
	return t.techType
}

// Level returns the level of the technology.
func (t *Technology) Level() int {
	return t.level
}

// SetName sets the technology name.
func (t *Technology) SetName(newName string) {
	t.name = newName
}

// SetType sets the type of technology.
func (t *Technology) SetType(newType string) {
	t.techType = newType
}

// SetLevel sets the level of use of the technology.
func (t *Technology) SetLevel(newLevel int) {
	t.level = newLevel
}

// Equals check if two technology objects are equals or not.
func (t *Technology) Equals(tech interface{}) bool {
	switch tech.(type) {
	case Technology:
		comp := tech.(Technology)
		return t.Name() == (&comp).Name() && t.Type() == (&comp).Type()
	case *Technology:
		comp := tech.(*Technology)
		return t.Name() == comp.Name() && t.Type() == comp.Type()
	default:
		return false
	}
}
