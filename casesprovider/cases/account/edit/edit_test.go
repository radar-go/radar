package edit

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

	"github.com/goware/emailx"
	"github.com/pkg/errors"

	"github.com/radar-go/radar/datastore"
	"github.com/radar-go/radar/datastore/user"
)

func TestEditCaseCreation(t *testing.T) {
	uc := New()
	if uc.Name != "AccountEdit" {
		t.Errorf("Expected 'AccountEdit', got '%s'", uc.Name)
	}

	ucNew := uc.New()
	if ucNew.GetName() != "AccountEdit" {
		t.Errorf("Expected 'AccountEdit', got '%s'", ucNew.GetName())
	}

	uc.SetDatastore(datastore.New())
	_, err := uc.Run()
	if !strings.Contains(fmt.Sprintf("%s", err), "User not logged in") {
		t.Error(err)
	}
}

func initializeTests(t *testing.T, session string) (*UseCase, int) {
	uc := New()
	uc.SetDatastore(datastore.New())

	/* User data. */
	user := "ritho"
	name := "ritho"
	email := "palvarez@ritho.net"
	password := "121212"

	/* Register the user. */
	id, err := uc.Datastore.UserRegistration(user, name, email, password)
	if err != nil {
		t.Errorf("Unexpected error registering the user: %s", err)
	}

	/* Login the user. */
	err = uc.Datastore.AddSession(session, user)
	if err != nil {
		t.Errorf("Unexpected error setting a session for the user: %s", err)
	}

	return uc, id
}

func addParam(t *testing.T, uc *UseCase, name string, value interface{}) {
	err := uc.AddParam(name, value)
	if err != nil {
		t.Errorf("Unexpected error adding the param '%s': %s", name, err)
	}
}

func TestEdit(t *testing.T) {
	/* Test initialization. */
	session := "00000000-0000-0000-0000-000000000000"
	uc, id := initializeTests(t, session)

	testCases := []struct {
		testName string
		username string
		name     string
		email    string
		password string
		err      error
		output   string
		compare  bool
	}{
		{
			"Success",
			"senoritho",
			"senoritho",
			"i02sopop@gmail.com",
			"212121",
			nil,
			fmt.Sprintf(`"id":%d`, id),
			true,
		},
		{
			"ErrorEmailInvalidFormat",
			"senoritho",
			"senoritho",
			"1@1",
			"212121",
			emailx.ErrInvalidFormat,
			"{}",
			false,
		},
		{
			"ErrorEmailUnresolvableHost",
			"senoritho",
			"senoritho",
			"email@domain.fakedomain",
			"212121",
			emailx.ErrUnresolvableHost,
			"{}",
			false,
		},
		{
			"ErrorUsernameShort",
			"s",
			"senoritho",
			"i02sopop@gmail.com",
			"212121",
			user.ErrUsernameTooShort,
			"{}",
			false,
		},
		{
			"ErrorPasswordShort",
			"senoritho",
			"senoritho",
			"i02sopop@gmail.com",
			"121",
			user.ErrPasswordTooShort,
			"{}",
			false,
		},
	}

	/* Add the new user data to the edit use case as params. */
	err := uc.AddParam("idd", id)
	if !strings.Contains(fmt.Sprintf("%s", err), "Unknown parameter for the use case") {
		t.Errorf("Unexpected error adding the 'idd' param: %s", err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			addParam(t, uc, "id", id)
			addParam(t, uc, "token", session)
			addParam(t, uc, "username", tc.username)
			addParam(t, uc, "name", tc.name)
			addParam(t, uc, "email", tc.email)
			addParam(t, uc, "password", tc.password)

			/* Edit the user. */
			res, err := uc.Run()
			if errors.Cause(err) != tc.err {
				t.Errorf("Unexpected error: %s", err)
			}

			/* Check the output. */
			ucRes, err := res.String()
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if !strings.Contains(ucRes, tc.output) {
				t.Errorf("Expected result to have %s, Got %s", tc.output, ucRes)
			}

			if tc.compare {
				/* Get the user from the session. */
				userData, err := uc.Datastore.GetUserBySession(session)
				if err != nil {
					t.Errorf("User %s doesn't have session in the datastore", tc.username)
				}

				/* Check if the user fields have been properly modified. */
				if userData.Name() != tc.name {
					t.Errorf("Expected %s, Got %s", tc.name, userData.Name())
				}

				if userData.Username() != tc.username {
					t.Errorf("Expected %s, Got %s", tc.username, userData.Username())
				}

				if userData.ID() != id {
					t.Errorf("Expected %d, Got %d", id, userData.ID())
				}

				if userData.Email() != tc.email {
					t.Errorf("Expected %s, Got %s", tc.email, userData.Email())
				}

				if userData.Password() != tc.password {
					t.Errorf("Expected %s, Got %s", tc.password, userData.Password())
				}

				/* Get the user from the data stored. */
				userDatastore, err := uc.Datastore.GetUserByUsername(tc.username)
				if err != nil {
					t.Errorf("User %s is not registered in the datastore", tc.username)
				}

				if !userDatastore.Equals(userData) {
					t.Error("User from datastore and user from session are not the same.")
				}
			}
		})
	}
}

func TestEditLogoutError(t *testing.T) {
	/* Test initialization. */
	session := "00000000-0000-0000-0000-000000000000"
	uc, id := initializeTests(t, session)

	/* Logout the user. */
	err := uc.Datastore.DeleteSession(session, "ritho")
	if err != nil {
		t.Errorf("Unexpected error removing a session for the user: %s", err)
	}

	addParam(t, uc, "id", id)
	addParam(t, uc, "token", session)
	addParam(t, uc, "username", "senoritho")
	addParam(t, uc, "name", "senoritho")
	addParam(t, uc, "email", "i02sopop@gmail.com")
	addParam(t, uc, "password", "212121")

	/* Edit the user. */
	res, err := uc.Run()
	if !strings.Contains(fmt.Sprintf("%s", err), "User not logged in") {
		t.Errorf("Unexpected error running the use case: %s", err)
	}

	ucRes, err := res.String()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if ucRes != "{}" {
		t.Error("Expected result to be empty")
	}
}

/* XXX: Test removing the user from the datastore when implemented.*/
