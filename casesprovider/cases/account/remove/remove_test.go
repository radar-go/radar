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
	tests.SaveGoldenData(t, "RemoveAccountCreation", []byte(err.Error()))
	expected := tests.GetGoldenData(t, "RemoveAccountCreation")
	if !bytes.Contains([]byte(err.Error()), expected) {
		t.Error(err)
	}
}

func TestRemoveAccountError(t *testing.T) {
	testCases := []struct {
		testName     string
		id           int
		username     string
		name         string
		email        string
		password     string
		session      string
		err          error
		register     bool
		login        bool
		checkAccount bool
	}{
		{
			testName:     "AccountNotExists",
			id:           1,
			username:     "ritho",
			name:         "ritho",
			email:        "palvarez@ritho.net",
			password:     "121212",
			session:      "00000000-0000-0000-0000-000000000000",
			err:          account.ErrUserNotLoggedIn,
			register:     false,
			login:        false,
			checkAccount: false,
		},
		{
			testName:     "AccountNotLogin",
			id:           1,
			username:     "ritho",
			name:         "ritho",
			email:        "palvarez@ritho.net",
			password:     "121212",
			session:      "00000000-0000-0000-0000-000000000000",
			err:          account.ErrUserNotLoggedIn,
			register:     true,
			login:        false,
			checkAccount: false,
		},
		{
			testName:     "AccountRemoveSuccessful",
			id:           2,
			username:     "ritho",
			name:         "ritho",
			email:        "palvarez@ritho.net",
			password:     "121212",
			session:      "00000000-0000-0000-0000-000000000000",
			err:          nil,
			register:     true,
			login:        true,
			checkAccount: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var err error

			uc := New()
			uc.SetDatastore(datastore.New())

			if tc.register {
				id, err := uc.Datastore.AccountRegistration(tc.username, tc.name,
					tc.email, tc.password)
				if err != nil {
					t.Errorf("Unexpected error registering the account: %s", err)
				}

				if id != tc.id {
					t.Errorf("Expected %d, Got %d", tc.id, id)
				}
			}

			if tc.login {
				err = uc.Datastore.AddSession(tc.session, tc.username)
				if err != nil {
					t.Errorf("Unexpected error setting a session for the user: %s", err)
				}
			}

			addParam(t, uc, "token", tc.session)
			addParam(t, uc, "id", tc.id)

			res, err := uc.Run()
			if err != nil {
				tests.SaveGoldenData(t, tc.testName+"_error", []byte(err.Error()))
				expected := tests.GetGoldenData(t, tc.testName+"_error")
				if !bytes.Equal([]byte(err.Error()), expected) {
					t.Errorf("Expected %s, Got %s", expected, err)
				}

				if errors.Cause(err) != tc.err {
					t.Errorf("Expect %s, Got %s", tc.err, errors.Cause(err))
				}
			} else if err != tc.err {
				t.Errorf("Expected error to be %s, Got nil", tc.err)
			}

			actual, err := res.Bytes()
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			tests.SaveGoldenData(t, tc.testName+"_result", actual)
			expected := tests.GetGoldenData(t, tc.testName+"_result")
			if !bytes.Equal(actual, expected) {
				t.Errorf("Expected %s, Got %s", expected, actual)
			}

			if tc.checkAccount {
				_, err = uc.Datastore.GetAccountByUsername(tc.username)
				if err == nil {
					t.Error("Expected error getting the account from the datastore")
				} else if errors.Cause(err) != account.ErrAccountNotExists {
					t.Errorf("Expected ErrAccountNotExists, Got %s", errors.Cause(err))
				}
			}
		})
	}
}

func addParam(t *testing.T, uc *UseCase, name string, value interface{}) {
	t.Helper()
	err := uc.AddParam(name, value)
	if err != nil {
		t.Errorf("Unexpected error adding the param '%s': %s", name, err)
	}
}
