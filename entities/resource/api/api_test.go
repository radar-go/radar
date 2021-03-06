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
)

func TestResource(t *testing.T) {
	res := New("Clean Arquitecture", "https://safari.oreilly.com/clean_arquitecture")

	if res.Name() != "Clean Arquitecture" {
		t.Errorf("Expected Clean Arquitecture, Got %s", res.Name())
	}

	if res.URL() != "https://safari.oreilly.com/clean_arquitecture" {
		t.Errorf("Expected https://safari.oreilly.com/clean_arquitecture, Got %s",
			res.URL())
	}
}
