package logout

/* Copyright (C) 2018 Radar team (see AUTHORS)

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
	"fmt"
	"strings"
	"testing"

	"github.com/radar-go/radar/datastore"
)

func TestLogout(t *testing.T) {
	uc := New()
	if uc.Name != "Logout" {
		t.Errorf("Expected Logout, Got %s", uc.Name)
	}

	uc.SetDatastore(datastore.New())
	_, err := uc.Run()
	if fmt.Sprintf("%v", err) != "Username too short" {
		t.Errorf("Expected 'Username too short', Got '%v'", err)
	}

	uc.AddParam("username", "rit")
	_, err = uc.Run()
	if fmt.Sprintf("%v", err) != "Username too short" {
		t.Errorf("Expected 'Username too short', Got '%v'", err)
	}

	uc.AddParam("username", "ritho")
	_, err = uc.Run()
	if fmt.Sprintf("%v", err) != "Token too short" {
		t.Errorf("Expected 'Token too short', Got '%v'", err)
	}

	uc.AddParam("token", "0000")
	_, err = uc.Run()
	if fmt.Sprintf("%v", err) != "Token too short" {
		t.Errorf("Expected 'Token too short', Got '%v'", err)
	}

	uc.AddParam("token", "00000000-0000-0000-0000-000000000000")
	_, err = uc.Run()
	if fmt.Sprintf("%v", err) != "ritho: User doesn't exists" {
		t.Errorf("Expected 'ritho: User doesn't exists', Got '%v'", err)
	}

	err = uc.AddParam("tokens", "12345")
	if !strings.Contains(fmt.Sprintf("%v", err), "Error adding the param tokens") {
		t.Errorf("Expected error to contain 'Error adding the param tokens', Got '%v'",
			err)
	}

	id, err := uc.Datastore.UserRegistration("ritho", "ritho", "palvarez@ritho.net", "12345")
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	_, err = uc.Run()
	if fmt.Sprintf("%v", err) != "ritho: User not logged in" {
		t.Errorf("Expected 'ritho: User not logged in', Got '%v'", err)
	}

	err = uc.Datastore.Login("00000000-0000-0000-0000-000000000000", "ritho")
	res, err := uc.Run()
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	plainResult, _ := res.String()
	if !strings.Contains(plainResult, fmt.Sprintf(`"id":%d`, id)) {
		t.Errorf("Expected id %d in result '%s'", id, plainResult)
	}

	if !strings.Contains(plainResult, `User logout successfully`) {
		t.Errorf("Expected to logout successfully, Got '%s'", plainResult)
	}
}