package datastore

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
	"testing"

	"github.com/goware/emailx"
	"github.com/pkg/errors"

	"github.com/radar-go/radar/datastore/account"
)

func TestDatastoreRegisterAccountSuccess(t *testing.T) {
	ds := New()

	id, err := ds.AccountRegistration("ritho", "ritho", "palvarez@ritho.net", "ritho")
	if err != nil {
		t.Errorf("Unexpected error registering an account: %+v", err)
	}

	if id != 1 {
		t.Errorf("Expected 1, got %d", id)
	}
}

func TestDatastoreAccountRegisterError(t *testing.T) {
	ds := New()

	_, err := ds.AccountRegistration("ritho", "ritho", "", "ritho")
	if errors.Cause(err) != emailx.ErrInvalidFormat {
		t.Errorf("Expected '%v', Got '%v'", account.ErrEmailEmpty, err)
	}

	_, err = ds.AccountRegistration("", "ritho", "palvarez@ritho.net", "ritho")
	if errors.Cause(err) != account.ErrUsernameTooShort {
		t.Errorf("Expected '%v', Got '%v'", account.ErrUsernameTooShort, err)
	}

	_, err = ds.AccountRegistration("ritho", "ritho", "palvarez@ritho.net", "")
	if errors.Cause(err) != account.ErrPasswordTooShort {
		t.Errorf("Expected '%v', Got '%v'", account.ErrPasswordTooShort, err)
	}

	_, err = ds.AccountRegistration("ritho", "ritho", "palvarez@ritho.net", "ritho")
	if err != nil {
		t.Errorf("Unexpected error registering an account: %+v", err)
	}

	_, err = ds.AccountRegistration("ritho", "ritho", "palvarez@ritho.net", "ritho")
	if errors.Cause(err) != account.ErrAccountExists {
		t.Errorf("Expected error %+v, Got %+v", account.ErrAccountExists, err)
	}
}

func TestDatastoreGetAccount(t *testing.T) {
	ds := New()

	_, err := ds.GetAccountByUsername("ritho")
	if fmt.Sprintf("%v", err) != "ritho: Account doesn't exists" {
		t.Errorf("Expected 'ritho: Account doesn't exists', Got '%v'", err)
	}

	ds.accounts["ritho"] = &account.Account{}
	_, err = ds.GetAccountByUsername("ritho")
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}
}

func TestDatastoreLogin(t *testing.T) {
	ds := New()

	err := ds.AddSession("00000000-0000-0000-0000-000000000000", "ritho")
	if fmt.Sprintf("%v", err) != "ritho: Account doesn't exists" {
		t.Errorf("Expected 'ritho: Account doesn't exists', Got '%v'", err)
	}

	ds.accounts["ritho"] = &account.Account{}
	err = ds.AddSession("00000000-0000-0000-0000-000000000000", "ritho")
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	err = ds.AddSession("00000000-0000-0000-0000-000000000000", "ritho")
	if fmt.Sprintf("%v", err) != "ritho: User already logged in" {
		t.Errorf("Expected 'ritho: User already logged in', Got '%v'", err)
	}
}

func TestDatastoreLogout(t *testing.T) {
	ds := New()

	ds.accounts["ritho"] = &account.Account{}
	err := ds.AddSession("00000000-0000-0000-0000-000000000000", "ritho")
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	err = ds.DeleteSession("00000000-0000-0000-0000-000000000000", "ritho")
	if err != nil {
		t.Errorf("Unexpected error '%v'", err)
	}

	err = ds.DeleteSession("00000000-0000-0000-0000-000000000000", "ritho")
	if fmt.Sprintf("%v", err) != "ritho: User not logged in" {
		t.Errorf("Expected 'ritho: User not logged in', Got '%v'", err)
	}

	err = ds.DeleteSession("00000000-0000-0000-0000-000000000000", "rit")
	if fmt.Sprintf("%v", err) != "rit: Account doesn't exists" {
		t.Errorf("Expected 'rit: Account doesn't exists', Got '%v'", err)
	}
}

func TestEndpoints(t *testing.T) {
	ds := New()

	numEndpoints := 4
	endpoints := ds.Endpoints()
	if len(endpoints) != numEndpoints {
		t.Errorf("Expected %d, Got %d", numEndpoints, len(endpoints))
	}
}
