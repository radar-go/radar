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
	"testing"
	"time"

	role "github.com/radar-go/radar/entities/role/api"
	technology "github.com/radar-go/radar/entities/technology/api"
)

type test struct {
	name         string
	roles        []role.Role
	technologies []technology.Technology
}

func TestMember(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		m := &Member{
			name: test.name,
		}

		if m.Name() != test.name {
			t.Errorf("Expected %s, got %s", test.name, m.Name())
		}
	}
}

func initializeTests() []test {
	be, _ := role.New("Backend Developer",
		time.Date(2015, time.September, 1, 0, 0, 0, 0, time.UTC),
		time.Time{})
	betl, _ := role.New("Backend Tech Lead",
		time.Date(2017, time.August, 1, 0, 0, 0, 0, time.UTC),
		time.Time{})
	fe, _ := role.New("Frontend Developer",
		time.Date(2012, time.September, 1, 0, 0, 0, 0, time.UTC),
		time.Time{})
	fetl, _ := role.New("Frontend Tech Lead",
		time.Date(2017, time.August, 1, 0, 0, 0, 0, time.UTC),
		time.Time{})

	return []test{
		test{
			name: "Ritho",
			roles: []role.Role{
				be,
				betl,
			},
			technologies: []technology.Technology{
				technology.New("Golang", "Language", 1),
				technology.New("C", "Language", 1),
			},
		},
		test{
			name: "Pepe",
			roles: []role.Role{
				fe,
				fetl,
			},
			technologies: []technology.Technology{
				technology.New("Javascript", "Language", 1),
				technology.New("React", "Framework", 1),
			},
		},
		test{
			name:         "Lolailo",
			roles:        []role.Role{},
			technologies: []technology.Technology{},
		},
		test{
			name:         "Lerele",
			roles:        []role.Role{},
			technologies: []technology.Technology{},
		},
	}
}
