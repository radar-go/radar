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
	"testing"
	"time"
)

type test struct {
	title    string
	started  time.Time
	finished time.Time
	active   bool
}

func TestRoleInterface(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		role, err := New(test.title, test.started, test.finished)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		if role.Title() != test.title {
			t.Errorf("Expected: %s, Got %s", test.title, role.Title())
		}

		if role.Started() != test.started {
			t.Errorf("Expected: %v, Got %v", test.started, role.Started())
		}

		if role.IsActive() != test.active {
			t.Errorf("Expected: %t, Got %t", test.active, role.IsActive())
		}
	}
}

func initializeTests() []test {
	return []test{
		test{
			title:    "Backend developer",
			started:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			finished: time.Date(2011, time.November, 10, 23, 0, 0, 0, time.UTC),
			active:   false,
		},
		test{
			title:    "Frontend developer",
			started:  time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC),
			finished: time.Date(2012, time.September, 10, 23, 0, 0, 0, time.UTC),
			active:   false,
		},
		test{
			title:    "Tech Lead",
			started:  time.Date(2013, time.November, 10, 23, 0, 0, 0, time.UTC),
			finished: time.Date(2016, time.April, 10, 23, 0, 0, 0, time.UTC),
			active:   false,
		},
		test{
			title:   "Manager",
			started: time.Date(2014, time.November, 10, 23, 0, 0, 0, time.UTC),
			active:  true,
		},
	}
}
