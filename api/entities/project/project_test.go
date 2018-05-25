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
	"testing"

	member "github.com/radar-go/radar/api/entities/member/api"
	technology "github.com/radar-go/radar/api/entities/technology/api"
)

type test struct {
	name string
}

func TestProject(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		p := &Project{
			name: test.name,
		}

		if p.Name() != test.name {
			t.Errorf("Expected %s, Got %s", test.name, p.Name())
		}

		p.SetName(test.name)
		if p.Name() != test.name {
			t.Errorf("Expected %s, Got %s", test.name, p.Name())
		}
	}
}

func TestProjectMember(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		p := &Project{
			name: test.name,
		}

		members := p.Members()
		if len(members) > 0 {
			t.Errorf("Expected the list of members to be 0, got %d", len(members))
		}

		newMember := member.New("Ritho")
		p.AddMember(newMember)
		members = p.Members()
		if len(members) != 1 {
			t.Errorf("Expected the list of members to be 1, got %d", len(members))
		}

		p.AddMember(newMember)
		members = p.Members()
		if len(members) != 2 {
			t.Errorf("Expected the list of members to be 2, got %d", len(members))
		}

		err := p.DeleteMember(newMember)
		if err != nil {
			t.Errorf("Unexpected error deleting a member: %+v", err)
		}

		members = p.Members()
		if len(members) != 1 {
			t.Errorf("Expected the list of members to be 1, got %d", len(members))
		}

		deleteMember := member.New("Lolailo")
		err = p.DeleteMember(deleteMember)
		if err == nil {
			t.Errorf("Expected error deleting the member %s", deleteMember.Name())
		}
	}
}

func TestProjectTechnology(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		p := &Project{
			name: test.name,
		}

		techs := p.Technologies()
		if len(techs) > 0 {
			t.Errorf("Expected the list of technologies to be 0, got %d", len(techs))
		}

		newTechnology := technology.New("Golang", "language", 1)
		p.AddTechnology(newTechnology)
		techs = p.Technologies()
		if len(techs) != 1 {
			t.Errorf("Expected the list of technologies to be 1, got %d", len(techs))
		}

		p.AddTechnology(newTechnology)
		techs = p.Technologies()
		if len(techs) != 2 {
			t.Errorf("Expected the list of technologies to be 2, got %d", len(techs))
		}

		err := p.DeleteTechnology(newTechnology)
		if err != nil {
			t.Errorf("Unexpected error deleting a technology: %+v", err)
		}

		techs = p.Technologies()
		if len(techs) != 1 {
			t.Errorf("Expected the list of technologies to be 1, got %d", len(techs))
		}

		deleteTechnology := technology.New("Linux", "os", 1)
		err = p.DeleteTechnology(deleteTechnology)
		if err == nil {
			t.Errorf("Expected error deleting the technology %s", deleteTechnology.Name())
		}
	}
}

func initializeTests() []test {
	return []test{
		test{
			name: "phoenix",
		},
		test{
			name: "serenity",
		},
		test{
			name: "rocket",
		},
	}
}
