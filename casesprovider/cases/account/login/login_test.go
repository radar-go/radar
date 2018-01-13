package login

/* Copyright (C) 2017-2018 Radar team (see AUTHORS)

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

func TestLogin(t *testing.T) {
	uc := New()
	if uc.Name != "AccountLogin" {
		t.Errorf("Expected Login, Got %s", uc.Name)
	}

	uc.SetDatastore(datastore.New())
	uc.AddParam("login", "ritho")
	uc.AddParam("password", "12345")
	_, err := uc.Run()
	if fmt.Sprintf("%v", err) != "ritho: User doesn't exists" {
		t.Errorf("Expected 'ritho: User doesn't exists', Got '%v'", err)
	}

	err = uc.AddParam("passwoed", "12345")
	if !strings.Contains(fmt.Sprintf("%v", err), "Error adding the param passwoed") {
		t.Errorf("Expected error to contain 'Error adding the param passwoed', Got '%v'",
			err)
	}

	id, err := uc.Datastore.UserRegistration("ritho", "ritho", "palvarez@ritho.net", "12345")
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	uc.AddParam("password", "123456")
	_, err = uc.Run()
	if fmt.Sprintf("%v", err) != "Password missmatch" {
		t.Errorf("Expected 'Password missmatch', Got '%v'", err)
	}

	uc.AddParam("password", "12345")
	res, err := uc.Run()
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	plainResult, _ := res.String()
	if !strings.Contains(plainResult, fmt.Sprintf(`"id":%d`, id)) {
		t.Errorf("Expected id %d in result '%s'", id, plainResult)
	}

	if !strings.Contains(plainResult, `"token":`) {
		t.Errorf("Expected to have token in the result '%s'", plainResult)
	}
}
