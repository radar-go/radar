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

	"github.com/goware/emailx"
	"github.com/pkg/errors"
)

func TestUser(t *testing.T) {
	usr, err := New("username", "name", "email@ritho.net", "password")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if usr.ID() != 1 {
		t.Errorf("Expected user id 1, Got %d", usr.ID())
	}

	if usr.Username() != "username" {
		t.Errorf("Expected name, got %s", usr.Name())
	}

	if usr.Name() != "name" {
		t.Errorf("Expected name, got %s", usr.Name())
	}

	if usr.Email() != "email@ritho.net" {
		t.Errorf("Expected 'email@ritho.net', Got %s", usr.email)
	}

	if usr.Password() != "password" {
		t.Errorf("Expected password, Got %s", usr.password)
	}
}

func TestUserFail(t *testing.T) {
	_, err := New("use", "name", "email@ritho.net", "password")
	if err != ErrUsernameTooShort {
		t.Errorf("Unexpected error: %s", err)
	}

	_, err = New("username", "name", "email@ritho.net", "pass")
	if err != ErrPasswordTooShort {
		t.Errorf("Unexpected error: %s", err)
	}

	_, err = New("username", "name", "email", "password")
	if errors.Cause(err) != emailx.ErrInvalidFormat {
		t.Errorf("Unexpected error: %s", err)
	}

	_, err = New("username", "name", "email@unknown.fake", "password")
	if errors.Cause(err) != emailx.ErrUnresolvableHost {
		t.Errorf("Unexpected error: %s", err)
	}
}
