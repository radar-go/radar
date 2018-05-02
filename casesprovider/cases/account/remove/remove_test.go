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
	"testing"

	"github.com/radar-go/radar/casesprovider/helper"
	"github.com/radar-go/radar/datastore"
	"github.com/radar-go/radar/datastore/account"
)

func TestRemoveAccountCreation(t *testing.T) {
	uc := New()
	helper.TestCaseName(t, New(), "AccountRemove")

	uc.SetDatastore(datastore.New())
	_, err := uc.Run()
	helper.SaveGoldenData(t, "RemoveAccountCreation", []byte(err.Error()))
	expected := helper.GetGoldenData(t, "RemoveAccountCreation")
	helper.ContainsBytes(t, []byte(err.Error()), expected)
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
				id := helper.RegisterUser(t, uc.Datastore, tc.username, tc.name,
					tc.email, tc.password)
				if id != tc.id {
					t.Errorf("Expected %d, Got %d", tc.id, id)
				}
			}

			if tc.login {
				helper.LoginUser(t, uc.Datastore, tc.session, tc.username)
			}

			helper.AddParam(t, uc, "token", tc.session)
			helper.AddParam(t, uc, "id", tc.id)
			res, err := uc.Run()
			if err != nil {
				helper.SaveGoldenData(t, tc.testName+"_error", []byte(err.Error()))
				expected := helper.GetGoldenData(t, tc.testName+"_error")
				helper.ContainsBytes(t, []byte(err.Error()), expected)
				if err != tc.err {
					t.Errorf("Expect %s, Got %s", tc.err, err)
				}
			} else if err != tc.err {
				t.Errorf("Expected error to be %s, Got nil", tc.err)
			}

			actual := helper.GetResultBytes(t, res)
			helper.SaveGoldenData(t, tc.testName+"_result", actual)
			expected := helper.GetGoldenData(t, tc.testName+"_result")
			helper.ContainsBytes(t, actual, expected)
			if tc.checkAccount {
				_, err = uc.Datastore.GetAccountByUsername(tc.username)
				if err == nil {
					t.Error("Expected error getting the account from the datastore")
				} else if err != account.ErrAccountNotExists {
					t.Errorf("Expected ErrAccountNotExists, Got %s", err)
				}
			}
		})
	}
}
