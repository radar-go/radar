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

func TestEndpoints(t *testing.T) {
	ds := New()

	numEndpoints := 5
	endpoints := ds.Endpoints()
	if len(endpoints) != numEndpoints {
		t.Errorf("Expected %d, Got %d", numEndpoints, len(endpoints))
	}
}

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

func TestGetAccountSession(t *testing.T) {
	ds := New()

	_, err := ds.GetAccountBySession(" ")
	if err == nil {
		t.Error("Expected error getting the account by session.")
	} else if errors.Cause(err) != account.ErrUserNotLoggedIn {
		t.Errorf("Expected %s, Got %s", account.ErrUserNotLoggedIn, errors.Cause(err))
	}

	_, err = ds.GetAccountBySession("00000000-0000-0000-0000-000000000000")
	if err == nil {
		t.Error("Expected error getting the account by session.")
	} else if errors.Cause(err) != account.ErrUserNotLoggedIn {
		t.Errorf("Expected %s, Got %s", account.ErrUserNotLoggedIn, errors.Cause(err))
	}

	ds.sessions["00000000-0000-0000-0000-000000000000"] = &account.Account{}
	_, err = ds.GetAccountBySession("00000000-0000-0000-0000-000000000000")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
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

func TestUpdateAccount(t *testing.T) {
	acc := &account.Account{}
	session := "00000000-0000-0000-0000-000000000000"
	ds := New()

	err := ds.UpdateAccountData(acc, " ")
	if err == nil {
		t.Error("Expected error updating the account data")
	} else if errors.Cause(err) != account.ErrUserNotLoggedIn {
		t.Errorf("Expected %s, Got %s", account.ErrUserNotLoggedIn, errors.Cause(err))
	}

	err = ds.UpdateAccountData(acc, session)
	if err == nil {
		t.Error("Expected error updating the account data")
	} else if errors.Cause(err) != account.ErrUserNotLoggedIn {
		t.Errorf("Expected %s, Got %s", account.ErrUserNotLoggedIn, errors.Cause(err))
	}

	ds.sessions[session] = acc
	err = ds.UpdateAccountData(acc, session)
	if err == nil {
		t.Error("Expected error updating the account data")
	} else if errors.Cause(err) != account.ErrAccountNotExists {
		t.Errorf("Expected %s, Got %s", account.ErrAccountNotExists, errors.Cause(err))
	}

	ds.accounts[acc.Username()] = acc
	err = ds.UpdateAccountData(acc, session)
	if err != nil {
		t.Errorf("Unexpected error updating the accoung data: $s", err)
	}
}

func TestRemoveAccount(t *testing.T) {
	acc := &account.Account{}
	session := "00000000-0000-0000-0000-000000000000"
	ds := New()

	err := ds.RemoveAccount(acc)
	if err == nil {
		t.Error("Expected error removing the account")
	} else if errors.Cause(err) != account.ErrAccountNotExists {
		t.Errorf("Expected %s, Got %s", account.ErrAccountNotExists, errors.Cause(err))
	}

	ds.accounts[acc.Username()] = acc
	err = ds.RemoveAccount(acc)
	if err != nil {
		t.Errorf("Unexpected error removing the account: %s", err)
	}

	ds.accounts[acc.Username()] = acc
	ds.sessions[session] = acc
	err = ds.RemoveAccount(acc)
	if err != nil {
		t.Errorf("Unexpected error removing the account: %s", err)
	}
}
