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

// New returns a new Role object.
func New(title string, started, finished time.Time) *Role {
	var finish time.Time

	if !finished.IsZero() && finished.After(started) {
		finish = finished
	}

	return &Role{
		title:    title,
		started:  started,
		finished: finish,
	}
}

// GetTitle returns the title of the role.
func (r *Role) GetTitle() string {
	return r.title
}

// GetExperience returns the time the member have been doing this Role.
func (r *Role) GetExperience() time.Duration {
	if !r.finished.IsZero() {
		return r.finished.Sub(r.started)
	}

	return time.Since(r.started)
}

// Finished sets the finishing time for the Role.
func (r *Role) Finished(t time.Time) error {
	if t.Before(r.started) {
		return errors.New("The finishing time for the role is before from the starting time")
	}

	r.finished = t
	return nil
}
