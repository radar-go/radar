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
	"testing"
	"time"
)

type test struct {
	title    string
	started  time.Time
	finished time.Time
	active   bool
	expected *Role
}

func TestRole(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		role := New(test.title, test.started, test.finished)
		if role.Title() != test.expected.title {
			t.Errorf("Expected: %s, Got %s", test.expected.title, role.Title())
		}

		if role.IsActive() != test.active {
			t.Errorf("Expected active to be %t, got %t", test.active, role.IsActive())
		}

		if !test.expected.finished.IsZero() {
			if role.Experience() != test.expected.finished.Sub(test.expected.started) {
				t.Errorf("Expected: %+v, Got %+v",
					test.expected.finished.Sub(test.expected.started),
					role.Experience())
			}
		} else {
			if role.Experience() <= 0 {
				t.Errorf("Expected > 0, Got %+v", role.Experience())
			}
		}

		err := role.SetStarted(test.started)
		if err != nil {
			t.Errorf("Unexpected error setting the started time for the role: %+v", err)
		}

		err = role.SetFinished(time.Now())
		if err != nil {
			t.Errorf("Unexpected error setting the finish time for the role: %+v", err)
		}

		if role.Experience() <= 0 {
			t.Errorf("Expected > 0, Got %+v", role.Experience())
		}

		if !role.Equals(test.expected) {
			t.Errorf("Expected '%s', got '%s'", test.expected.title, role.Title())
		}

		err = role.SetStarted(time.Now())
		if err == nil {
			t.Error("Expected to have an error setting the started role")
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
			expected: &Role{
				title:    "Backend developer",
				started:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				finished: time.Date(2011, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
		},
		test{
			title:    "Frontend developer",
			started:  time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC),
			finished: time.Date(2012, time.September, 10, 23, 0, 0, 0, time.UTC),
			active:   false,
			expected: &Role{
				title:    "Frontend developer",
				started:  time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC),
				finished: time.Date(2012, time.September, 10, 23, 0, 0, 0, time.UTC),
			},
		},
		test{
			title:    "Tech Lead",
			started:  time.Date(2013, time.November, 10, 23, 0, 0, 0, time.UTC),
			finished: time.Date(2016, time.April, 10, 23, 0, 0, 0, time.UTC),
			active:   false,
			expected: &Role{
				title:    "Tech Lead",
				started:  time.Date(2013, time.November, 10, 23, 0, 0, 0, time.UTC),
				finished: time.Date(2016, time.April, 10, 23, 0, 0, 0, time.UTC),
			},
		},
		test{
			title:   "Manager",
			started: time.Date(2014, time.November, 10, 23, 0, 0, 0, time.UTC),
			active:  true,
			expected: &Role{
				title:   "Manager",
				started: time.Date(2014, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
		},
	}
}
