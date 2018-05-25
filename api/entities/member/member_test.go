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

	role "github.com/radar-go/radar/api/entities/role/api"
	technology "github.com/radar-go/radar/api/entities/technology/api"
)

type test struct {
	name string
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

		m.SetName(test.name)
		if m.Name() != test.name {
			t.Errorf("Expected %s, got %s", test.name, m.Name())
		}

		if !m.Equals(m) {
			t.Errorf("Expected %s member to be equals to %s", test.name, m.Name())
		}

		failMember := &Member{
			name: test.name + "2",
		}
		if m.Equals(failMember) {
			t.Errorf("Expected %s not to be equals to %s", failMember.Name(),
				m.Name())
		}
	}
}

func TestMemberRoles(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		m := &Member{
			name: test.name,
		}

		roles := m.Roles()
		if len(roles) > 0 {
			t.Errorf("Expected 0 roles, got %d", len(roles))
		}

		currentRole := m.CurrentRole()
		if currentRole != nil {
			t.Errorf("Expected the current role to be nil, got %s", currentRole.Title())
		}

		newRole, _ := role.New("Backend Developer",
			time.Date(2015, time.September, 1, 0, 0, 0, 0, time.UTC),
			time.Time{})

		m.AddRole(newRole)
		roles = m.Roles()
		if len(roles) != 1 {
			t.Errorf("Expected 1 roles, got %d", len(roles))
		}

		m.AddRole(newRole)
		roles = m.Roles()
		if len(roles) != 2 {
			t.Errorf("Expected 2 roles, got %d", len(roles))
		}

		err := m.DeleteRole(newRole)
		if err != nil {
			t.Errorf("Unexpected error deleting the role: %+v", err)
		}

		roles = m.Roles()
		if len(roles) != 1 {
			t.Errorf("Expected 1 roles, got %d", len(roles))
		}

		currentRole = m.CurrentRole()
		if !currentRole.Equals(newRole) {
			t.Errorf("Expected the current role to be %s, got %s", newRole.Title(),
				currentRole.Title())
		}

		deleteRole, _ := role.New("Tech Lead",
			time.Date(2015, time.September, 1, 0, 0, 0, 0, time.UTC),
			time.Time{})
		err = m.DeleteRole(deleteRole)
		if err == nil {
			t.Errorf("Expected error deleting the role %s", deleteRole.Title())
		}
	}
}

func TestMemberTechnologies(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		m := &Member{
			name: test.name,
		}

		techs := m.Technologies()
		if len(techs) > 0 {
			t.Errorf("Expected 0 technologies, got %d", len(techs))
		}

		tech := technology.New("Golang", "Language", 1)
		m.AddTechnology(tech)
		techs = m.Technologies()
		if len(techs) != 1 {
			t.Errorf("Expected 1 technologies, got %d", len(techs))
		}

		m.AddTechnology(tech)
		techs = m.Technologies()
		if len(techs) != 2 {
			t.Errorf("Expected 2 technologies, got %d", len(techs))
		}

		err := m.DeleteTechnology(tech)
		if err != nil {
			t.Errorf("Unexpected error deleting the technology: %+v", err)
		}

		techs = m.Technologies()
		if len(techs) != 1 {
			t.Errorf("Expected 1 technologies, got %d", len(techs))
		}

		deleteTech := technology.New("Linux", "os", 1)
		err = m.DeleteTechnology(deleteTech)
		if err == nil {
			t.Errorf("Expected error deleting the technology %s", deleteTech.Name())
		}
	}
}

func initializeTests() []test {
	return []test{
		test{
			name: "Ritho",
		},
		test{
			name: "Pepe",
		},
		test{
			name: "Lolailo",
		},
		test{
			name: "Lerele",
		},
	}
}
