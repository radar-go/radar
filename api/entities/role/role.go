// Package role implements the role entity.
package role

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
	"errors"
	"time"
)

// Role represents a member role in the organization.
type Role struct {
	title    string
	started  time.Time
	finished time.Time
}

// Title returns the title of the role.
func (r *Role) Title() string {
	return r.title
}

// Started returns when the Role started to being applicable.
func (r *Role) Started() time.Time {
	return r.started
}

// Experience returns the time the member have been doing this Role.
func (r *Role) Experience() time.Duration {
	if !r.finished.IsZero() {
		return r.finished.Sub(r.started)
	}

	return time.Since(r.started)
}

// IsActive returns true if the Role is active and false otherwise.
func (r *Role) IsActive() bool {
	return r.finished.IsZero()
}

// SetTitle sets the title of the role.
func (r *Role) SetTitle(title string) {
	r.title = title
}

// SetStarted sets the starting time of the role.
func (r *Role) SetStarted(t time.Time) error {
	if !r.finished.IsZero() && t.After(r.finished) {
		return errors.New("The finishing time for the role is before from the starting time")
	}

	r.started = t

	return nil
}

// SetFinished sets the finishing time for the Role.
func (r *Role) SetFinished(t time.Time) error {
	if t.IsZero() {
		return nil
	}

	if t.Before(r.started) {
		return errors.New("The finishing time for the role is before from the starting time")
	}

	r.finished = t
	return nil
}

// Equals check if two role objects are equals.
func (r *Role) Equals(role interface{}) bool {
	switch role.(type) {
	case *Role:
		comp := role.(*Role)
		return r.Title() == comp.Title()
	case Role:
		comp := role.(Role)
		return r.Title() == (&comp).Title()
	default:
		return false
	}
}
