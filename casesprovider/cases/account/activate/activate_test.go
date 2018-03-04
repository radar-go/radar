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

func unexpectedError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
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
	err := uc.AddParam("tokens", "12345")
	if !strings.Contains(fmt.Sprintf("%v", err), "Error adding the param tokens") {
		t.Errorf("Expected error to contain 'Error adding the param tokens', Got '%s'",
			err)
	}

	testCases := map[string]struct {
		params        map[string]interface{}
		expected      string
		expectedError bool

		/* Data for user registration. */
		username string
		name     string
		email    string
		password string
		register bool

		/* Data for user login. */
		token string
		login bool
	}{
		"UserNotRegistered": {
			params: map[string]interface{}{
				"id":    1,
				"token": "00000000-0000-0000-0000-000000000000",
			},
			expected:      "00000000-0000-0000-0000-000000000000: User not logged in",
			expectedError: true,
		},
		"UserNotLogged": {
			params: map[string]interface{}{
				"token": "00000000-0000-0000-0000-000000000000",
			},
			expected:      "00000000-0000-0000-0000-000000000000: User not logged in",
			expectedError: true,
			username:      "ritho",
			name:          "ritho",
			email:         "palvarez@ritho.net",
			password:      "121212",
			register:      true,
		},
		"ActivationSuccess": {
			params:        make(map[string]interface{}),
			expected:      "Account activated successfully",
			expectedError: false,
			username:      "ritho",
			name:          "ritho",
			email:         "palvarez@ritho.net",
			password:      "121212",
			register:      true,
			token:         "00000000-0000-0000-0000-000000000000",
			login:         true,
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			var id int
			uc.SetDatastore(datastore.New())
			for name, value := range tc.params {
				addParam(t, uc, name, value)
			}

			if tc.register {
				id, err = uc.Datastore.AccountRegistration(tc.username, tc.name,
					tc.email, tc.password)
				unexpectedError(t, err)
				addParam(t, uc, "id", id)
			}

			if tc.login {
				err = uc.Datastore.AddSession(tc.token, tc.username)
				unexpectedError(t, err)
				addParam(t, uc, "token", tc.token)
			}

			res, err := uc.Run()
			if tc.expectedError {
				if err == nil {
					t.Error("Expected error running the use case")
				} else if fmt.Sprint(err) != tc.expected {
					t.Errorf("Expected '%s', Got '%s'", tc.expected, err)
				}
			} else {
				unexpectedError(t, err)
				plainResult, err := res.String()
				unexpectedError(t, err)

				if !strings.Contains(plainResult, fmt.Sprintf(`"id":%d`, id)) {
					t.Errorf("Expected id %d in result '%s'", id, plainResult)
				}

				if !strings.Contains(plainResult, tc.expected) {
					t.Errorf("Expected '%s', Got '%s'", tc.expected, plainResult)
				}
			}
		})
	}
}
