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
	"testing"

	"github.com/radar-go/radar/api/casesprovider/helper"
	"github.com/radar-go/radar/api/datastore"
)

func TestCaseName(t *testing.T) {
	helper.TestCaseName(t, New(), "AccountActivate")
}

func TestCaseParams(t *testing.T) {
	uc := New()
	uc.SetDatastore(datastore.New())

	testCases := map[string]helper.ParamsTestCases{
		"UnknownTokens": {
			Params: map[string]interface{}{
				"tokens": "00000000-0000-0000-0000-000000000000",
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
		"AddParamsSuccessfully": {
			Params: map[string]interface{}{
				"id":    1,
				"token": "00000000-0000-0000-0000-000000000000",
			},
			ExpectedError: false,
		},
	}
	helper.TestCaseParams(t, uc, testCases)
}

func TestAccountActivation(t *testing.T) {
	uc := New()

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
			var err error

			helper.SetupUseCase(t, uc, tc.params)
			if tc.register {
				id = helper.RegisterUser(t, uc.Datastore, tc.username, tc.name,
					tc.email, tc.password)
				helper.AddParam(t, uc, "id", id)
			}

			if tc.login {
				helper.LoginUser(t, uc.Datastore, tc.token, tc.username)
				helper.AddParam(t, uc, "token", tc.token)
			}

			res, err := uc.Run()
			if tc.expectedError {
				if err == nil {
					t.Error("Expected error running the use case")
				}

				helper.Contains(t, fmt.Sprint(err), tc.expected)
			} else {
				helper.UnexpectedError(t, err)
				plainResult := helper.GetResultString(t, res)
				helper.Contains(t, plainResult, fmt.Sprintf(`"id":%d`, id))
				helper.Contains(t, plainResult, tc.expected)
			}
		})
	}
}
