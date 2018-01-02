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

	"github.com/pkg/errors"

	"github.com/radar-go/radar/datastore/user"
)

func TestDatastoreRegisterUser(t *testing.T) {
	ds := New()

	id, err := ds.UserRegistration("ritho", "ritho", "palvarez@ritho.net", "ritho")
	if err != nil {
		t.Errorf("Unexpected error registering an user: %+v", err)
	}

	if id != 1 {
		t.Errorf("Expected 1, got %d", id)
	}
}

func TestDatastoreUserRegistered(t *testing.T) {
	ds := New()

	_, err := ds.UserRegistration("ritho", "ritho", "", "ritho")
	if errors.Cause(err) != user.ErrEmailEmpty {
		t.Errorf("Expected '%v', Got '%v'", user.ErrEmailEmpty, err)
	}

	_, err = ds.UserRegistration("", "ritho", "palvarez@ritho.net", "ritho")
	if errors.Cause(err) != user.ErrUsernameEmpty {
		t.Errorf("Expected '%v', Got '%v'", user.ErrUsernameEmpty, err)
	}

	_, err = ds.UserRegistration("ritho", "ritho", "palvarez@ritho.net", "")
	if errors.Cause(err) != user.ErrPasswordEmpty {
		t.Errorf("Expected '%v', Got '%v'", user.ErrPasswordEmpty, err)
	}

	_, err = ds.UserRegistration("ritho", "ritho", "palvarez@ritho.net", "ritho")
	if err != nil {
		t.Errorf("Unexpected error registering an user: %+v", err)
	}

	_, err = ds.UserRegistration("ritho", "ritho", "palvarez@ritho.net", "ritho")
	if errors.Cause(err) != user.ErrUserExists {
		t.Errorf("Expected error %+v, Got %+v", user.ErrUserExists, err)
	}
}

func TestDatastoreGetUser(t *testing.T) {
	ds := New()

	_, err := ds.GetUser("ritho")
	if fmt.Sprintf("%v", err) != "ritho: User doesn't exists" {
		t.Errorf("Expected 'ritho: User doesn't exists', Got '%v'", err)
	}

	ds.users["ritho"] = &user.User{}
	_, err = ds.GetUser("ritho")
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}
}

func TestDatastoreLogin(t *testing.T) {
	ds := New()

	err := ds.Login("1234", "ritho")
	if fmt.Sprintf("%v", err) != "ritho: User doesn't exists" {
		t.Errorf("Expected 'ritho: User doesn't exists', Got '%v'", err)
	}

	ds.users["ritho"] = &user.User{}
	err = ds.Login("1234", "ritho")
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	err = ds.Login("1234", "ritho")
	if fmt.Sprintf("%v", err) != "ritho: User already logged in" {
		t.Errorf("Expected 'ritho: User already logged in', Got '%v'", err)
	}
}

func TestDatastoreLogout(t *testing.T) {
	ds := New()

	ds.users["ritho"] = &user.User{}
	err := ds.Login("1234", "ritho")
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	err = ds.Logout("1234", "ritho")
	if err != nil {
		t.Errorf("Unexpected error '%v'", err)
	}

	err = ds.Logout("1234", "ritho")
	if fmt.Sprintf("%v", err) != "ritho: User not logged in" {
		t.Errorf("Expected 'ritho: User not logged in', Got '%v'", err)
	}

	err = ds.Logout("1234", "rit")
	if fmt.Sprintf("%v", err) != "rit: User doesn't exists" {
		t.Errorf("Expected 'rit: User doesn't exists', Got '%v'", err)
	}
}
