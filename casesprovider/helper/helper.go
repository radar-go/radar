// Package helper contains helper functions to test the use cases.
package helper

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
	"fmt"
	"strings"
	"testing"

	"github.com/radar-go/radar/casesprovider"
	"github.com/radar-go/radar/datastore"
	"github.com/radar-go/radar/helper"
)

// SaveGoldenData saves test data in a golden file.
func SaveGoldenData(t *testing.T, name string, data []byte) {
	helper.SaveGoldenData(t, name, data)
}

// GetGoldenData gets data from a golden file.
func GetGoldenData(t *testing.T, name string) []byte {
	return helper.GetGoldenData(t, name)
}

// AddParam helper function to add a param to a use case.
func AddParam(t *testing.T, uc casesprovider.UseCase, param string, value interface{}) {
	t.Helper()

	err := uc.AddParam(param, value)
	if err != nil {
		t.Errorf("Unknown param %s", param)
	}
}

// UnexpectedError helper function to show an unexpected error in the tests.
func UnexpectedError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

// GetResultString helper function to convert the result to string.
func GetResultString(t *testing.T, res casesprovider.ResultPrinter) string {
	t.Helper()
	plainResult, err := res.String()
	UnexpectedError(t, err)

	return plainResult
}

// GetResultBytes helper function to convert the result to []byte.
func GetResultBytes(t *testing.T, res casesprovider.ResultPrinter) []byte {
	t.Helper()
	plainResult, err := res.Bytes()
	UnexpectedError(t, err)

	return plainResult
}

// Contains helper function to check if got contains expected in the tests.
func Contains(t *testing.T, got, expected string) bool {
	t.Helper()
	if !strings.Contains(got, expected) {
		t.Errorf("Expected '%s', Got '%s'", expected, got)
		return false
	}

	return true
}

// ContainsBytes helper function to check if got contains expected in the tests.
func ContainsBytes(t *testing.T, got, expected []byte) bool {
	t.Helper()
	if !bytes.Contains(got, expected) {
		t.Errorf("Expected '%s', Got '%s'", expected, got)
		return false
	}

	return true
}

// AddParams helper function to add params to an use case in the tests.
func AddParams(t *testing.T, uc casesprovider.UseCase, params map[string]interface{}) {
	t.Helper()
	for name, value := range params {
		AddParam(t, uc, name, value)
	}
}

// RegisterUser helper function to register an user in the datastore for the tests.
func RegisterUser(t *testing.T, ds *datastore.Datastore, username, name, email, password string) int {
	t.Helper()
	id, err := ds.AccountRegistration(username, name, email, password)
	UnexpectedError(t, err)

	return id
}

// LoginUser helper function to login an user into the datastore for the tests.
func LoginUser(t *testing.T, ds *datastore.Datastore, token, username string) {
	t.Helper()
	err := ds.AddSession(token, username)
	UnexpectedError(t, err)
}

// SetupUseCase helper function to initialize an use case for the tests.
func SetupUseCase(t *testing.T, uc casesprovider.UseCase, params map[string]interface{}) {
	t.Helper()
	uc.SetDatastore(datastore.New())
	AddParams(t, uc, params)
}

// ParamsTestCases struct to test the use case parameters.
type ParamsTestCases struct {
	Params        map[string]interface{}
	Expected      string
	ExpectedError bool
}

// TestCaseParams helper function to tests the use case parameters.
func TestCaseParams(t *testing.T, uc casesprovider.UseCase, testCases map[string]ParamsTestCases) {
	t.Helper()
	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			for name, value := range tc.Params {
				if tc.ExpectedError {
					err := uc.AddParam(name, value)
					if err == nil {
						t.Error("Expected error adding the tokens param.")
					}

					Contains(t, fmt.Sprint(err), tc.Expected)
				} else {
					AddParam(t, uc, name, value)
				}
			}
		})
	}
}

// TestCaseName helper function to test the use case name.
func TestCaseName(t *testing.T, uc casesprovider.UseCase, name string) {
	t.Helper()
	if uc.GetName() != name {
		t.Errorf("Expected %s, Got %s", name, uc.GetName())
	}

	newUC := uc.New()
	if newUC.GetName() != name {
		t.Errorf("Expected %s, Got %s", name, newUC.GetName())
	}
}
