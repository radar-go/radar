package remove

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
	"bytes"
	"testing"

	"github.com/pkg/errors"

	"github.com/radar-go/radar/datastore"
	"github.com/radar-go/radar/datastore/account"
	"github.com/radar-go/radar/tests"
)

func TestRemoveAccountCreation(t *testing.T) {
	uc := New()
	if uc.Name != "AccountRemove" {
		t.Errorf("Expected 'AccountRemove', Got '%s'", uc.Name)
	}

	ucNew := uc.New()
	if ucNew.GetName() != "AccountRemove" {
		t.Errorf("Expected 'AccountRemove', Got '%s'", ucNew.GetName())
	}

	uc.SetDatastore(datastore.New())
	_, err := uc.Run()
	tests.SaveGoldenData(t, "user_not_logged", []byte(err.Error()))
	expected := tests.GetGoldenData(t, "user_not_logged")
	if !bytes.Contains([]byte(err.Error()), expected) {
		t.Error(err)
	}
}

func TestRemoveAccountSuccess(t *testing.T) {
	uc := New()
	uc.SetDatastore(datastore.New())

	/* Account data. */
	user := "ritho"
	name := "ritho"
	email := "palvarez@ritho.net"
	password := "121212"

	/* Register the account. */
	id, err := uc.Datastore.AccountRegistration(user, name, email, password)
	if err != nil {
		t.Errorf("Unexpected error registering the account: %s", err)
	}

	/* Login the user. */
	session := "00000000-0000-0000-0000-000000000000"
	err = uc.Datastore.AddSession(session, user)
	if err != nil {
		t.Errorf("Unexpected error setting a session for the user: %s", err)
	}

	// Remove the account.
	addParam(t, uc, "id", id)
	addParam(t, uc, "token", session)
	res, err := uc.Run()
	if err != nil {
		t.Error(err)
	}

	actual, err := res.Bytes()
	if err != nil {
		t.Errorf("Unexpected error:%s", err)
	}

	tests.SaveGoldenData(t, "remove_account_successfully", actual)
	expected := tests.GetGoldenData(t, "remove_account_successfully")
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %s, Got %s", expected, actual)
	}

	// Check that the account doesn't exists anymore in the datastore.
	_, err = uc.Datastore.GetAccountByUsername(user)
	if err == nil {
		t.Error("Expected error getting the account from the datastore")
	} else if errors.Cause(err) != account.ErrAccountNotExists {
		t.Errorf("Expected ErrAccountNotExists, Got %s", errors.Cause(err))
	}
}

func TestRemoveAccountError(t *testing.T) {
	t.Error("Test not yet implemented")
}

func addParam(t *testing.T, uc *UseCase, name string, value interface{}) {
	t.Helper()
	err := uc.AddParam(name, value)
	if err != nil {
		t.Errorf("Unexpected error adding the param '%s': %s", name, err)
	}
}
