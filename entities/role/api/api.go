// Package api defines the protocol for the role entity.
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
	"time"

	"github.com/radar-go/radar/entities/role"
)

// Role entity defines the role that a member have in the organization
// (developer, product owner, collaborator, ...)
type Role interface {
	Title() string
	Started() time.Time
	Experience() time.Duration
	IsActive() bool
	Equals(role interface{}) bool

	SetTitle(title string)
	SetStarted(t time.Time) error
	SetFinished(t time.Time) error
}

// New returns a new Role object.
func New(title string, started, finished time.Time) (*role.Role, error) {
	var err error

	r := &role.Role{}
	r.SetTitle(title)
	err = r.SetStarted(started)
	if err != nil {
		return r, err
	}

	err = r.SetFinished(finished)

	return r, err
}
