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

	"github.com/radar-go/radar/api/casesprovider/helper"
	"github.com/radar-go/radar/api/datastore"
	"github.com/radar-go/radar/api/datastore/account"
)

func TestEditCaseCreation(t *testing.T) {
	uc := New()
	helper.TestCaseName(t, uc, "AccountEdit")

	uc.SetDatastore(datastore.New())
	_, err := uc.Run()
	if !strings.Contains(fmt.Sprintf("%s", err), "User not logged in") {
		t.Error(err)
	}
}

func TestCaseParams(t *testing.T) {
	uc := New()
	uc.SetDatastore(datastore.New())

	testCases := map[string]helper.ParamsTestCases{
		"UnknownParam": {
			Params: map[string]interface{}{
				"idd": 1,
			},
			Expected:      "key doesn't exists: Unknown parameter for the use case",
			ExpectedError: true,
		},
		"IdFormatError": {
			Params: map[string]interface{}{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			Expected:      "id: Param is not from the right type",
			ExpectedError: true,
		},
		"TokenFormatError": {
			Params: map[string]interface{}{
				"token": 1,
			},
			Expected:      "token: Param is not from the right type",
			ExpectedError: true,
		},
		"UsernameFormatError": {
			Params: map[string]interface{}{
				"username": 1,
			},
			Expected:      "username: Param is not from the right type",
			ExpectedError: true,
		},
		"NameFormatError": {
			Params: map[string]interface{}{
				"name": 1,
			},
			Expected:      "name: Param is not from the right type",
			ExpectedError: true,
		},
		"EmailFormatError": {
			Params: map[string]interface{}{
				"email": 1,
			},
			Expected:      "email: Param is not from the right type",
			ExpectedError: true,
		},
		"PasswordFormatError": {
			Params: map[string]interface{}{
				"password": 1,
			},
			Expected:      "password: Param is not from the right type",
			ExpectedError: true,
		},
		"AddParamsSuccessfully": {
			Params: map[string]interface{}{
				"id":       1,
				"token":    "00000000-0000-0000-0000-000000000000",
				"username": "ritho",
				"name":     "ritho",
				"email":    "palvarez@ritho.net",
				"password": "121212",
			},
			ExpectedError: false,
		},
	}
	helper.TestCaseParams(t, uc, testCases)
}

func TestEdit(t *testing.T) {
	/* Test initialization. */
	session := "00000000-0000-0000-0000-000000000000"
	uc, id := initializeTests(t, session)

	testCases := []struct {
		testName string
		params   map[string]interface{}
		err      error
		output   string
		compare  bool
	}{
		{
			"Success",
			map[string]interface{}{
				"username": "senoritho",
				"name":     "senoritho",
				"email":    "i02sopop@gmail.com",
				"password": "212121",
			},
			nil,
			fmt.Sprintf(`"id":%d`, id),
			true,
		},
		{
			"ErrorEmailInvalidFormat",
			map[string]interface{}{
				"username": "senoritho",
				"name":     "senoritho",
				"email":    "1@1",
				"password": "212121",
			},
			emailx.ErrInvalidFormat,
			"{}",
			false,
		},
		{
			"ErrorEmailUnresolvableHost",
			map[string]interface{}{
				"username": "senoritho",
				"name":     "senoritho",
				"email":    "email@domain.fakedomain",
				"password": "212121",
			},
			emailx.ErrUnresolvableHost,
			"{}",
			false,
		},
		{
			"ErrorUsernameShort",
			map[string]interface{}{
				"username": "s",
				"name":     "senoritho",
				"email":    "i02sopop@gmail.com",
				"password": "212121",
			},
			account.ErrUsernameTooShort,
			"{}",
			false,
		},
		{
			"ErrorPasswordShort",
			map[string]interface{}{
				"username": "senoritho",
				"name":     "senoritho",
				"email":    "i02sopop@gmail.com",
				"password": "121",
			},
			account.ErrPasswordTooShort,
			"{}",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			helper.AddParam(t, uc, "id", id)
			helper.AddParam(t, uc, "token", session)
			helper.AddParams(t, uc, tc.params)

			/* Edit the account. */
			res, err := uc.Run()
			if err != tc.err {
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
				/* Get the account from the session. */
				accountData, err := uc.Datastore.GetAccountBySession(session)
				if err != nil {
					t.Errorf("User %s doesn't have session in the datastore", tc.params["username"])
				}

				/* Check if the account fields have been properly modified. */
				if accountData.Name() != tc.params["name"] {
					t.Errorf("Expected %s, Got %s", tc.params["name"], accountData.Name())
				}

				if accountData.Username() != tc.params["username"] {
					t.Errorf("Expected %s, Got %s", tc.params["username"],
						accountData.Username())
				}

				if accountData.ID() != id {
					t.Errorf("Expected %d, Got %d", id, accountData.ID())
				}

				if accountData.Email() != tc.params["email"] {
					t.Errorf("Expected %s, Got %s", tc.params["email"], accountData.Email())
				}

				if accountData.Password() != tc.params["password"] {
					t.Errorf("Expected %s, Got %s", tc.params["password"],
						accountData.Password())
				}

				/* Get the user from the data stored. */
				accountDatastore, err := uc.Datastore.GetAccountByUsername(tc.params["username"].(string))
				if err != nil {
					t.Errorf("User %s is not registered in the datastore", tc.params["username"])
				}

				if !accountDatastore.Equals(accountData) {
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
	helper.UnexpectedError(t, err)

	helper.AddParam(t, uc, "id", id)
	helper.AddParam(t, uc, "token", session)
	helper.AddParam(t, uc, "username", "senoritho")
	helper.AddParam(t, uc, "name", "senoritho")
	helper.AddParam(t, uc, "email", "i02sopop@gmail.com")
	helper.AddParam(t, uc, "password", "212121")

	/* Edit the account. */
	res, err := uc.Run()
	helper.Contains(t, fmt.Sprintf("%s", err), "User not logged in")

	ucRes, err := res.String()
	helper.UnexpectedError(t, err)
	if ucRes != "{}" {
		t.Error("Expected result to be empty")
	}
}

/* XXX: Test removing the account from the datastore when implemented.*/

func initializeTests(t *testing.T, session string) (*UseCase, int) {
	t.Helper()
	uc := New()
	uc.SetDatastore(datastore.New())

	/* Account data. */
	user := "ritho"
	name := "ritho"
	email := "palvarez@ritho.net"
	password := "121212"

	/* Register the account. */
	id := helper.RegisterUser(t, uc.Datastore, user, name, email, password)

	/* Login the user. */
	helper.LoginUser(t, uc.Datastore, session, user)

	return uc, id
}
