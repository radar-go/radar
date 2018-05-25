// Package resource implements the resource entity.
package resource

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

	technology "github.com/radar-go/radar/api/entities/technology/api"
)

type test struct {
	name string
	url  string
}

func TestResource(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		r := &Resource{
			name: test.name,
			url:  test.url,
		}

		if r.Name() != test.name {
			t.Errorf("Expected %s, Got %s", test.name, r.Name())
		}

		r.SetName(test.name)
		if r.Name() != test.name {
			t.Errorf("Expected %s, Got %s", test.name, r.Name())
		}

		if r.URL() != test.url {
			t.Errorf("Expected %s, Got %s", test.url, r.URL())
		}

		r.SetURL(test.url)
		if r.URL() != test.url {
			t.Errorf("Expected %s, Got %s", test.url, r.URL())
		}

		if r.Rate() != 0.0 {
			t.Errorf("Expected 0.0, Got %f", r.Rate())
		}

		r.AddRate(2.0)
		if r.Rate() != 2.0 {
			t.Errorf("Expected 2.0, Got %f", r.Rate())
		}

		r.AddRate(1.0)
		if r.Rate() != 1.5 {
			t.Errorf("Expected 1.5, Got %f", r.Rate())
		}

		err := r.DeleteRate(1.0)
		if err != nil {
			t.Errorf("Unexpected error removing a rate: %+v", err)
		}

		if r.Rate() != 2.0 {
			t.Errorf("Expected 2.0, Got %f", r.Rate())
		}

		err = r.DeleteRate(1.0)
		if err == nil {
			t.Error("Expected error removing a rate didn't happen")
		}
	}
}

func TestResourceTechnology(t *testing.T) {
	tests := initializeTests()
	for _, test := range tests {
		r := &Resource{
			name: test.name,
			url:  test.url,
		}

		techs := r.Technologies()
		if len(techs) > 0 {
			t.Errorf("Expected 0 technologies, got %d", len(techs))
		}

		newTech := technology.New("Golang", "Language", 1)
		r.AddTechnology(newTech)
		techs = r.Technologies()
		if len(techs) != 1 {
			t.Errorf("Expected 1 technologies, got %d", len(techs))
		}

		r.AddTechnology(newTech)
		techs = r.Technologies()
		if len(techs) != 2 {
			t.Errorf("Expected 2 technologies, got %d", len(techs))
		}

		err := r.DeleteTechnology(newTech)
		if err != nil {
			t.Errorf("Unexpected error deleting the technology: %+v", err)
		}

		techs = r.Technologies()
		if len(techs) != 1 {
			t.Errorf("Expected 1 technologies, got %d", len(techs))
		}

		deleteTech := technology.New("Linux", "os", 1)
		err = r.DeleteTechnology(deleteTech)
		if err == nil {
			t.Errorf("Expected error deleting the technology %s", deleteTech.Name())
		}
	}
}

func initializeTests() []test {
	return []test{
		test{
			name: "Clean arquitecture",
			url:  "https://safari.oreilly.com/clean_arquitecture",
		},
		test{
			name: "Clean code",
			url:  "https://safari.oreilly.com/clean_code",
		},
	}
}
