package user

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

func TestUser(t *testing.T) {
	usr := New("name", "email", "password")

	if usr.ID() != 1 {
		t.Errorf("Expected user id 1, Got %d", usr.ID())
	}

	if usr.Name() != "name" {
		t.Errorf("Expected name, got %s", usr.Name())
	}

	if usr.email != "email" {
		t.Errorf("Expected email, Got %s", usr.email)
	}

	if usr.password != "password" {
		t.Errorf("Expected password, Got %s", usr.password)
	}
}
