package activate

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

func addParam(t *testing.T, uc *UseCase, param string, value interface{}) {
	t.Helper()

	err := uc.AddParam(param, value)
	if err != nil {
		t.Errorf("Unknown param %s", param)
	}
}

func TestAccountActivation(t *testing.T) {
	uc := New()
	if uc.Name != "AccountActivate" {
		t.Errorf("Expected AccountActivate, Got %s", uc.Name)
	}

	newUC := uc.New()
	if newUC.GetName() != "AccountActivate" {
		t.Errorf("Expected AccountActivate, Got %s", newUC.GetName())
	}

	uc.SetDatastore(datastore.New())
	addParam(t, uc, "id", 1)
	addParam(t, uc, "token", "00000000-0000-0000-0000-000000000000")

	_, err := uc.Run()
	if err == nil {
		t.Error("Expected error running the use case")
	} else if fmt.Sprint(err) != "00000000-0000-0000-0000-000000000000: User not logged in" {
		t.Errorf("Expected '00000000-0000-0000-0000-000000000000: User not logged in', got %s", err)
	}

	err = uc.AddParam("tokens", "12345")
	if !strings.Contains(fmt.Sprintf("%v", err), "Error adding the param tokens") {
		t.Errorf("Expected error to contain 'Error adding the param tokens', Got '%v'",
			err)
	}

	id, err := uc.Datastore.AccountRegistration("ritho", "ritho", "palvarez@ritho.net",
		"12345")
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	addParam(t, uc, "id", id)
	_, err = uc.Run()
	if err == nil {
		t.Error("Expected error running the use case")
	} else if fmt.Sprint(err) != "00000000-0000-0000-0000-000000000000: User not logged in" {
		t.Errorf("Expected '00000000-0000-0000-0000-000000000000: User not logged in', got %s", err)
	}

	err = uc.Datastore.AddSession("00000000-0000-0000-0000-000000000000", "ritho")
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	res, err := uc.Run()
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	plainResult, err := res.String()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !strings.Contains(plainResult, fmt.Sprintf(`"id":%d`, id)) {
		t.Errorf("Expected id %d in result '%s'", id, plainResult)
	}

	if !strings.Contains(plainResult, `Account activated successfully`) {
		t.Errorf("Expected to activate the account successfully, Got '%s'", plainResult)
	}
}
