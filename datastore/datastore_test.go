package datastore

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

	"github.com/goware/emailx"
	"github.com/pkg/errors"

	"github.com/radar-go/radar/datastore/user"
)

func TestDatastoreRegisterUser(t *testing.T) {
	ds := New()

	id, err := ds.UserRegistration("ritho", "palvarez@ritho.net", "ritho")
	if err != nil {
		t.Errorf("Unexpected error registering an user: %+v", err)
	}

	if id != 1 {
		t.Errorf("Expected 1, got %d", id)
	}
}

func TestDatastoreEmailError(t *testing.T) {
	ds := New()

	_, err := ds.UserRegistration("ritho", "ritho", "ritho")
	if errors.Cause(err) != emailx.ErrInvalidFormat {
		t.Errorf("Unexpected error registering an user: %+v", err)
	}

	_, err = ds.UserRegistration("ritho", "palvarez@invalid.es", "ritho")
	if errors.Cause(err) != emailx.ErrUnresolvableHost {
		t.Errorf("Unexpected error registering an user: %+v", err)
	}
}

func TestDatastoreUserRegistered(t *testing.T) {
	ds := New()

	_, err := ds.UserRegistration("ritho", "palvarez@ritho.net", "ritho")
	if err != nil {
		t.Errorf("Unexpected error registering an user: %+v", err)
	}

	_, err = ds.UserRegistration("ritho", "palvarez@ritho.net", "ritho")
	if errors.Cause(err) != user.ErrUserExists {
		t.Errorf("Expected error %+v, Got %+v", user.ErrUserExists, err)
	}
}
